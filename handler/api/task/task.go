/**
 * Created by zc on 2020/8/10.
**/
package task

import (
	"encoding/json"
	"github.com/drone/drone-yaml/yaml"
	"github.com/pkgms/go/ctr"
	"luban/pkg/api/response"
	"luban/pkg/store"
	"luban/pkg/wrap"
	"luban/service"
	"net/http"
)

func List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := service.New().Task().List(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, list)
	}
}

func Info() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		info, err := service.New().Task().Find(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		list, err := service.New().Task().StepList(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		steps := make([]response.TaskStepResultItem, 0, len(list))
		for _, step := range list {
			steps = append(steps, response.TaskStepResultItem{
				StepID:    step.StepID,
				Name:      step.Name,
				Status:    step.Status,
				Log:       step.Log,
				StartTS:   step.StartAt.Unix(),
				EndTS:     step.EndAt.Unix(),
				CreatedTS: step.CreatedAt.Unix(),
				UpdatedTS: step.UpdatedAt.Unix(),
			})
		}
		ctr.OK(w, response.TaskResult{
			TaskID:     info.TaskID,
			PipelineID: info.PipelineID,
			Status:     info.Status,
			StartTS:    info.StartAt.Unix(),
			EndTS:      info.EndAt.Unix(),
			CreatedTS:  info.CreatedAt.Unix(),
			UpdatedTS:  info.UpdatedAt.Unix(),
			Steps:      steps,
		})
	}
}

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 查出 pipeline 的 spec
		pipeline, err := service.New().Pipeline().Find(r.Context())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		// 解析 spec
		var spec yaml.Pipeline
		if err := json.Unmarshal([]byte(pipeline.Spec), &spec); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		task := &store.Task{Spec: pipeline.Spec}
		if pipeline.ResourceID != "" {
			ctx := wrap.ContextWithResource(r.Context(), pipeline.ResourceID)
			resource, err := service.New().Resource().Find(ctx)
			if err != nil {
				ctr.BadRequest(w, err)
				return
			}
			task.Data = resource.Content
		}
		steps := make([]store.TaskStep, 0, len(spec.Steps))
		for _, step := range spec.Steps {
			steps = append(steps, store.TaskStep{
				Name: step.Name,
			})
		}
		if err := service.New().Task().Create(r.Context(), task, steps); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
