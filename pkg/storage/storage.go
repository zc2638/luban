/**
 * Created by zc on 2020/7/25.
 */
package storage

import (
	"github.com/spf13/afero"
	"os"
	"path/filepath"
)

type Interface interface {
	// PathKeys returns the path name collection under path
	PathKeys(path string) ([]string, error)

	// PathUpdate updates the path name
	PathUpdate(target, path string) error

	// Keys returns the value key name collection under path
	Keys(path string) ([]string, error)

	// Find returns the value by key under path
	Find(path, key string) ([]byte, error)

	// Create creates a value by key under path
	Create(path, key string, data []byte) error

	// Update updates a value by key under path
	Update(path, key string, data []byte) error

	// Delete deletes a value by key under path
	Delete(path, key string) error
}

func New() Interface {
	return DefaultFStorage
}

var DefaultFStorage = NewFile(afero.NewOsFs())

type FileSystem struct {
	Fs afero.Fs
}

func NewFile(fs afero.Fs) *FileSystem {
	return &FileSystem{Fs: fs}
}

func (fs *FileSystem) PathKeys(path string) ([]string, error) {
	file, err := fs.Fs.Open(path)
	if err != nil {
		return nil, err
	}
	infos, err := file.Readdir(-1)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(infos))
	for _, info := range infos {
		if !info.IsDir() {
			continue
		}
		keys = append(keys, info.Name())
	}
	return keys, nil
}

func (fs *FileSystem) PathUpdate(target, path string) error {
	exists, err := afero.Exists(fs.Fs, target)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	return fs.Fs.Rename(target, path)
}

func (fs *FileSystem) Keys(path string) ([]string, error) {
	file, err := fs.Fs.Open(path)
	if err != nil {
		return nil, err
	}
	infos, err := file.Readdir(-1)
	if err != nil {
		return nil, err
	}
	keys := make([]string, 0, len(infos))
	for _, info := range infos {
		if info.IsDir() {
			continue
		}
		keys = append(keys, info.Name())
	}
	return keys, nil
}

func (fs *FileSystem) Find(path, key string) ([]byte, error) {
	fp := filepath.Join(path, key)
	exists, err := afero.Exists(fs.Fs, fp)
	if err != nil {
		return nil, err
	}
	if !exists {
		return []byte{}, nil
	}
	return afero.ReadFile(fs.Fs, fp)
}

func (fs *FileSystem) Create(path, key string, data []byte) error {
	fp := filepath.Join(path, key)
	exists, err := afero.Exists(fs.Fs, fp)
	if err != nil {
		return err
	}
	if exists {
		return os.ErrExist
	}
	if err := fs.Fs.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	file, err := fs.Fs.Create(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

func (fs *FileSystem) Update(path, key string, data []byte) error {
	fp := filepath.Join(path, key)
	exists, err := afero.Exists(fs.Fs, fp)
	if err != nil {
		return err
	}
	if !exists {
		if err := fs.Fs.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	file, err := fs.Fs.Create(fp)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

func (fs *FileSystem) Delete(path, key string) error {
	fp := filepath.Join(path, key)
	return fs.Fs.RemoveAll(fp)
}
