/**
 * Created by zc on 2020/8/10.
**/
package service

import (
	"context"
	"luban/pkg/ctr"
	"luban/pkg/database/data"
	"luban/pkg/uuid"
	"time"
)

type taskService struct{ service }

func (s *taskService) List(ctx context.Context) ([]data.Task, error) {
	pipeline := ctr.ContextPipelineValue(ctx)
	var tasks []data.Task
	db := s.db.Where(&data.Task{
		PipelineID: pipeline,
	}).Find(&tasks)
	if db.Error == nil || db.RecordNotFound() {
		return tasks, nil
	}
	return nil, db.Error
}

func (s *taskService) Find(ctx context.Context, id string) (*data.Task, error) {
	pipeline := ctr.ContextPipelineValue(ctx)
	var task data.Task
	db := s.db.Where(&data.Task{
		PipelineID: pipeline,
		TaskID:     id,
	}).First(&task)
	if db.Error == nil {
		return &task, nil
	}
	if db.RecordNotFound() {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *taskService) Create(ctx context.Context, task *data.Task, steps []data.TaskStep) error {
	s.db = s.db.Begin()
	pipeline := ctr.ContextPipelineValue(ctx)
	task.PipelineID = pipeline
	task.TaskID = uuid.New()
	task.StartAt = time.Now()
	if err := s.db.Create(task).Error; err != nil {
		s.db.Rollback()
		return err
	}
	ctx = ctr.ContextWithTask(ctx, task.TaskID)
	for _, step := range steps {
		if err := s.StepCreate(ctx, &step); err != nil {
			s.db.Rollback()
			return err
		}
	}
	s.db.Commit()
	return nil
}

func (s *taskService) Update(ctx context.Context, id string, task *data.Task) error {
	current, err := s.Find(ctx, id)
	if err != nil {
		return err
	}
	task.TaskID = current.TaskID
	task.PipelineID = current.PipelineID
	task.CreatedAt = current.CreatedAt
	return s.db.Save(task).Error
}

func (s *taskService) StepList(ctx context.Context) ([]data.TaskStep, error) {
	task := ctr.ContextTaskValue(ctx)
	var steps []data.TaskStep
	db := s.db.Where(&data.TaskStep{
		TaskID: task,
	}).Find(&steps)
	if db.Error == nil {
		return steps, nil
	}
	if db.RecordNotFound() {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *taskService) StepFind(ctx context.Context, id string) (*data.TaskStep, error) {
	task := ctr.ContextTaskValue(ctx)
	var step data.TaskStep
	db := s.db.Where(&data.TaskStep{
		TaskID: task,
		StepID: id,
	}).First(&step)
	if db.Error == nil {
		return &step, nil
	}
	if db.RecordNotFound() {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *taskService) StepFindByName(ctx context.Context, name string) (*data.TaskStep, error) {
	task := ctr.ContextTaskValue(ctx)
	var step data.TaskStep
	db := s.db.Where(&data.TaskStep{
		TaskID: task,
		Name:   name,
	}).First(&step)
	if db.Error == nil {
		return &step, nil
	}
	if db.RecordNotFound() {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *taskService) StepCreate(ctx context.Context, step *data.TaskStep) error {
	if _, err := s.StepFindByName(ctx, step.Name); err == nil {
		return err
	}
	task := ctr.ContextTaskValue(ctx)
	step.TaskID = task
	step.StepID = uuid.New()
	step.StartAt = time.Now()
	return s.db.Create(step).Error
}

func (s *taskService) StepUpdate(ctx context.Context, id string, step *data.TaskStep) error {
	current, err := s.StepFind(ctx, id)
	if err != nil {
		return err
	}
	step.TaskID = current.TaskID
	step.StepID = current.StepID
	step.CreatedAt = current.CreatedAt
	return s.db.Save(step).Error
}
