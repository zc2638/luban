/**
 * Created by zc on 2020/7/18.
 */
package fs

import "os"

type Interface interface {
	Dir() DirInterface
	File() FileInterface
}

type DirInterface interface {
	IsExist(path string) bool
	List(path string) ([]string, error)
	ListFile(path string) ([]os.FileInfo, error)
	Create(path string) error
	Update(target, path string) error
	Delete(path string) error
}

type FileInterface interface {
	Find(path string) ([]byte, error)
	Create(dir, name string) (*os.File, error)
	Update(path string, data []byte) error
	Delete(path string) error
}

type FS struct{}

func New() *FS {
	return &FS{}
}

func (*FS) Dir() DirInterface {
	return NewDir()
}

func (*FS) File() FileInterface {
	return NewFile()
}
