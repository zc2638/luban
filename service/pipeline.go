/**
 * Created by zc on 2020/8/10.
**/
package service

import (
	"context"
	"luban/pkg/ctr"
	"luban/pkg/database/data"
	"luban/pkg/uuid"
)

type pipelineService struct{ service }

func (s *pipelineService) List(ctx context.Context) ([]data.Pipeline, error) {
	space := ctr.ContextSpaceValue(ctx)
	var pipelines []data.Pipeline
	db := s.db.Where(&data.Pipeline{
		SpaceID: space,
	}).Find(&pipelines)
	if db.Error == nil || db.RecordNotFound() {
		return pipelines, nil
	}
	return nil, db.Error
}

func (s *pipelineService) Find(ctx context.Context) (*data.Pipeline, error) {
	space := ctr.ContextSpaceValue(ctx)
	target := ctr.ContextPipelineValue(ctx)
	var pipeline data.Pipeline
	db := s.db.Where(&data.Pipeline{
		SpaceID:    space,
		PipelineID: target,
	}).First(&pipeline)
	if db.Error == nil {
		return &pipeline, nil
	}
	if db.RecordNotFound() {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *pipelineService) FindByName(ctx context.Context, name string) (*data.Pipeline, error) {
	space := ctr.ContextSpaceValue(ctx)
	var pipeline data.Pipeline
	db := s.db.Where(&data.Pipeline{
		SpaceID: space,
		Name:    name,
	}).First(&pipeline)
	if db.Error == nil {
		return &pipeline, nil
	}
	if db.RecordNotFound() {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *pipelineService) Create(ctx context.Context, pipeline *data.Pipeline) error {
	if _, err := s.FindByName(ctx, pipeline.Name); err == nil {
		return ErrExist
	}
	space := ctr.ContextSpaceValue(ctx)
	pipeline.SpaceID = space
	pipeline.PipelineID = uuid.New()
	return s.db.Create(pipeline).Error
}

func (s *pipelineService) Update(ctx context.Context, pipeline *data.Pipeline) error {
	current, err := s.Find(ctx)
	if err != nil {
		return err
	}
	pipeline.PipelineID = current.PipelineID
	pipeline.CreatedAt = current.CreatedAt
	return s.db.Save(pipeline).Error
}

func (s *pipelineService) Delete(ctx context.Context) error {
	space := ctr.ContextSpaceValue(ctx)
	target := ctr.ContextPipelineValue(ctx)
	return s.db.Where(&data.Pipeline{
		SpaceID:    space,
		PipelineID: target,
	}).Delete(&data.Pipeline{}).Error
}
