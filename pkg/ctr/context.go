/**
 * Created by zc on 2020/6/10.
 */
package ctr

import (
	"context"
)

const ContextUserKey = "user"

// WithUser returns a copy of parent in which the user value is set
func ContextWithUser(parent context.Context, user *JwtUserInfo) context.Context {
	return context.WithValue(parent, ContextUserKey, user)
}

// UserFrom returns the value of the user key on the ctx
func ContextUserFrom(ctx context.Context) (*JwtUserInfo, bool) {
	user, ok := ctx.Value(ContextUserKey).(*JwtUserInfo)
	return user, ok
}
