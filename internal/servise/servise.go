package servise

import (
	"context"
	"log"

	"github.com/DKeshavarz/armis/internal/storage"
	"github.com/google/uuid"
)

type service struct {
	id string
	isMaster bool
	storage storage.StorageInterface
}

func New()*service{
	st := storage.New()
	return &service{
		id: uuid.New().String(),
		isMaster: true,
		storage: st,
	}
}

func (s *service)Put(ctx context.Context, key, value string) error{
	log.Println(s.id, "doing put...")
	if s.isMaster {
		return s.storage.Put(ctx, key, value)
	}
	return nil
}

func (s *service)Get(ctx context.Context, key string) (string, error){
	log.Println(s.id, "doing get...")
	if s.isMaster {
		return s.storage.Get(ctx, key)
	}
	return "",nil
}

func (s *service)Delete(ctx context.Context, key string) error{
	log.Println(s.id, "doing delete...")
	if s.isMaster {
		return s.storage.Delete(ctx, key)
	}
	return nil
}
