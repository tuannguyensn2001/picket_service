package middlewares

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/status"
	"myclass_service/src/app"
	errpkg "myclass_service/src/packages/err"
	"net/http"
	"strings"
)

func HandleError(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {

	s := status.Convert(err)
	zap.S().Error(err)
	writer.Header().Set("Content-Type", "application/json")

	httpErr := errpkg.General.Internal
	if len(s.Details()) > 0 {
		item := s.Details()[0].(*errdetails.BadRequest_FieldViolation)
		http := &app.Error{}
		f := bufio.NewReader(strings.NewReader(item.GetDescription()))
		errDecode := json.NewDecoder(f).Decode(http)
		if errDecode == nil {
			httpErr = http
		}
	}
	writer.WriteHeader(httpErr.StatusCode)
	json.NewEncoder(writer).Encode(httpErr)
	return
}
