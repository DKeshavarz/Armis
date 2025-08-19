package storage

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrPathNotSet  = errors.New("No path set for saving")
)
