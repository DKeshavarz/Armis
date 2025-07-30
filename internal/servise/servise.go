package servise

import (
	"context"

	"github.com/DKeshavarz/armis/internal/storage"
)

type service struct {
	storage storage.StorageInterface
}

func New()*service{
	st := storage.New()
	return &service{
		storage: st,
	}
}

func (s *service)Put(ctx context.Context, key, value string) error{
	return s.storage.Put(ctx, key, value)
}
func (s *service)Get(ctx context.Context, key string) (string, error){
	return s.storage.Get(ctx, key)
}
func (s *service)Delete(ctx context.Context, key string) error{
	return s.storage.Delete(ctx, key)
}