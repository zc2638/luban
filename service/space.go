/**
 * Created by zc on 2020/6/11.
 */
package service

import (
	"context"
	"luban/pkg/ctr"
	"luban/pkg/database"
	"luban/pkg/database/data"
	"luban/pkg/uuid"
)

type spaceService struct{ service }

func (s *spaceService) List(ctx context.Context) ([]data.Space, error) {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	var spaces []data.Space
	db := s.db.Where(&data.Space{
		UserID: user.UserID,
	}).Find(&spaces)
	if db.Error == nil || database.RecordNotFound(db.Error) {
		return spaces, nil
	}
	return nil, db.Error
}

func (s *spaceService) Find(ctx context.Context) (*data.Space, error) {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	target := ctr.ContextSpaceValue(ctx)
	var space data.Space
	db := s.db.Where(&data.Space{
		UserID:  user.UserID,
		SpaceID: target,
	}).First(&space)
	if db.Error == nil {
		return &space, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *spaceService) FindByName(ctx context.Context, name string) (*data.Space, error) {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	var space data.Space
	db := s.db.Where(&data.Space{
		UserID: user.UserID,
		Name:   name,
	}).First(&space)
	if db.Error == nil {
		return &space, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *spaceService) Create(ctx context.Context, space *data.Space) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	if _, err = s.FindByName(ctx, space.Name); err == nil {
		return ErrExist
	}
	space.SpaceID = uuid.New()
	space.UserID = user.UserID
	return s.db.Create(space).Error
}

func (s *spaceService) Update(ctx context.Context, space *data.Space) error {
	current, err := s.Find(ctx)
	if err != nil {
		return err
	}
	space.SpaceID = current.SpaceID
	space.UserID = current.UserID
	space.CreatedAt = current.CreatedAt
	return s.db.Save(space).Error
}

func (s *spaceService) Delete(ctx context.Context) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	target := ctr.ContextSpaceValue(ctx)
	return s.db.Where(&data.Space{
		UserID:  user.UserID,
		SpaceID: target,
	}).Delete(&data.Space{}).Error
}
