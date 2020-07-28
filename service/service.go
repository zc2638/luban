/**
 * Created by zc on 2020/6/10.
 */
package service

import (
	"context"
	"luban/pkg/api/store"
	"luban/pkg/errs"
)

type Interface interface {
	// User returns the UserService interface definition
	User() UserService

	// Space returns the SpaceService interface definition
	Space() SpaceService

	// Config returns the ConfigService interface definition
	Config() ConfigService
}

type (
	UserService interface {
		// All returns the all users list
		All(ctx context.Context) ([]store.User, error)

		// FindByNameAndPwd returns the current user by username and password
		FindByNameAndPwd(ctx context.Context, username, password string) (*store.User, error)

		// Create creates a user
		Create(ctx context.Context, user *store.User) error
	}
	SpaceService interface {
		// List returns the space list
		List(ctx context.Context) ([]store.Space, error)

		// Create creates a space
		Create(ctx context.Context, name string) error

		// Update updates the space info
		Update(ctx context.Context, name string) error

		// Delete deletes a space
		Delete(ctx context.Context) error
	}
	ConfigService interface {
		// List returns the config list
		List(ctx context.Context) ([]store.Config, error)

		// Find returns the current config
		Find(ctx context.Context) (*store.Config, error)

		// Create creates a config in space
		Create(ctx context.Context, config *store.Config) error

		// Update updates the config info
		Update(ctx context.Context, config *store.Config) error

		// Delete deletes a config
		Delete(ctx context.Context) error

		// Raw returns the current config content
		Raw(ctx context.Context, username, space, config string) ([]byte, error)

		// VersionList returns the version config list
		VersionList(ctx context.Context) ([]store.ConfigVersion, error)

		// VersionFind returns the current version config
		VersionFind(ctx context.Context, version string) ([]byte, error)

		// VersionCreate creates a version config
		VersionCreate(ctx context.Context, version, desc string) error

		// VersionDelete deletes a version config
		VersionDelete(ctx context.Context, version string) error

		// VersionRaw returns the current version config content
		VersionRaw(ctx context.Context, username, space, config, version string) ([]byte, error)

		// VersionDefaultSetting updates a config with a default version
		VersionDefaultSetting(ctx context.Context, version string) error
	}
)

type Service struct{}

func New() Interface {
	return &Service{}
}

func (s *Service) User() UserService {
	return &userService{}
}

func (s *Service) Space() SpaceService {
	return &spaceService{}
}

func (s *Service) Config() ConfigService {
	return &configService{}
}

const (
	ErrNotExist = errs.Error("not exist")
	ErrExist    = errs.Error("already exist")
)
