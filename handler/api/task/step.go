/**
 * Created by zc on 2020/8/10.
**/
package task

import (
	"context"
	"github.com/go-chi/chi"
	"github.com/pkgms/go/ctr"
	"luban/pkg/api/request"
	"luban/pkg/errs"
	"luban/pkg/store"
	"luban/pkg/wrap"
	"luban/service"
	"net/http"
	"time"
)

func StepList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		list, err := service.New().Task().StepList(context.Background())
		if err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.OK(w, list)
	}
}

func StepUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params request.TaskStepParams
		if err := wrap.JSONParseReader(r.Body, &params); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		id := chi.URLParam(r, "step")
		if params.Status != store.TaskStatusSuccess && params.Status != store.TaskStatusFail {
			ctr.BadRequest(w, errs.New("the step status not support to complete"))
		}
		step := &store.TaskStep{
			Status: params.Status,
			Log:    params.Log,
			EndAt:  time.Now(),
		}
		if err := service.New().Task().StepUpdate(r.Context(), id, step); err != nil {
			ctr.BadRequest(w, err)
			return
		}
		ctr.Success(w)
	}
}
