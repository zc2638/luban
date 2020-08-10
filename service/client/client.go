/**
 * Created by zc on 2020/8/2.
 */
package client

import (
	"context"
	"encoding/json"
	"github.com/zc2638/gotool/curlx"
	"luban/pkg/api"
)

type client struct {
	host string
}

func (c *client) Request(ctx context.Context) (*api.Task, error) {
	h := curlx.Request{}
	h.Url = c.host + "/v1/api/async/request"
	body, err := h.Get()
	if err != nil {
		return nil, err
	}
	var result api.Task
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (c *client) Watch(ctx context.Context, taskId string) (*api.Task, error) {
	h := curlx.Request{}
	h.Url = c.host + "/v1/api/async/watch/" + taskId
	body, err := h.Get()
	if err != nil {
		return nil, err
	}
	var result api.Task
	err = json.Unmarshal(body, &result)
	return &result, err
}

func (c *client) FlowConfig() {

}

func (c *client) CustomConfig() {

}