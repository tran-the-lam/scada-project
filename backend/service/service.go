package service

import "context"

type IService interface {
	PutState(ctx context.Context, key string, val any) error
	GetState(ctx context.Context, key string) (any, error)
}

type service struct{}

func NewService() IService {
	return &service{}
}

func (s *service) PutState(ctx context.Context, key string, val any) error {
	return nil
}

func (s *service) GetState(ctx context.Context, key string) (any, error) {
	return nil, nil
}
