/**
 * Created by zc on 2020/6/9.
 */
package service

import (
	"bytes"
	"context"
	"luban/global"
	"luban/pkg/api/store"
	"luban/pkg/errs"
	"luban/pkg/fs"
	"luban/pkg/uuid"
	"strings"
)

type userService struct{}

func (s *userService) all() ([]store.User, error) {
	ub, err := fs.New().File().Find(global.FilePathUser)
	if err != nil {
		return nil, err
	}
	us := string(ub)
	arr := strings.Split(us, "\n")
	users := make([]store.User, 0, len(arr))
	for _, v := range arr {
		infos := strings.Split(v, ",")
		if len(infos) != 3 {
			continue
		}
		users = append(users, store.User{
			Code:     infos[0],
			Username: infos[1],
			Pwd:      infos[2],
		})
	}
	return users, nil
}

func (s *userService) FindStatus(ctx context.Context, name string) (user *store.User, exist bool, err error) {
	users, err := s.all()
	if err != nil {
		return nil, false, err
	}
	for _, u := range users {
		if u.Username == name {
			return &u, true, nil
		}
	}
	return &store.User{}, false, nil
}

func (s *userService) Find(ctx context.Context, name string) (*store.User, error) {
	user, exist, err := s.FindStatus(ctx, name)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, errs.Error("user not found")
	}
	return user, nil
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
	users, err := s.all()
	if err != nil {
		return err
	}
	// check if the user name is duplicate
	_, exist, err := s.FindStatus(ctx, user.Username)
	if err != nil {
		return err
	}
	if exist {
		return errs.New("Duplicate username")
	}
	user.Code = uuid.New()
	users = append(users, *user)
	var buf bytes.Buffer
	for _, u := range users {
		buf.WriteString(u.Code)
		buf.WriteString(",")
		buf.WriteString(u.Username)
		buf.WriteString(",")
		buf.WriteString(u.Pwd)
		buf.WriteString("\n")
	}
	return fs.New().File().Update(global.FilePathUser, buf.Bytes())
}
