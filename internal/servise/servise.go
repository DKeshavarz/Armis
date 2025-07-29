package servise

import (
	"context"

	"github.com/DKeshavarz/armis/internal/storage"
)

type service struct {
	storage storage.StorageInterface
}

func New()*service{
	st := storage.NewStorage()
	return &service{
		storage: st,
	}
}

func (s *service)Put(ctx context.Context, key, value string) error{
	return nil
}
func (s *service)Get(ctx context.Context, key string) (string, error){
	return "nigga",nil
}
func (s *service)Delete(ctx context.Context, key string) error{
	return nil
}