/**
 * Created by zc on 2020/6/10.
 */
package ctr

import (
	"context"
	"stone/global"
)

// WithUser returns a copy of parent in which the user value is set
func WithUser(parent context.Context, user *JwtUserInfo) context.Context {
	return context.WithValue(parent, global.ContextUserKey, user)
}

// UserFrom returns the value of the user key on the ctx
func UserFrom(ctx context.Context) (*JwtUserInfo, bool) {
	user, ok := ctx.Value(global.ContextUserKey).(*JwtUserInfo)
	return user, ok
}
