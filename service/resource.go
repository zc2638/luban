/**
 * Created by zc on 2020/7/18.
 */
package service

import (
	"context"
	"luban/global/database"
	"luban/pkg/store"
	"luban/pkg/util"
	"luban/pkg/wrapper"
)

type resourceService struct{ service }

func (s *resourceService) List(ctx context.Context, kind string) ([]store.Resource, error) {
	user, err := wrapper.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	resources := make([]store.Resource, 0)
	db := s.db.Where(&store.Resource{
		UserID: user.UserID,
		Kind:   kind,
	}).Order("created_at desc").Find(&resources)
	if db.Error != nil && !database.RecordNotFound(db.Error) {
		return nil, db.Error
	}
	return resources, nil
}

func (s *resourceService) FindByID(ctx context.Context, resourceID string) (*store.Resource, error) {
	user, err := wrapper.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	var resource store.Resource
	db := s.db.Where(&store.Resource{
		ResourceID: resourceID,
		UserID:     user.UserID,
	}).First(&resource)
	if db.Error == nil {
		return &resource, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *resourceService) Find(ctx context.Context, name string) (*store.Resource, error) {
	user, err := wrapper.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	var resource store.Resource
	db := s.db.Where(&store.Resource{
		Name:   name,
		UserID: user.UserID,
	}).First(&resource)
	if db.Error == nil {
		return &resource, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *resourceService) Create(ctx context.Context, resource *store.Resource) error {
	if _, err := s.Find(ctx, resource.Name); err == nil {
		return ErrExist
	}
	resource.ResourceID = util.UUID()
	return s.db.Create(resource).Error
}

func (s *resourceService) Update(ctx context.Context, resource *store.Resource) error {
	current, err := s.Find(ctx, resource.Name)
	if err != nil {
		return err
	}
	resource.ResourceID = current.ResourceID
	resource.SpaceID = current.SpaceID
	resource.CreatedAt = current.CreatedAt
	return s.db.Save(resource).Error
}

func (s *resourceService) Delete(ctx context.Context, name string) error {
	user, err := wrapper.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	return s.db.Where(&store.Resource{
		Name:   name,
		UserID: user.UserID,
	}).Delete(&store.Resource{}).Error
}

func (s *resourceService) Raw(ctx context.Context, username, resource string) ([]byte, error) {
	user, err := New().User().Find(ctx, username)
	if err != nil {
		return nil, err
	}
	ctx = wrapper.ContextWithUser(ctx, &wrapper.JwtUserInfo{
		UserID: user.UserID,
		Pwd:    user.Pwd,
	})
	resourceData, err := s.Find(ctx, resource)
	if err != nil {
		return nil, err
	}
	return []byte(resourceData.Data), nil
}

func (s *resourceService) VersionList(ctx context.Context, resource string) ([]store.Version, error) {
	var versions []store.Version
	db := s.db.Where(&store.Version{
		ResourceID: resource,
	}).First(&versions)
	if db.Error == nil || database.RecordNotFound(db.Error) {
		return versions, nil
	}
	return nil, db.Error
}

func (s *resourceService) VersionFind(ctx context.Context, resource, ver string) (*store.Version, error) {
	rs, err := s.Find(ctx, resource)
	if err != nil {
		return nil, err
	}
	var version store.Version
	db := s.db.Where(&store.Version{
		ResourceID: rs.ResourceID,
		Version:    ver,
	}).First(&version)
	if db.Error == nil {
		return &version, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *resourceService) VersionDelete(ctx context.Context, resource, ver string) error {
	return s.db.Where(&store.Version{
		ResourceID: resource,
		Version:    ver,
	}).Delete(&store.Version{}).Error
}

func (s *resourceService) VersionRaw(ctx context.Context, username, resource, ver string) ([]byte, error) {
	user, err := New().User().Find(ctx, username)
	if err != nil {
		return nil, err
	}
	ctx = wrapper.ContextWithUser(ctx, &wrapper.JwtUserInfo{
		UserID: user.UserID,
		Pwd:    user.Pwd,
	})
	versionData, err := s.VersionFind(ctx, resource, ver)
	if err != nil {
		return nil, err
	}
	return []byte(versionData.Data), nil
}
