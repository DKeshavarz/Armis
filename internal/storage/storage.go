package storage

import (
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"
)

type storage struct {
	mutex        sync.RWMutex
	data         map[string]string
	autoSave     bool
	saveInterval time.Duration
	filePath     string
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

func (s *storage) save() error{
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.filePath == "" {
		return ErrPathNotSet
	}

	data, err := json.Marshal(s.data)
	if err != nil {
		return err
	}
	
	return  os.WriteFile(s.filePath, data, 0664)
}

func (s *storage) load() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.filePath == "" {
		return ErrPathNotSet
	}

	data, err := os.ReadFile(s.filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNoFileExist
		}
		return err
	}

	return json.Unmarshal(data, &s.data)
}
