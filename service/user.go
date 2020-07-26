/**
 * Created by zc on 2020/6/9.
 */
package service

import (
	"context"
	"gopkg.in/yaml.v2"
	"luban/global"
	"luban/pkg/api/store"
	"luban/pkg/errs"
	"luban/pkg/storage"
	"luban/pkg/uuid"
)

type userService struct{}

func (s *userService) All(ctx context.Context) ([]store.User, error) {
	ub, err := storage.New().Find(global.PathRoot, global.KeyUserFile)
	if err != nil {
		return nil, err
	}
	var users []store.User
	err = yaml.Unmarshal(ub, &users)
	return users, err
}

func (s *userService) Find(ctx context.Context, name string) (*store.User, error) {
	users, err := s.All(ctx)
	if err != nil {
		return nil, err
	}
	for _, u := range users {
		if u.Username == name {
			return &u, nil
		}
	}
	return nil, ErrNotExist
}

func (s *userService) FindByNameAndPwd(ctx context.Context, username, password string) (*store.User, error) {
	user, err := s.Find(ctx, username)
	if err != nil {
		return nil, err
	}
	if user.Pwd != password {
		return nil, errs.New("Invalid username or password")
	}
	return user, nil
}

func (s *userService) Create(ctx context.Context, user *store.User) error {
	users, err := s.All(ctx)
	if err != nil {
		return err
	}
	// check if the user name is duplicate
	for _, u := range users {
		if u.Username == user.Username {
			return errs.New("Duplicate username")
		}
	}
	user.Code = uuid.New()
	users = append(users, *user)
	b, err := yaml.Marshal(&users)
	if err != nil {
		return err
	}
	return storage.New().Update(global.PathRoot, global.KeyUserFile, b)
}
