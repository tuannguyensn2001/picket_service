package storage_service

import (
	"context"
	"io"
)

type instance interface {
	Write(ctx context.Context, fileName string, r io.Reader) error
	Put(ctx context.Context, dir string, r io.Reader) error
}

type service struct {
	instance instance
}

func New(disk string, path string) *service {
	var instance instance
	if disk == "fs" {
		instance = newFs(path)
	}
	return &service{instance: instance}
}

func (s *service) Write(ctx context.Context, fileName string, r io.Reader) error {
	return s.instance.Write(ctx, fileName, r)
}

func (s *service) Put(ctx context.Context, dir string, r io.Reader) error {
	return s.instance.Put(ctx, dir, r)
}
