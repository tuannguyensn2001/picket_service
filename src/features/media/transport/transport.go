package media_transport

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	errpkg "picket/src/packages/err"
)

const PATH = "src/storage/public"

type IUsecase interface {
	Upload(ctx context.Context, file *multipart.FileHeader) error
}

type transport struct {
	usecase IUsecase
}

func New(ctx context.Context, usecase IUsecase) *transport {
	return &transport{usecase: usecase}
}
func (t *transport) Upload(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseMultipartForm(10 << 20)

	_, header, err := r.FormFile("file")
	if err != nil {
		err = errpkg.General.BadRequest
		json.NewEncoder(w).Encode(err)
		return
	}

	err = t.usecase.Upload(r.Context(), header)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"message": "upload success",
	}
	json.NewEncoder(w).Encode(resp)
	return
}
