/**
 * Created by zc on 2020/6/11.
 */
package service

import (
	"context"
	"luban/pkg/ctr"
	"luban/pkg/fs"
	"path/filepath"
)

type spaceService struct{}

func (s *spaceService) List(ctx context.Context) ([]string, error) {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	return fs.New().Dir().List(user.UserPath())
}

func (s *spaceService) Create(ctx context.Context, name string) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	path := filepath.Join(user.UserPath(), name)
	return fs.New().Dir().Create(path)
}

func (s *spaceService) Update(ctx context.Context, target, name string) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	userPath := user.UserPath()
	targetPath := filepath.Join(userPath, target)
	path := filepath.Join(userPath, name)
	return fs.New().Dir().Update(targetPath, path)
}

func (s *spaceService) Delete(ctx context.Context, name string) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	path := filepath.Join(user.UserPath(), name)
	return fs.New().Dir().Delete(path)
}
