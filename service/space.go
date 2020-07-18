/**
 * Created by zc on 2020/6/11.
 */
package service

import (
	"context"
	"github.com/google/uuid"
	"luban/global"
	"luban/pkg/api/store"
	"luban/pkg/ctr"
)

type spaceService struct{}

func (s *spaceService) List(ctx context.Context) ([]*store.Space, error) {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	var list []*store.Space
	db := global.DB().Where(&store.Space{Owner: user.UID}).Find(&list)
	return list, db.Error
}

func (s *spaceService) Find(ctx context.Context, id string) (*store.Space, error) {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return nil, err
	}
	var space store.Space
	db := global.DB().Where(&store.Space{Owner: user.UID, SID: id}).First(&space)
	if db.Error != nil {
		return nil, db.Error
	}
	return &space, nil
}

func (s *spaceService) Create(ctx context.Context, space *store.Space) error {
	user, err := ctr.ContextUserFrom(ctx)
	if err != nil {
		return err
	}
	space.SID = uuid.New().String()
	space.Owner = user.UID
	return global.DB().Create(space).Error
}

func (s *spaceService) Update(ctx context.Context, space *store.Space) error {
	info, err := s.Find(ctx, space.SID)
	if err != nil {
		return err
	}
	return global.DB().Model(info).Updates(space).Error
}

func (s *spaceService) Delete(ctx context.Context, id string) error {
	info, err := s.Find(ctx, id)
	if err != nil {
		return err
	}
	return global.DB().Delete(info).Error
}
