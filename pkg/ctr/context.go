/**
 * Created by zc on 2020/6/10.
 */
package ctr

import (
	"context"
	"luban/pkg/errs"
)

type key int

const (
	ContextUserKey key = iota
)

// ContextWithUser returns a copy of parent in which the user value is set
func ContextWithUser(parent context.Context, user *JwtUserInfo) context.Context {
	return context.WithValue(parent, ContextUserKey, user)
}

// ContextUserFrom returns the value of the user key on the ctx
func ContextUserFrom(ctx context.Context) (*JwtUserInfo, error) {
	user, ok := ctx.Value(ContextUserKey).(*JwtUserInfo)
	if !ok {
		return nil, errs.ErrUnauthorized
	}
	return user, nil
}
