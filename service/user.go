/**
 * Created by zc on 2020/6/9.
 */
package service

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/zc2638/gotool/utilx"
	"stone/global"
	"stone/pkg/api/store"
	"stone/pkg/errs"
)

type userService struct{}

func (s *userService) FindByEmail(ctx context.Context, email string) (*store.User, error) {
	var user store.User
	db := global.DB().Where(store.User{Email: email}).First(&user)
	if db.Error != nil {
		return nil, db.Error
	}
	return &user, nil
}

func (s *userService) FindByNameAndPwd(ctx context.Context, username, password string) (*store.User, error) {
	user, err := s.FindByEmail(ctx, username)
	if err != nil {
		return nil, errs.Error("Invalid username or password").With(err)
	}
	h := md5.New()
	h.Write([]byte(password))
	hb := h.Sum([]byte(user.Salt))
	pwd := base64.StdEncoding.EncodeToString(hb)
	if user.Pwd != pwd {
		return nil, errs.New("Invalid username or password")
	}
	return user, nil
}

func (s *userService) Create(ctx context.Context, user *store.User) error {
	// check if the user name is duplicate
	user, err := s.FindByEmail(ctx, user.Email)
	if err != nil {
		return err
	}
	if user.UID != "" {
		return errs.New("Duplicate email")
	}
	user.UID = uuid.New().String()
	user.Salt = utilx.RandomStr(6)
	h := md5.New()
	h.Write([]byte(user.Pwd))
	hb := h.Sum([]byte(user.Salt))
	user.Pwd = base64.StdEncoding.EncodeToString(hb)
	return global.DB().Create(user).Error
}
