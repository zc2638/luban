/**
 * Created by zc on 2020/6/10.
 */
package service

import (
	"context"
	"luban/pkg/api/store"
)

type Interface interface {
	User() UserService
	Space() SpaceService
}

type UserService interface {
	FindByNameAndPwd(ctx context.Context, username, password string) (*store.User, error)
	Create(ctx context.Context, user *store.User) error
}

type SpaceService interface {
	List(ctx context.Context) ([]string, error)
	Create(ctx context.Context, name string) error
	Update(ctx context.Context, target string, name string) error
	Delete(ctx context.Context, name string) error
}

type Service struct{}

func New() Interface {
	return &Service{}
}

func (s *Service) User() UserService {
	return &userService{}
}

func (s *Service) Space() SpaceService {
	return &spaceService{}
}
