/**
 * Created by zc on 2020/8/10.
**/
package service

import (
	"context"
	"luban/global/database"
	"luban/pkg/store"
	"luban/pkg/util"
	"luban/pkg/wrap"
	"time"
)

type taskService struct{ service }

func (s *taskService) List(ctx context.Context) ([]store.Task, error) {
	pipeline := wrap.ContextPipelineValue(ctx)
	var tasks []store.Task
	db := s.db.Where(&store.Task{
		PipelineID: pipeline,
	}).Find(&tasks)
	if db.Error == nil || database.RecordNotFound(db.Error) {
		return tasks, nil
	}
	return nil, db.Error
}

func (s *taskService) ListUnComplete(ctx context.Context) ([]store.Task, error) {
	var tasks []store.Task
	db := s.db.Where(&store.Task{
		Status: store.TaskStatusPending,
	}).Or(&store.Task{
		Status: store.TaskStatusRunning,
	}).Find(&tasks)
	if db.Error == nil || database.RecordNotFound(db.Error) {
		return tasks, nil
	}
	return nil, db.Error
}

func (s *taskService) Find(ctx context.Context) (*store.Task, error) {
	pipeline := wrap.ContextPipelineValue(ctx)
	target := wrap.ContextTaskValue(ctx)
	var task store.Task
	db := s.db.Where(&store.Task{
		PipelineID: pipeline,
		TaskID:     target,
	}).First(&task)
	if db.Error == nil {
		return &task, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *taskService) Create(ctx context.Context, task *store.Task, steps []store.TaskStep) error {
	s.db = s.db.Begin()
	pipeline, err := New().Pipeline().Find(ctx)
	if err != nil {
		return err
	}
	task.PipelineID = pipeline.PipelineID
	task.Spec = pipeline.Spec
	task.Status = store.TaskStatusPending
	task.TaskID = util.UUID()
	task.StartAt = time.Now()
	if err := s.db.Create(task).Error; err != nil {
		s.db.Rollback()
		return err
	}
	ctx = wrap.ContextWithTask(ctx, task.TaskID)
	for _, step := range steps {
		if err := s.StepCreate(ctx, &step); err != nil {
			s.db.Rollback()
			return err
		}
	}
	s.db.Commit()
	return nil
}

func (s *taskService) Update(ctx context.Context, task *store.Task) error {
	current, err := s.Find(ctx)
	if err != nil {
		return err
	}
	task.TaskID = current.TaskID
	task.PipelineID = current.PipelineID
	task.Spec = current.Spec
	task.Data = current.Data
	task.CreatedAt = current.CreatedAt
	return s.db.Save(task).Error
}

func (s *taskService) StepList(ctx context.Context) ([]store.TaskStep, error) {
	task := wrap.ContextTaskValue(ctx)
	var steps []store.TaskStep
	db := s.db.Where(&store.TaskStep{
		TaskID: task,
	}).Find(&steps)
	if db.Error == nil {
		return steps, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *taskService) StepFind(ctx context.Context, id string) (*store.TaskStep, error) {
	task := wrap.ContextTaskValue(ctx)
	var step store.TaskStep
	db := s.db.Where(&store.TaskStep{
		TaskID: task,
		StepID: id,
	}).First(&step)
	if db.Error == nil {
		return &step, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *taskService) StepFindByName(ctx context.Context, name string) (*store.TaskStep, error) {
	task := wrap.ContextTaskValue(ctx)
	var step store.TaskStep
	db := s.db.Where(&store.TaskStep{
		TaskID: task,
		Name:   name,
	}).First(&step)
	if db.Error == nil {
		return &step, nil
	}
	if database.RecordNotFound(db.Error) {
		return nil, ErrNotExist
	}
	return nil, db.Error
}

func (s *taskService) StepCreate(ctx context.Context, step *store.TaskStep) error {
	if _, err := s.StepFindByName(ctx, step.Name); err == nil {
		return err
	}
	task := wrap.ContextTaskValue(ctx)
	step.TaskID = task
	step.Status = store.TaskStatusPending
	step.StepID = util.UUID()
	step.StartAt = time.Now()
	return s.db.Create(step).Error
}

func (s *taskService) StepUpdate(ctx context.Context, id string, step *store.TaskStep) error {
	current, err := s.StepFind(ctx, id)
	if err != nil {
		return err
	}
	step.TaskID = current.TaskID
	step.StepID = current.StepID
	step.Name = current.Name
	step.CreatedAt = current.CreatedAt
	if err := s.db.Save(step).Error; err != nil {
		return err
	}
	// 更新task状态
	if step.Status == store.TaskStatusFail {
		task := &store.Task{
			Status: store.TaskStatusFail,
			EndAt:  step.EndAt,
		}
		return s.Update(ctx, task)
	}
	task := wrap.ContextTaskValue(ctx)
	var not store.TaskStep
	db := s.db.Where(&store.TaskStep{
		TaskID: task,
	}).Not(&store.TaskStep{
		Status: store.TaskStatusSuccess,
	}).First(&not)
	if database.RecordNotFound(db.Error) {
		task := &store.Task{
			Status: store.TaskStatusSuccess,
			EndAt:  step.EndAt,
		}
		return s.Update(ctx, task)
	}
	return db.Error
}
