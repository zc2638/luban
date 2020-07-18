/**
 * Created by zc on 2020/6/9.
 */
package ctr

import (
	"encoding/json"
	"io"
	"luban/pkg/errs"
)

// JSONParseReader decode Reader to any struct
func JSONParseReader(rc io.ReadCloser, v interface{}) error {
	defer rc.Close()
	if err := json.NewDecoder(rc).Decode(v); err != nil {
		return errs.ErrBodyParse.With(err)
	}
	return nil
}
