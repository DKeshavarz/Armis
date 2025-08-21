package storage

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrPathNotSet  = errors.New("no path set for saving")
	ErrNoFileExist = errors.New("no file exist")
)
