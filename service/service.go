/**
 * Created by zc on 2020/6/10.
 */
package service

import (
	"context"
	"gorm.io/gorm"
	"luban/global"
	"luban/pkg/errs"
	"luban/pkg/store"
)

type Service interface {
	Init()
}

type service struct {
	db *gorm.DB
}

func (s *service) Init() {
	s.db = global.DB()
}

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
		// Find returns the current user by username
		Find(ctx context.Context, name string) (*store.User, error)

		// FindByNameAndPwd returns the current user by username and password
		FindByNameAndPwd(ctx context.Context, username, password string) (*store.User, error)

		// FindByUserID returns the current user by user id
		FindByUserID(ctx context.Context, userID string) (*store.User, error)

		// Create creates a user
		Create(ctx context.Context, user *store.User) error

		// PwdReset resets the password to user
		PwdReset(ctx context.Context, username, password string) error
	}

	// SpaceService defines the space related operations
	SpaceService interface {
		// List returns the space list
		List(ctx context.Context) ([]store.Space, error)

		// Find returns the current space
		Find(ctx context.Context) (*store.Space, error)

		// FindByName returns the current space by name
		FindByName(ctx context.Context, name string) (*store.Space, error)

		// Create creates a space
		Create(ctx context.Context, space *store.Space) error

		// Update updates the space info
		Update(ctx context.Context, space *store.Space) error

		// Delete deletes a space
		Delete(ctx context.Context) error
	}

	// ResourceService defines the resource related operations
	ResourceService interface {
		// List returns the resource list
		List(ctx context.Context) ([]store.Resource, error)

		// Find returns the current resource
		Find(ctx context.Context) (*store.Resource, error)

		// Create creates a resource in space
		Create(ctx context.Context, resource *store.Resource) error

		// Update updates the resource info
		Update(ctx context.Context, resource *store.Resource) error

		// Delete deletes a resource
		Delete(ctx context.Context) error

		// Raw returns the current resource content
		Raw(ctx context.Context, username, space, resource string) ([]byte, error)

		// VersionList returns the version resource list
		VersionList(ctx context.Context) ([]store.Version, error)

		// VersionFind returns the current version resource
		VersionFind(ctx context.Context, id string) (*store.Version, error)

		// VersionCreate creates a version resource
		VersionCreate(ctx context.Context, version *store.Version) error

		// VersionDelete deletes a version resource
		VersionDelete(ctx context.Context, id string) error

		// VersionRaw returns the current version resource content
		VersionRaw(ctx context.Context, username, space, resource, version string) ([]byte, error)
	}

	// PipelineService defines the pipeline related operations
	PipelineService interface {
		// List returns the pipeline list
		List(ctx context.Context) ([]store.Pipeline, error)

		// Find returns the current pipeline
		Find(ctx context.Context) (*store.Pipeline, error)

		// Create creates a pipeline
		Create(ctx context.Context, pipeline *store.Pipeline) error

		// Update updates the pipeline info
		Update(ctx context.Context, pipeline *store.Pipeline) error

		// Delete deletes a pipeline
		Delete(ctx context.Context) error
	}

	// TaskService defines the task related operations
	TaskService interface {
		// List returns the task list
		List(ctx context.Context) ([]store.Task, error)

		// ListUnComplete returns the not complete task list
		ListUnComplete(ctx context.Context) ([]store.Task, error)

		// Find returns the current task
		Find(ctx context.Context) (*store.Task, error)

		// Create creates a task
		Create(ctx context.Context, task *store.Task, steps []store.TaskStep) error

		// Update updates the task info
		Update(ctx context.Context, task *store.Task) error

		// StepList returns the task step list
		StepList(ctx context.Context) ([]store.TaskStep, error)

		// StepFind returns the current task step
		StepFind(ctx context.Context, id string) (*store.TaskStep, error)

		// StepCreate creates a task step
		StepCreate(ctx context.Context, step *store.TaskStep) error

		// StepUpdate updates the task step info
		StepUpdate(ctx context.Context, id string, step *store.TaskStep) error
	}
)

// Default returns the default service, change it if need.
var Default = &DefaultService{}

type DefaultService struct{}

func New() Interface {
	return Default
}

func (s *DefaultService) Init(svc Service) { svc.Init() }

func (s *DefaultService) User() UserService {
	svc := &userService{}
	s.Init(svc)
	return svc
}

func (s *DefaultService) Space() SpaceService {
	svc := &spaceService{}
	s.Init(svc)
	return svc
}

func (s *DefaultService) Resource() ResourceService {
	svc := &resourceService{}
	s.Init(svc)
	return svc
}

func (s *DefaultService) Pipeline() PipelineService {
	svc := &pipelineService{}
	s.Init(svc)
	return svc
}

func (s *DefaultService) Task() TaskService {
	svc := &taskService{}
	s.Init(svc)
	return svc
}

const (
	ErrNotExist = errs.Error("not exist")
	ErrExist    = errs.Error("already exist")
)
