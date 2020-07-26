/**
 * Created by zc on 2020/7/18.
 */
package service

import (
	"context"
	"gopkg.in/yaml.v2"
	"luban/global"
	"luban/pkg/api/store"
	"luban/pkg/ctr"
	"luban/pkg/errs"
	"luban/pkg/storage"
	"luban/pkg/utils"
	"path/filepath"
	"strings"
)

type configService struct{}

func (s *configService) List(ctx context.Context) ([]store.Config, error) {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	space := ctr.ContextSpaceValue(ctx)
	path := filepath.Join(user.UserPath(), space)
	manifest, err := storage.New().Find(path, global.KeyManifest)
	if err != nil {
		return nil, err
	}
	var configs []store.Config
	if len(manifest) > 0 {
		err = yaml.Unmarshal(manifest, &configs)
	}
	return configs, err
}

func (s *configService) Find(ctx context.Context) (*store.Config, error) {
	list, err := s.List(ctx)
	if err != nil {
		return nil, err
	}
	config := ctr.ContextConfigValue(ctx)
	for _, v := range list {
		if v.Name == config {
			return &v, nil
		}
	}
	return nil, ErrNotExist
}

func (s *configService) Create(ctx context.Context, config *store.Config) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	list, err := s.List(ctx)
	if err != nil {
		return err
	}
	for _, v := range list {
		if v.Name == config.Name {
			return errs.Error("Duplicate config")
		}
	}
	list = append(list, *config)
	manifest, err := yaml.Marshal(&list)
	if err != nil {
		return err
	}
	space := ctr.ContextSpaceValue(ctx)
	path := filepath.Join(user.UserPath(), space)
	// 更新清单
	if err := storage.New().Update(path, global.KeyManifest, manifest); err != nil {
		return err
	}
	// 创建配置文件
	fp := filepath.Join(path, config.Name)
	return storage.New().Update(fp, global.KeyConfigDefault, []byte(config.Content))
}

func (s *configService) Update(ctx context.Context, config *store.Config) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	list, err := s.List(ctx)
	if err != nil {
		return err
	}
	target := ctr.ContextConfigValue(ctx)
	current := make([]store.Config, 0, len(list))
	for _, v := range list {
		if v.Name == target {
			v = *config
		}
		current = append(current, v)
	}
	manifest, err := yaml.Marshal(&current)
	if err != nil {
		return err
	}
	space := ctr.ContextSpaceValue(ctx)
	path := filepath.Join(user.UserPath(), space)
	// 更新清单
	if err := storage.New().Update(path, global.KeyManifest, manifest); err != nil {
		return err
	}
	keys, err := storage.New().PathKeys(path)
	if err != nil {
		return err
	}
	fp := filepath.Join(path, config.Name)
	if _, exist := utils.InStringSlice(keys, target); exist {
		targetPath := filepath.Join(path, target)
		if err := storage.New().PathUpdate(targetPath, fp); err != nil {
			return err
		}
	}
	// 创建/更新配置文件
	return storage.New().Update(fp, global.KeyConfigDefault, []byte(config.Content))
}

func (s *configService) Delete(ctx context.Context) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	list, err := s.List(ctx)
	if err != nil {
		return err
	}
	target := ctr.ContextConfigValue(ctx)
	current := make([]store.Config, 0, len(list))
	for _, v := range list {
		if v.Name == target {
			continue
		}
		current = append(current, v)
	}
	manifest, err := yaml.Marshal(&current)
	if err != nil {
		return err
	}
	space := ctr.ContextSpaceValue(ctx)
	path := filepath.Join(user.UserPath(), space)
	// 更新清单
	if err := storage.New().Update(path, global.KeyManifest, manifest); err != nil {
		return err
	}
	// 删除配置
	return storage.New().Delete(path, target)
}

func (s *configService) Raw(ctx context.Context, username, space, config string) ([]byte, error) {
	us := &userService{}
	user, err := us.Find(ctx, username)
	if err != nil {
		return nil, err
	}
	path := filepath.Join(global.PathData, user.Code, space, config)
	return storage.New().Find(path, global.KeyConfigDefault)
}

func (s *configService) getConfigPath(ctx context.Context) (string, error) {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return "", err
	}
	space := ctr.ContextSpaceValue(ctx)
	config := ctr.ContextConfigValue(ctx)
	path := filepath.Join(user.UserPath(), space, config)
	return path, nil
}

func (s *configService) VersionList(ctx context.Context) ([]store.ConfigVersion, error) {
	path, err := s.getConfigPath(ctx)
	if err != nil {
		return nil, err
	}
	keys, err := storage.New().Keys(path)
	if err != nil {
		return nil, err
	}
	list := make([]store.ConfigVersion, 0, len(keys))
	for _, key := range keys {
		ext := filepath.Ext(key)
		if ext == "" {
			continue
		}
		list = append(list, store.ConfigVersion{
			Version: strings.Trim(key, ext),
			Format:  strings.Trim(ext, "."),
		})
	}
	return list, nil
}

func (s *configService) VersionFind(ctx context.Context, name string) ([]byte, error) {
	path, err := s.getConfigPath(ctx)
	if err != nil {
		return nil, err
	}
	return storage.New().Find(path, name)
}

func (s *configService) VersionCreate(ctx context.Context, version string) error {
	info, err := s.Find(ctx)
	if err != nil {
		return err
	}
	path, err := s.getConfigPath(ctx)
	if err != nil {
		return err
	}
	keys, err := storage.New().Keys(path)
	if err != nil {
		return err
	}
	for _, key := range keys {
		ext := filepath.Ext(key)
		keyVersion := strings.Trim(key, ext)
		if keyVersion == version {
			return ErrExist
		}
	}
	name := version + "." + info.Format
	return storage.New().Create(path, name, []byte(info.Content))
}

func (s *configService) VersionDelete(ctx context.Context, name string) error {
	path, err := s.getConfigPath(ctx)
	if err != nil {
		return err
	}
	return storage.New().Delete(path, name)
}

func (s *configService) VersionRaw(ctx context.Context, username, space, config, version string) ([]byte, error) {
	us := &userService{}
	user, err := us.Find(ctx, username)
	if err != nil {
		return nil, err
	}
	path := filepath.Join(global.PathData, user.Code, space, config)
	keys, err := storage.New().Keys(path)
	if err != nil {
		return nil, err
	}
	for _, key := range keys {
		ext := filepath.Ext(key)
		keyVersion := strings.Trim(key, ext)
		if keyVersion == version {
			return storage.New().Find(path, key)
		}
	}
	return nil, nil
}
