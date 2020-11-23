/**
 * Created by zc on 2020/6/10.
 */
package wrap

import (
	"context"
	"luban/pkg/errs"
)

type key int

const (
	ContextUserKey key = iota
	ContextSpaceKey
	ContextResourceKey
	ContextPipelineKey
	ContextTaskKey
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

// ContextWithResource returns a copy of parent in which the resource value is set
func ContextWithResource(parent context.Context, resource string) context.Context {
	return context.WithValue(parent, ContextResourceKey, resource)
}

// ContextResourceValue returns the value of the resource key on the ctx
func ContextResourceValue(ctx context.Context) string {
	return ctx.Value(ContextResourceKey).(string)
}

// ContextWithPipeline returns a copy of parent in which the pipeline value is set
func ContextWithPipeline(parent context.Context, pipeline string) context.Context {
	return context.WithValue(parent, ContextPipelineKey, pipeline)
}

// ContextPipelineValue returns the value of the pipeline key on the ctx
func ContextPipelineValue(ctx context.Context) string {
	return ctx.Value(ContextPipelineKey).(string)
}

// ContextWithTask returns a copy of parent in which the task value is set
func ContextWithTask(parent context.Context, task string) context.Context {
	return context.WithValue(parent, ContextTaskKey, task)
}

// ContextTaskValue returns the value of the task key on the ctx
func ContextTaskValue(ctx context.Context) string {
	return ctx.Value(ContextTaskKey).(string)
}
