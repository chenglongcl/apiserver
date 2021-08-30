package esservice

import (
	"apiserver/elastic/esdao"
	"apiserver/elastic/esmodel"
	"context"
)

type UserService struct {
	es *esdao.UserES
}

func NewUserService(es *esdao.UserES) *UserService {
	return &UserService{
		es: es,
	}
}

func (s *UserService) BatchAdd(ctx context.Context, user []*esmodel.UserEs) error {
	return s.es.BatchAdd(ctx, user)
}

func (s *UserService) BatchUpdate(ctx context.Context, user []*esmodel.UserEs) error {
	return s.es.BatchUpdate(ctx, user)
}

func (s *UserService) BatchDel(ctx context.Context, user []*esmodel.UserEs) error {
	return s.es.BatchDel(ctx, user)
}

func (s *UserService) MGet(ctx context.Context, IDS []uint64) ([]*esmodel.UserEs, error) {
	return s.es.MGet(ctx, IDS)
}

func (s *UserService) Get(ctx context.Context, id uint64) (*esmodel.UserEs, error) {
	return s.es.Get(ctx, id)
}

func (s *UserService) Search(ctx context.Context, req *esmodel.SearchUserRequest) ([]*esmodel.UserEs, uint64, error) {
	return s.es.Search(ctx, req.ToFilter())
}
