/**
 * Created by zc on 2020/6/11.
 */
package service

import (
	"context"
	"gopkg.in/yaml.v2"
	"luban/global"
	"luban/pkg/api/store"
	"luban/pkg/ctr"
	"luban/pkg/storage"
	"luban/pkg/utils"
	"path/filepath"
)

type spaceService struct{}

func (s *spaceService) List(ctx context.Context) ([]store.Space, error) {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	data, err := storage.New().Find(user.UserPath(), global.KeyManifest)
	if err != nil {
		return nil, err
	}
	var spaces []store.Space
	if len(data) == 0 {
		return spaces, nil
	}
	err = yaml.Unmarshal(data, &spaces)
	return spaces, err
}

func (s *spaceService) Find(ctx context.Context, name string) (*store.Space, error) {
	spaces, err := s.List(ctx)
	if err != nil {
		return nil, err
	}
	for _, space := range spaces {
		if space.Name == name {
			return &space, nil
		}
	}
	return nil, ErrNotExist
}

func (s *spaceService) Create(ctx context.Context, name string) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	spaces, err := s.List(ctx)
	if err != nil {
		return err
	}
	for _, space := range spaces {
		if space.Name == name {
			return ErrExist
		}
	}
	spaces = append(spaces, store.Space{
		Name: name,
	})
	b, err := yaml.Marshal(&spaces)
	if err != nil {
		return err
	}
	return storage.New().Update(user.UserPath(), global.KeyManifest, b)
}

func (s *spaceService) Update(ctx context.Context, name string) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	spaces, err := s.List(ctx)
	if err != nil {
		return err
	}
	target := ctr.ContextSpaceValue(ctx)
	list := make([]store.Space, 0, len(spaces))
	for _, space := range spaces {
		if space.Name == target {
			space.Name = name
		}
		list = append(list, space)
	}
	b, err := yaml.Marshal(&list)
	if err != nil {
		return err
	}
	if err := storage.New().Update(user.UserPath(), global.KeyManifest, b); err != nil {
		return err
	}
	keys, err := storage.New().PathKeys(user.UserPath())
	if err != nil {
		return err
	}
	if _, exist := utils.InStringSlice(keys, target); !exist {
		return nil
	}
	targetPath := filepath.Join(user.UserPath(), target)
	newPath := filepath.Join(user.UserPath(), name)
	return storage.New().PathUpdate(targetPath, newPath)
}

func (s *spaceService) Delete(ctx context.Context) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	spaces, err := s.List(ctx)
	if err != nil {
		return err
	}
	target := ctr.ContextSpaceValue(ctx)
	list := make([]store.Space, 0, len(spaces))
	for _, space := range spaces {
		if space.Name == target {
			continue
		}
		list = append(list, space)
	}
	b, err := yaml.Marshal(&list)
	if err != nil {
		return err
	}
	if err := storage.New().Update(user.UserPath(), global.KeyManifest, b); err != nil {
		return err
	}
	return storage.New().Delete(user.UserPath(), target)
}
