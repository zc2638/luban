/**
 * Created by zc on 2020/7/18.
 */
package uuid

import (
	"github.com/google/uuid"
	"strings"
)

func New() string {
	id := uuid.New().String()
	return strings.ReplaceAll(id, "-", "")
}