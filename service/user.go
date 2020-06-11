/**
 * Created by zc on 2020/6/9.
 */
package service

import (
	"crypto/md5"
	"encoding/base64"
	"github.com/google/uuid"
	"github.com/zc2638/gotool/utilx"
	"stone/global"
	"stone/pkg/api/store"
	"stone/pkg/errs"
)

type userService struct{}

func (s *userService) FindByEmail(email string) (*store.User, bool) {
	var user store.User
	global.DB().Where(store.User{Email: email}).First(&user)
	if user.UID == "" {
		return nil, false
	}
	return &user, true
}

func (s *userService) FindByNameAndPwd(username, password string) (*store.User, error) {
	user, ok := s.FindByEmail(username)
	if !ok {
		return nil, errs.New("Invalid username or password")
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

func (s *userService) Create(user *store.User) error {
	user.UID = uuid.New().String()
	user.Salt = utilx.RandomStr(6)
	h := md5.New()
	h.Write([]byte(user.Pwd))
	hb := h.Sum([]byte(user.Salt))
	user.Pwd = base64.StdEncoding.EncodeToString(hb)
	return global.DB().Create(user).Error
}
