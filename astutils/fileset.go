package astutils

import (
	"go/token"
)

type FileSet struct {
	fs *token.FileSet
}

func NewFileSet() *FileSet {
	return &FileSet{
		fs: token.NewFileSet(),
	}
}

func (fs *FileSet) AddFile(name, packageName string) *File {
	return NewFileWithFileSet(fs, name, packageName)
}
