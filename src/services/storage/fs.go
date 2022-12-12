package storage_service

import (
	"context"
	"fmt"
	"io"
	"os"
	"picket/src/packages/random"
)

type fs struct {
	path string
}

func newFs(path string) *fs {
	return &fs{path: path}
}

func (f *fs) Write(ctx context.Context, fileName string, r io.Reader) error {
	dst, err := os.Create(fmt.Sprintf("%s%s", f.path, fileName))
	if err != nil {
		return err
	}
	if _, err := io.Copy(dst, r); err != nil {
		return err
	}

	return nil
}

func (f *fs) Put(ctx context.Context, dir string, r io.Reader) error {
	fileName := fmt.Sprintf("%s%s/%s", f.path, dir, randompkg.String())
	return f.Write(ctx, fileName, r)
}
