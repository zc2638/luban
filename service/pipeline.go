/**
 * Created by zc on 2020/8/10.
**/
package service

import (
	"context"
	"luban/global/database"
	"luban/pkg/store"
	"luban/pkg/util"
	"luban/pkg/wrapper"
)

type pipelineService struct{ service }

func (s *pipelineService) List(ctx context.Context) ([]store.Pipeline, error) {
	space := wrapper.ContextSpaceValue(ctx)
	var pipelines []store.Pipeline
	db := s.db.Where(&store.Pipeline{
		SpaceID: space,
	}).Find(&pipelines)
	if db.Error == nil || database.RecordNotFound(db.Error) {
		return pipelines, nil
	}
	return nil, db.Error
}

func (s *pipelineService) Find(ctx context.Context) (*store.Pipeline, error) {
	space := wrapper.ContextSpaceValue(ctx)
	target := wrapper.ContextPipelineValue(ctx)
	var pipeline store.Pipeline
	db := s.db.Where(&store.Pipeline{
		SpaceID:    space,
		PipelineID: target,
	}).First(&pipeline)
	if db.Error == nil {
		return &pipeline, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *pipelineService) FindByName(ctx context.Context, name string) (*store.Pipeline, error) {
	space := wrapper.ContextSpaceValue(ctx)
	var pipeline store.Pipeline
	db := s.db.Where(&store.Pipeline{
		SpaceID: space,
		Name:    name,
	}).First(&pipeline)
	if db.Error == nil {
		return &pipeline, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *pipelineService) Create(ctx context.Context, pipeline *store.Pipeline) error {
	if _, err := s.FindByName(ctx, pipeline.Name); err == nil {
		return ErrExist
	}
	space := wrapper.ContextSpaceValue(ctx)
	pipeline.SpaceID = space
	pipeline.PipelineID = util.UUID()
	return s.db.Create(pipeline).Error
}

func (s *pipelineService) Update(ctx context.Context, pipeline *store.Pipeline) error {
	current, err := s.Find(ctx)
	if err != nil {
		return err
	}
	pipeline.PipelineID = current.PipelineID
	pipeline.CreatedAt = current.CreatedAt
	return s.db.Save(pipeline).Error
}

func (s *pipelineService) Delete(ctx context.Context) error {
	space := wrapper.ContextSpaceValue(ctx)
	target := wrapper.ContextPipelineValue(ctx)
	return s.db.Where(&store.Pipeline{
		SpaceID:    space,
		PipelineID: target,
	}).Delete(&store.Pipeline{}).Error
}
