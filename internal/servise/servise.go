package servise

import (
	"context"
	"log"

	"github.com/DKeshavarz/armis/internal/storage"
	"github.com/google/uuid"
)

type service struct {
	id string
	masterID string
	storage storage.StorageInterface
}

func New(st storage.StorageInterface)*service{
	tmp_id := uuid.New().String()
	return &service{
		id: tmp_id,
		masterID: tmp_id,
		storage: st,
	}
}

func (s *service)Put(ctx context.Context, key, value string) error{
	log.Println(s.id, "doing put...")
	if s.masterID == s.id {
		return s.storage.Put(ctx, key, value)
	}
	return nil
}

func (s *service)Get(ctx context.Context, key string) (string, error){
	log.Println(s.id, "doing get...")
	
	return s.storage.Get(ctx, key)
}

func (s *service)Delete(ctx context.Context, key string) error{
	log.Println(s.id, "doing delete...")
	if s.masterID == s.id  {
		return s.storage.Delete(ctx, key)
	}
	return nil
}
