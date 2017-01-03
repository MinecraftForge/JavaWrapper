package asset

import (
	"golang.org/x/tools/godoc/vfs"
	"fmt"
	"os"
)

type failFS struct {
	err error
}

func failfs(err error) vfs.FileSystem {
	return &failFS{fmt.Errorf("asset: %s", err.Error())}
}

func (fs *failFS) Open(name string) (vfs.ReadSeekCloser, error) {
	return nil, fs.err
}
func (fs *failFS) Lstat(path string) (os.FileInfo, error) {
	return nil, fs.err
}
func (fs *failFS) Stat(path string) (os.FileInfo, error) {
	return nil, fs.err
}
func (fs *failFS) ReadDir(path string) ([]os.FileInfo, error) {
	return nil, fs.err
}
func (fs *failFS) String() string {
	return fmt.Sprintf("failfs(%q)", fs.err)
}
