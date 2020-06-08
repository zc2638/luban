package hook

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

/**
 * Created by zc on 2020-04-16.
 */
const (
	// DefaultPath 默认地址
	DefaultPath = "_log"
)

type FileHook struct {
	w        io.ReadWriteCloser
	dir      string
	datetime string
	name     string
}

func NewHook(name string) *FileHook {
	return &FileHook{name: name, dir: DefaultPath}
}

func (h *FileHook) SetDir(dir string) *FileHook {
	filterDir := strings.TrimRight(dir, "/")
	if filterDir != "" {
		h.dir = filterDir
	}
	return h
}

func (h *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// 自定义钩子执行（默认协程安全）
func (h *FileHook) Fire(e *logrus.Entry) error {
	// 判断log文件是否变更
	now := time.Now().Format("20060102")
	path := h.dir + "/" + h.name + now + ".log"
	if h.datetime != now {
		h.datetime = now
		if h.w != nil {
			if err := h.w.Close(); err != nil {
				return err
			}
			h.w = nil
		}
	}
	if _, err := os.Stat(path); err == nil && h.w != nil {
		return nil
	}

	// 自动创建文件
	var pathArr = strings.Split(path, "/")
	var pathLen = len(pathArr)
	if pathLen > 1 {
		dir := strings.Join(pathArr[:pathLen-1], "/")
		// 自动创建日志文件夹
		_, err := os.Stat(dir)
		if err != nil {
			mkErr := os.MkdirAll(dir, os.ModePerm)
			if mkErr != nil {
				return mkErr
			}
		}
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return err
	}

	h.w = f
	e.Logger.Out = h.w
	return nil
}
