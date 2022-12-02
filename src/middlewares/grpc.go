package middlewares

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"picket/src/app"
)

func HandleGrpcError(ctx context.Context, p interface{}) error {
	zap.S().Error(p)

	err, ok := p.(*app.Error)
	if !ok {
		return status.Error(codes.Internal, "internal server error")
	}
	str, _ := err.ToJSON()

	resp := status.New(codes.Code(err.GrpcCode), err.Message)

	f1 := &errdetails.BadRequest_FieldViolation{
		Field:       "json",
		Description: str,
	}
	resp, _ = resp.WithDetails(f1)

	return resp.Err()
}
