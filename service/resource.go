/**
 * Created by zc on 2020/7/18.
 */
package service

import (
	"context"
	"github.com/jinzhu/gorm"
	"luban/global"
	"luban/pkg/ctr"
	"luban/pkg/database/data"
	"luban/pkg/uuid"
)

type resourceService struct{}

func (s *resourceService) List(ctx context.Context) ([]data.Resource, error) {
	space := ctr.ContextSpaceValue(ctx)
	var resources []data.Resource
	db := global.DB().Where(&data.Resource{SpaceID: space}).Find(&resources)
	if db.Error == nil || gorm.IsRecordNotFoundError(db.Error) {
		return resources, nil
	}
	return nil, db.Error
}

func (s *resourceService) Find(ctx context.Context) (*data.Resource, error) {
	space := ctr.ContextSpaceValue(ctx)
	target := ctr.ContextResourceValue(ctx)
	var resource data.Resource
	db := global.DB().Where(&data.Resource{
		SpaceID:    space,
		ResourceID: target,
	}).First(&resource)
	if db.Error == nil {
		return &resource, nil
	}
	if gorm.IsRecordNotFoundError(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *resourceService) FindByName(ctx context.Context, name string) (*data.Resource, error) {
	space := ctr.ContextSpaceValue(ctx)
	var resource data.Resource
	db := global.DB().Where(&data.Resource{
		SpaceID: space,
		Name:    name,
	}).First(&resource)
	if db.Error == nil {
		return &resource, nil
	}
	if gorm.IsRecordNotFoundError(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *resourceService) Create(ctx context.Context, resource *data.Resource) error {
	space := ctr.ContextSpaceValue(ctx)
	if _, err := s.FindByName(ctx, resource.Name); err == nil {
		return ErrExist
	}
	resource.ResourceID = uuid.New()
	resource.SpaceID = space
	return global.DB().Create(resource).Error
}

func (s *resourceService) Update(ctx context.Context, resource *data.Resource) error {
	current, err := s.Find(ctx)
	if err != nil {
		return err
	}
	resource.ResourceID = current.ResourceID
	resource.SpaceID = current.SpaceID
	resource.CreatedAt = current.CreatedAt
	return global.DB().Save(resource).Error
}

func (s *resourceService) Delete(ctx context.Context) error {
	space := ctr.ContextSpaceValue(ctx)
	target := ctr.ContextResourceValue(ctx)
	return global.DB().Where(&data.Resource{
		ResourceID: target,
		SpaceID:    space,
	}).Delete(&data.Resource{}).Error
}

func (s *resourceService) Raw(ctx context.Context, username, space, resource string) ([]byte, error) {
	us := &userService{}
	user, err := us.Find(ctx, username)
	if err != nil {
		return nil, err
	}
	ctx = ctr.ContextWithUser(ctx, &ctr.JwtUserInfo{
		UserID: user.UserID,
		Pwd:    user.Pwd,
	})
	ss := &spaceService{}
	spaceData, err := ss.FindByName(ctx, space)
	if err != nil {
		return nil, err
	}
	ctx = ctr.ContextWithSpace(ctx, spaceData.SpaceID)
	resourceData, err := s.FindByName(ctx, resource)
	if err != nil {
		return nil, err
	}
	return []byte(resourceData.Content), nil
}

func (s *resourceService) VersionList(ctx context.Context) ([]data.Version, error) {
	resource := ctr.ContextResourceValue(ctx)
	var versions []data.Version
	db := global.DB().Where(&data.Version{
		ResourceID: resource,
	}).First(&versions)
	if db.Error == nil || gorm.IsRecordNotFoundError(db.Error) {
		return versions, nil
	}
	return nil, db.Error
}

func (s *resourceService) VersionFind(ctx context.Context, id string) (*data.Version, error) {
	resource := ctr.ContextResourceValue(ctx)
	var version data.Version
	db := global.DB().Where(&data.Version{
		ResourceID: resource,
		VersionID:  id,
	}).First(&version)
	if db.Error == nil {
		return &version, nil
	}
	if gorm.IsRecordNotFoundError(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *resourceService) VersionFindByName(ctx context.Context, name string) (*data.Version, error) {
	resource := ctr.ContextResourceValue(ctx)
	var version data.Version
	db := global.DB().Where(&data.Version{
		ResourceID: resource,
		Version:    name,
	}).First(&version)
	if db.Error == nil {
		return &version, nil
	}
	if gorm.IsRecordNotFoundError(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *resourceService) VersionCreate(ctx context.Context, version *data.Version) error {
	if _, err := s.FindByName(ctx, version.Version); err == nil {
		return ErrExist
	}
	resource, err := s.Find(ctx)
	if err != nil {
		return err
	}
	version.VersionID = uuid.New()
	version.ResourceID = resource.ResourceID
	version.Format = resource.Format
	version.Content = resource.Content
	return global.DB().Create(version).Error
}

func (s *resourceService) VersionDelete(ctx context.Context, id string) error {
	resource := ctr.ContextResourceValue(ctx)
	return global.DB().Where(&data.Version{
		ResourceID: resource,
		VersionID:  id,
	}).Delete(&data.Version{}).Error
}

func (s *resourceService) VersionRaw(ctx context.Context, username, space, resource, version string) ([]byte, error) {
	us := &userService{}
	user, err := us.Find(ctx, username)
	if err != nil {
		return nil, err
	}
	ctx = ctr.ContextWithUser(ctx, &ctr.JwtUserInfo{
		UserID: user.UserID,
		Pwd:    user.Pwd,
	})
	ss := &spaceService{}
	spaceData, err := ss.FindByName(ctx, space)
	if err != nil {
		return nil, err
	}
	ctx = ctr.ContextWithSpace(ctx, spaceData.SpaceID)
	resourceData, err := s.FindByName(ctx, resource)
	if err != nil {
		return nil, err
	}
	ctx = ctr.ContextWithResource(ctx, resourceData.ResourceID)
	versionData, err := s.VersionFindByName(ctx, version)
	if err != nil {
		return nil, err
	}
	return []byte(versionData.Content), nil
}
