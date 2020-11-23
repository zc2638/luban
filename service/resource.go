/**
 * Created by zc on 2020/7/18.
 */
package service

import (
	"context"
	"luban/global/database"
	"luban/pkg/store"
	"luban/pkg/util"
	"luban/pkg/wrap"
)

type resourceService struct{ service }

func (s *resourceService) List(ctx context.Context) ([]store.Resource, error) {
	space := wrap.ContextSpaceValue(ctx)
	var resources []store.Resource
	db := s.db.Where(&store.Resource{SpaceID: space}).Find(&resources)
	if db.Error == nil || database.RecordNotFound(db.Error) {
		return resources, nil
	}
	return nil, db.Error
}

func (s *resourceService) Find(ctx context.Context) (*store.Resource, error) {
	space := wrap.ContextSpaceValue(ctx)
	target := wrap.ContextResourceValue(ctx)
	var resource store.Resource
	db := s.db.Where(&store.Resource{
		SpaceID:    space,
		ResourceID: target,
	}).First(&resource)
	if db.Error == nil {
		return &resource, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *resourceService) FindByName(ctx context.Context, name string) (*store.Resource, error) {
	space := wrap.ContextSpaceValue(ctx)
	var resource store.Resource
	db := s.db.Where(&store.Resource{
		SpaceID: space,
		Name:    name,
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
	space := wrap.ContextSpaceValue(ctx)
	if _, err := s.FindByName(ctx, resource.Name); err == nil {
		return ErrExist
	}
	resource.ResourceID = util.UUID()
	resource.SpaceID = space
	return s.db.Create(resource).Error
}

func (s *resourceService) Update(ctx context.Context, resource *store.Resource) error {
	current, err := s.Find(ctx)
	if err != nil {
		return err
	}
	resource.ResourceID = current.ResourceID
	resource.SpaceID = current.SpaceID
	resource.CreatedAt = current.CreatedAt
	return s.db.Save(resource).Error
}

func (s *resourceService) Delete(ctx context.Context) error {
	space := wrap.ContextSpaceValue(ctx)
	target := wrap.ContextResourceValue(ctx)
	return s.db.Where(&store.Resource{
		ResourceID: target,
		SpaceID:    space,
	}).Delete(&store.Resource{}).Error
}

func (s *resourceService) Raw(ctx context.Context, username, space, resource string) ([]byte, error) {
	us := New().User()
	user, err := us.Find(ctx, username)
	if err != nil {
		return nil, err
	}
	ctx = wrap.ContextWithUser(ctx, &wrap.JwtUserInfo{
		UserID: user.UserID,
		Pwd:    user.Pwd,
	})
	ss := New().Space()
	spaceData, err := ss.FindByName(ctx, space)
	if err != nil {
		return nil, err
	}
	ctx = wrap.ContextWithSpace(ctx, spaceData.SpaceID)
	resourceData, err := s.FindByName(ctx, resource)
	if err != nil {
		return nil, err
	}
	return []byte(resourceData.Content), nil
}

func (s *resourceService) VersionList(ctx context.Context) ([]store.Version, error) {
	resource := wrap.ContextResourceValue(ctx)
	var versions []store.Version
	db := s.db.Where(&store.Version{
		ResourceID: resource,
	}).First(&versions)
	if db.Error == nil || database.RecordNotFound(db.Error) {
		return versions, nil
	}
	return nil, db.Error
}

func (s *resourceService) VersionFind(ctx context.Context, id string) (*store.Version, error) {
	resource := wrap.ContextResourceValue(ctx)
	var version store.Version
	db := s.db.Where(&store.Version{
		ResourceID: resource,
		VersionID:  id,
	}).First(&version)
	if db.Error == nil {
		return &version, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *resourceService) VersionFindByName(ctx context.Context, name string) (*store.Version, error) {
	resource := wrap.ContextResourceValue(ctx)
	var version store.Version
	db := s.db.Where(&store.Version{
		ResourceID: resource,
		Version:    name,
	}).First(&version)
	if db.Error == nil {
		return &version, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *resourceService) VersionCreate(ctx context.Context, version *store.Version) error {
	if _, err := s.FindByName(ctx, version.Version); err == nil {
		return ErrExist
	}
	resource, err := s.Find(ctx)
	if err != nil {
		return err
	}
	version.VersionID = util.UUID()
	version.ResourceID = resource.ResourceID
	version.Format = resource.Format
	version.Content = resource.Content
	return s.db.Create(version).Error
}

func (s *resourceService) VersionDelete(ctx context.Context, id string) error {
	resource := wrap.ContextResourceValue(ctx)
	return s.db.Where(&store.Version{
		ResourceID: resource,
		VersionID:  id,
	}).Delete(&store.Version{}).Error
}

func (s *resourceService) VersionRaw(ctx context.Context, username, space, resource, version string) ([]byte, error) {
	us := New().User()
	user, err := us.Find(ctx, username)
	if err != nil {
		return nil, err
	}
	ctx = wrap.ContextWithUser(ctx, &wrap.JwtUserInfo{
		UserID: user.UserID,
		Pwd:    user.Pwd,
	})
	ss := New().Space()
	spaceData, err := ss.FindByName(ctx, space)
	if err != nil {
		return nil, err
	}
	ctx = wrap.ContextWithSpace(ctx, spaceData.SpaceID)
	resourceData, err := s.FindByName(ctx, resource)
	if err != nil {
		return nil, err
	}
	ctx = wrap.ContextWithResource(ctx, resourceData.ResourceID)
	versionData, err := s.VersionFindByName(ctx, version)
	if err != nil {
		return nil, err
	}
	return []byte(versionData.Content), nil
}
