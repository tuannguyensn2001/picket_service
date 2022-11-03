package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"myclass_service/src/config"
	"myclass_service/src/packages/err"
	authpb "myclass_service/src/pb/auth"
	"myclass_service/src/routes"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

type handler = func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

func server(config config.IConfig) *cobra.Command {
	return &cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			err.LoadError()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

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
	)))
	reflection.Register(server)

	lis, err := net.Listen("tcp", config.GetGrpcAddress())
	if err != nil {
		zap.S().Fatalln(err)
	}

	routes.RouteGrpc(ctx, server)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	signal.Notify(sigint, os.Kill)

	go func() {
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
	gw := runtime.NewServeMux(runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {

		s := status.Convert(err)
		zap.S().Error(s.Details())
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		type Response struct {
			Message string `json:"message"`
		}
		resp := Response{
			Message: err.Error(),
		}
		json.NewEncoder(writer).Encode(resp)
	}))

	lists := []handler{authpb.RegisterAuthServiceHandler}

	for _, item := range lists {
		err = item(ctx, gw, conn)
		if err != nil {
			zap.S().Fatalln(err)
		}
	}

	gwServer := &http.Server{
		Addr:    ":21000",
		Handler: gw,
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
