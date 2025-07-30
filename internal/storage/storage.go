package storage

import "context"

type storageServise struct {
	//TODO:
}

func New()*storageServise{
	return &storageServise{
		
	}
}
func (s *storageServise)Put(ctx context.Context, key, value string) error{
	// TODO
	return  nil
}
func (s *storageServise)Get(ctx context.Context, key string) (string, error){
	// TODO
	return "", nil
}
func (s *storageServise)Delete(ctx context.Context, key string) error{
	// TODO
	return  nil
}