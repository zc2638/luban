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
	"time"
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
	fp := filepath.Join(path, config.Name)
	// 创建/更新配置文件
	if err := storage.New().Update(fp, global.KeyConfigDefault, []byte(config.Content)); err != nil {
		return err
	}
	if target != config.Name {
		keys, err := storage.New().PathKeys(path)
		if err != nil {
			return err
		}
		if _, exist := utils.InStringSlice(keys, target); exist {
			targetPath := filepath.Join(path, target)
			if err := storage.New().PathUpdate(targetPath, fp); err != nil {
				return err
			}
		}
	}
	// 更新清单
	return storage.New().Update(path, global.KeyManifest, manifest)
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
	manifest, err := storage.New().Find(path, global.KeyManifest)
	if err != nil {
		return nil, err
	}
	var verGroup []store.ConfigVersion
	if len(manifest) > 0 {
		if err := yaml.Unmarshal(manifest, &verGroup); err != nil {
			return nil, err
		}
	}
	return verGroup, nil
}

func (s *configService) VersionFind(ctx context.Context, version string) ([]byte, error) {
	path, err := s.getConfigPath(ctx)
	if err != nil {
		return nil, err
	}
	return storage.New().Find(path, version+global.KeyConfigVersionExt)
}

func (s *configService) VersionCreate(ctx context.Context, version, desc string) error {
	info, err := s.Find(ctx)
	if err != nil {
		return err
	}
	path, err := s.getConfigPath(ctx)
	if err != nil {
		return err
	}
	vs, err := s.VersionList(ctx)
	if err != nil {
		return err
	}
	list := make([]store.ConfigVersion, 0, len(vs)+1)
	list = append(list, store.ConfigVersion{
		Version:   version,
		Format:    info.Format,
		Desc:      desc,
		CreatedAt: time.Now().Unix(),
	})
	for _, v := range vs {
		if v.Version == version {
			return ErrExist
		}
		list = append(list, v)
	}
	if err := storage.New().Create(path, version+global.KeyConfigVersionExt, []byte(info.Content)); err != nil {
		return err
	}
	manifest, err := yaml.Marshal(&list)
	if err != nil {
		return err
	}
	return storage.New().Update(path, global.KeyManifest, manifest)
}

func (s *configService) VersionDelete(ctx context.Context, version string) error {
	path, err := s.getConfigPath(ctx)
	if err != nil {
		return err
	}
	vs, err := s.VersionList(ctx)
	if err != nil {
		return err
	}
	list := make([]store.ConfigVersion, 0, len(vs))
	for _, v := range vs {
		if v.Version == version {
			continue
		}
		list = append(list, v)
	}
	manifest, err := yaml.Marshal(&list)
	if err != nil {
		return err
	}
	if err := storage.New().Update(path, global.KeyManifest, manifest); err != nil {
		return err
	}
	return storage.New().Delete(path, version+global.KeyConfigVersionExt)
}

func (s *configService) VersionRaw(ctx context.Context, username, space, config, version string) ([]byte, error) {
	us := &userService{}
	user, err := us.Find(ctx, username)
	if err != nil {
		return nil, err
	}
	path := filepath.Join(global.PathData, user.Code, space, config)
	return storage.New().Find(path, version+global.KeyConfigVersionExt)
}
