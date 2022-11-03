package media_usecase

import (
	"context"
	"io"
	"mime/multipart"
)

type IStorageService interface {
	Write(ctx context.Context, fileName string, r io.Reader) error
	Put(ctx context.Context, dir string, r io.Reader) error
}

type usecase struct {
	storageService IStorageService
}

func New(storage IStorageService) *usecase {
	return &usecase{storageService: storage}
}

func (u *usecase) Upload(ctx context.Context, file *multipart.FileHeader) error {
	dir := "/app/avatar/"
	f, err := file.Open()
	if err != nil {
		return err
	}
	err = u.storageService.Put(ctx, dir, f)
	if err != nil {
		return err
	}

	return nil
}
