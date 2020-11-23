/**
 * Created by zc on 2020/6/9.
 */
package service

import (
	"context"
	"github.com/zc2638/gotool/utilx"
	"luban/pkg/database"
	"luban/pkg/database/data"
	"luban/pkg/errs"
	"luban/pkg/uuid"
)

type userService struct{ service }

func (s *userService) Find(ctx context.Context, name string) (*data.User, error) {
	var user data.User
	db := s.db.Where(&data.User{Username: name}).First(&user)
	if db.Error == nil {
		return &user, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, nil
	}
	return nil, db.Error
}

func (s *userService) FindByNameAndPwd(ctx context.Context, username, password string) (*data.User, error) {
	user, err := s.Find(ctx, username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errs.New("user not found")
	}
	if user.Pwd != password {
		return nil, errs.New("Invalid username or password")
	}
	return user, nil
}

func (s *userService) FindByUserID(ctx context.Context, userID string) (*data.User, error) {
	var user data.User
	db := s.db.Where(&data.User{UserID: userID}).First(&user)
	if db.Error == nil {
		return &user, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, errs.New("user not found")
	}
	return nil, db.Error
}

func (s *userService) Create(ctx context.Context, user *data.User) error {
	current, err := s.Find(ctx, user.Username)
	if err != nil {
		return err
	}
	// check if the user name is duplicate
	if current != nil {
		return errs.New("Duplicate username")
	}
	user.UserID = uuid.New()
	user.Salt = utilx.RandomStr(6)
	return s.db.Create(user).Error
}

func (s *userService) PwdReset(ctx context.Context, username, password string) error {
	current, err := s.Find(ctx, username)
	if err != nil {
		return err
	}
	if current == nil {
		return errs.New("user not found")
	}
	current.Pwd = password
	return s.db.Save(current).Error
}
