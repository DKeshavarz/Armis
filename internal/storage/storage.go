package storage

import (
	"context"
	"sync"
	"time"
)

type storage struct {
	mutex        sync.RWMutex
	data         map[string]string
	autoSave     bool
	saveInterval time.Duration
}

func New() *storage {
	return &storage{
		data: map[string]string{},
	}
}
func (s *storage) Put(ctx context.Context, key, value string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.data[key] = value
	return nil
}
func (s *storage) Get(ctx context.Context, key string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	val, ok := s.data[key]
	if !ok {
		return "moew", ErrKeyNotFound
	}
	return val, nil
}
func (s *storage) Delete(ctx context.Context, key string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.data, key)
	return nil
}
