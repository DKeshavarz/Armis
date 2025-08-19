package storage

import (
	"context"
	"encoding/json"
	"log"
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

func New(autoSave bool, saveInterval int, filePath string) *storage {
	s :=  &storage{
		data: map[string]string{},
		autoSave: autoSave,
		saveInterval: time.Duration(saveInterval) * time.Second,
		filePath: filePath,
	}

	if err := s.load() ; err != nil {
		log.Printf("err in reading file in package storage %s", err.Error())
	}

	if s.autoSave {
		go s.autosaver()
	}
	return s
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
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	if s.filePath == "" {
		return ErrPathNotSet
	}

	data, err := json.Marshal(s.data)
	if err != nil {
		return err
	}
	
	return  os.WriteFile(s.filePath, data, 0777)
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

func (s *storage) autosaver() {
	ticker := time.NewTicker(s.saveInterval)

	for range ticker.C {
		if err := s.save(); err != nil {
			log.Printf("get error in saving in package storage : %s\n", err.Error())
		}
	}
}
