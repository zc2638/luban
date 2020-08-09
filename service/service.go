/**
 * Created by zc on 2020/6/10.
 */
package service

import (
	"context"
	"luban/pkg/database/data"
	"luban/pkg/errs"
)

type Interface interface {
	// User returns the UserService interface definition
	User() UserService

	// Space returns the SpaceService interface definition
	Space() SpaceService

	// Resource returns the ResourceService interface definition
	Resource() ResourceService

	// Pipeline returns the PipelineService interface definition
	Pipeline() PipelineService

	// Task returns the TaskService interface definition
	Task() TaskService
}

type (
	// UserService defines the user related operations
	UserService interface {
		// FindByNameAndPwd returns the current user by username and password
		FindByNameAndPwd(ctx context.Context, username, password string) (*data.User, error)

		// FindByUserID returns the current user by user id
		FindByUserID(ctx context.Context, userID string) (*data.User, error)

		// Create creates a user
		Create(ctx context.Context, user *data.User) error

		// PwdReset resets the password to user
		PwdReset(ctx context.Context, username, password string) error
	}

	// SpaceService defines the space related operations
	SpaceService interface {
		// List returns the space list
		List(ctx context.Context) ([]data.Space, error)

		// Find returns the current space
		Find(ctx context.Context) (*data.Space, error)

		// Create creates a space
		Create(ctx context.Context, space *data.Space) error

		// Update updates the space info
		Update(ctx context.Context, space *data.Space) error

		// Delete deletes a space
		Delete(ctx context.Context) error
	}

	// ResourceService defines the resource related operations
	ResourceService interface {
		// List returns the config list
		List(ctx context.Context) ([]data.Resource, error)

		// Find returns the current config
		Find(ctx context.Context) (*data.Resource, error)

		// Create creates a config in space
		Create(ctx context.Context, resource *data.Resource) error

		// Update updates the config info
		Update(ctx context.Context, resource *data.Resource) error

		// Delete deletes a config
		Delete(ctx context.Context) error

		// Raw returns the current config content
		Raw(ctx context.Context, username, space, resource string) ([]byte, error)

		// VersionList returns the version config list
		VersionList(ctx context.Context) ([]data.Version, error)

		// VersionFind returns the current version config
		VersionFind(ctx context.Context, id string) (*data.Version, error)

		// VersionCreate creates a version config
		VersionCreate(ctx context.Context, version *data.Version) error

		// VersionDelete deletes a version config
		VersionDelete(ctx context.Context, id string) error

		// VersionRaw returns the current version config content
		VersionRaw(ctx context.Context, username, space, resource, version string) ([]byte, error)
	}

	// PipelineService defines the pipeline related operations
	PipelineService interface {
	}

	// TaskService defines the task related operations
	TaskService interface {
	}
)

// Default returns the default service, change it if need.
var Default = &Service{}

type Service struct{}

func New() Interface {
	return Default
}

func (s *Service) User() UserService {
	return &userService{}
}

func (s *Service) Space() SpaceService {
	return &spaceService{}
}

func (s *Service) Resource() ResourceService {
	return &resourceService{}
}

func (s *Service) Pipeline() PipelineService {
	panic("implement me")
}

func (s *Service) Task() TaskService {
	panic("implement me")
}

const (
	ErrNotExist = errs.Error("not exist")
	ErrExist    = errs.Error("already exist")
)
