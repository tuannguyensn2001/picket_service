package cmd

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"net"
	"net/http"
	"os"
	"os/signal"
	"picket/src/config"
	"picket/src/middlewares"
	"picket/src/packages/err"
	"picket/src/routes"
	"sync"
)

func server(config config.IConfig) *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			errpkg.LoadError()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			tp, _ := tracerProvider("http://localhost:14268/api/traces")
			otel.SetTracerProvider(tp)

			var wg sync.WaitGroup
			wg.Add(2)
			go runGrpc(ctx, config, &wg)

			go runGateway(ctx, config, &wg)
			wg.Wait()
			zap.S().Info("shutdown application")
		},
	}
}

func runGrpc(ctx context.Context, config config.IConfig, wg *sync.WaitGroup) {

	server := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		grpc_zap.UnaryServerInterceptor(zap.L()),
		grpc_recovery.UnaryServerInterceptor([]grpc_recovery.Option{grpc_recovery.WithRecoveryHandlerContext(middlewares.HandleGrpcError)}...),
		grpc_auth.UnaryServerInterceptor(middlewares.Auth(config)),
	)))
	reflection.Register(server)

	lis, err := net.Listen("tcp", config.GetGrpcAddress())
	if err != nil {
		zap.S().Fatalln(err)
	}

	routes.RouteGrpc(ctx, server, config)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, os.Kill)

	go func() {
		zap.S().Info("grpc is running at :" + config.GetGrpcAddress())
		err = server.Serve(lis)
		if err != nil {
			zap.S().Fatalln(err)
		}
	}()

	<-sigint
	server.GracefulStop()
	wg.Done()
	zap.S().Info("shutdown grpc server")
}

func runGateway(ctx context.Context, config config.IConfig, wg *sync.WaitGroup) {
	conn, err := grpc.DialContext(ctx, config.GetGrpcAddress(), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Fatalln(err)
	}
	gw := runtime.NewServeMux(runtime.WithErrorHandler(middlewares.HandleError), runtime.WithMetadata(middlewares.HttpPath), runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	}))

	routes.RouteGw(ctx, gw, conn)

	c := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
	})

	handler := c.Handler(gw)

	gwServer := &http.Server{
		Addr:    config.GetHttpAddress(),
		Handler: handler,
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, os.Kill)

	go func() {
		zap.S().Info(fmt.Sprintf("grpc gateway server is running at %s", "21000"))
		err = gwServer.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			zap.S().Fatalln(err)
		}
	}()

	<-sigint
	gwServer.Shutdown(ctx)
	wg.Done()
	zap.S().Info("shutdown grpc gateway server")
}

func tracerProvider(url string) (*tracesdk.TracerProvider, error) {
	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("my-class-service"),
		)),
	)
	return tp, nil
}
