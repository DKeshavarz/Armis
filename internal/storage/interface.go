package storage

import "context"

type storageInterface interface {
	Put(ctx context.Context, key, value string) error
    Get(ctx context.Context, key string) (string, error)
    Delete(ctx context.Context, key string) error
}