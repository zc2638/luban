/**
 * Created by zc on 2020/7/18.
 */
package fs

import (
	"io/ioutil"
	"luban/pkg/errs"
	"os"
)

type Dir struct{}

func NewDir() *Dir {
	return &Dir{}
}

func (d *Dir) IsExist(path string) bool {
	_, err := os.Stat(path)
	return os.IsExist(err)
}

func (d *Dir) List(path string) ([]string, error) {
	stat, err := os.Stat(path)
	if os.IsNotExist(err) {
		return []string{}, nil
	}
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errs.New("not dir")
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return f.Readdirnames(-1)
}

func (d *Dir) ListFile(path string) ([]os.FileInfo, error) {
	fs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	files := make([]os.FileInfo, 0, len(fs))
	for _, f := range fs {
		if f.IsDir() {
			continue
		}
		files = append(files, f)
	}
	return files, err
}

func (d *Dir) Create(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func (d *Dir) Update(target, path string) error {
	return os.Rename(target, path)
}

func (d *Dir) Delete(path string) error {
	return os.RemoveAll(path)
}
