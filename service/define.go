/**
 * Created by zc on 2020/6/9.
 */
package service

import (
	"context"
	"stone/pkg/api/store"
)

type Interface interface {
	User() UserService
	Space() SpaceService
}

type UserService interface {
	FindByEmail(ctx context.Context, email string) (*store.User, error)
	FindByNameAndPwd(ctx context.Context, username, password string) (*store.User, error)
	Create(ctx context.Context, user *store.User) error
}

type SpaceService interface {
	List(ctx context.Context) ([]*store.Space, error)
	Find(ctx context.Context, id string) (*store.Space, error)
	Create(ctx context.Context, space *store.Space) error
	Update(ctx context.Context, space *store.Space) error
	Delete(ctx context.Context, id string) error
}
