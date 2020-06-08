/**
 * Created by zc on 2020/5/22.
 */
package hook

import (
	"bytes"
	"github.com/sirupsen/logrus"
	"runtime"
	"strconv"
	"strings"
)

type FrameContextHook struct {
	ws  []string
	cap uint32
}

func NewFrameContextHook(c uint32, ws []string) *FrameContextHook {
	if ws == nil {
		ws = []string{"github.com/sirupsen/logrus", "github.com/gin-gonic/gin", "/pkg/logger", "net/http", "/runtime"}
	}
	if c == 0 {
		c = 30
	}
	return &FrameContextHook{ws: ws, cap: c}
}

func (f *FrameContextHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	}
}

func (f *FrameContextHook) Fire(entry *logrus.Entry) error {
	pc := make([]uintptr, f.cap)
	n := runtime.Callers(-1, pc)
	if n > 0 {
		pc = pc[:n]
		frames := runtime.CallersFrames(pc)
		var chains []string
		for {
			frame, more := frames.Next()
			skip := false
			for _, w := range f.ws {
				if strings.Contains(frame.File, w) {
					skip = true
					break
				}
			}
			if !skip {
				var buffer bytes.Buffer
				_, _ = buffer.WriteString(frame.File)
				_, _ = buffer.WriteRune(':')
				_, _ = buffer.WriteString(strconv.Itoa(frame.Line))
				_, _ = buffer.WriteRune(' ')
				_, _ = buffer.WriteString(frame.Function)
				chains = append(chains, buffer.String())
			}
			if !more {
				break
			}
		}
		entry.Data["chain"] = chains
	}
	return nil
}
