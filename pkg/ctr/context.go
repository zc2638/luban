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
	ContextSpaceKey
	ContextResourceKey
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

// ContextWithSpace returns a copy of parent in which the space value is set
func ContextWithSpace(parent context.Context, space string) context.Context {
	return context.WithValue(parent, ContextSpaceKey, space)
}

// ContextSpaceValue returns the value of the space key on the ctx
func ContextSpaceValue(ctx context.Context) string {
	return ctx.Value(ContextSpaceKey).(string)
}

// ContextWithResource returns a copy of parent in which the config value is set
func ContextWithResource(parent context.Context, config string) context.Context {
	return context.WithValue(parent, ContextResourceKey, config)
}

// ContextResourceValue returns the value of the config key on the ctx
func ContextResourceValue(ctx context.Context) string {
	return ctx.Value(ContextResourceKey).(string)
}
