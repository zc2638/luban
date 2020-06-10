/**
 * Created by zc on 2020/6/9.
 */
package ctr

import (
	"encoding/json"
	"io"
	"stone/pkg/errs"
)

func JSONParseReader(rc io.ReadCloser, v interface{}) error {
	defer rc.Close()
	if err := json.NewDecoder(rc).Decode(v); err != nil {
		return errs.ErrBodyParse.With(err)
	}
	return nil
}
