package media_transport

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	errpkg "myclass_service/src/packages/err"
	"net/http"
	"os"
)

const PATH = "src/storage/public"

type transport struct {
}

func New(ctx context.Context) *transport {
	return &transport{}
}
func (t *transport) Upload(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	r.ParseMultipartForm(10 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		err = errpkg.General.BadRequest
		json.NewEncoder(w).Encode(err)
		return
	}
	defer file.Close()

	dst, err := os.Create(fmt.Sprintf("%s/%s", PATH, header.Filename))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]interface{}{
		"message": "upload success",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
	return
}
