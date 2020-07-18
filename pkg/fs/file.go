/**
 * Created by zc on 2020/7/18.
 */
package fs

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

type File struct{}

func NewFile() *File {
	return &File{}
}

func (f *File) Find(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func (f *File) Create(path, name string) (*os.File, error) {
	dir := NewDir()
	if err := dir.Create(path); err != nil {
		return nil, err
	}
	return os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, os.ModePerm)
}

func (f *File) Update(path string, data []byte) error {
	dirPath := filepath.Dir(path)
	dir := NewDir()
	if err := dir.Create(dirPath); err != nil {
		return err
	}
	return ioutil.WriteFile(path, data, os.ModePerm)
}

func (f *File) Delete(path string) error {
	return os.Remove(path)
}
