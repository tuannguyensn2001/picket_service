package middlewares

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/metadata"
	"net/http"
)

func HttpPath(ctx context.Context, r *http.Request) metadata.MD {
	md := make(map[string]string)

	if pattern, ok := runtime.HTTPPathPattern(ctx); ok {
		md["pattern"] = pattern
	}

	return metadata.New(md)
}
