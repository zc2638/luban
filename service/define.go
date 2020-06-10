/**
 * Created by zc on 2020/6/9.
 */
package service

import (
	"stone/pkg/api/store"
)

type Interface interface {
	User() UserService
}

type UserService interface {
	FindByEmail(email string) (*store.User, bool)
	FindByNameAndPwd(username, password string) (*store.User, error)
	Create(user *store.User) error
}
