/*
Package asset lets you embed files in your go binary by simply appending a zip file to said binary.
Typically on launch of your program, you call asset.Fs() to get a
handle to the vfs.FileSystem that represents the appended zip file.
It is on this vfs that you should call Open and friends, as opposed to
the normal os.Open.

NOTE: for serving files through http, use the newer https://godoc.org/bitbucket.org/mjl/httpasset.

See https://godoc.org/golang.org/x/tools/godoc/vfs for vfs.FileSystem.

An example:

	// mybinary.go

	import (
		"golang.org/x/tools/godoc/vfs" // optional, for vfs.OS
		"bitbucket.org/mjl/asset"
	)

	func main() {
		// the error-check is optional, asset.Fs() always returns a non-nil vfs.FileSystem.
		// however, after failed initialization (eg no zip file was appended to the binary),
		// vfs operations return an error.
		fs := asset.Fs()
		if err := asset.Error(); err != nil {
			log.Fatal(err)
			// or alternatively fallback to to local file system:
			// fs = vfs.OS(".")
		}

		// note: paths must be absolute, i.e. starting with a slash.
		f, err := fs.Open("/test.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		io.Copy(os.Stdout, f)
	}

Build your program, let's say the result is "mybinary".
Now create a zip file, eg on unix:

	zip -r0 mybinary.zip test.txt

Append it to the binary:

	cat mybinary.zip >>mybinary

If you run mybinary, it will print the contents of file "test.txt"
in the zip file that you appended to the binary.

To make this work, an assumption about zip files is made:
That the central directory (with a list of files inside the zip file)
comes right before the "end of central directory" marker.  This is almost
always the case with zip files.  With this assumption, asset can locate
the start and end of the zip file that is appended to the binary, which
archive/zip needs in order to parse the zip file.

Some existing tools for reading zip files can still read the
binary-with-zipfile as a zip file.  For example 7z, and the unzip
command-line tool.  Windows XP's explorer zip opener does NOT seem to
understand it, and also Mac OS X's archive utility gets confused.

This has been tested with binaries on Linux (Ubuntu 12.04), Mac OS X
(10.9.2) and Windows 8.1.  These operating systems don't
seem to mind extra data at the end of the binary.
*/
package asset

import (
	"archive/zip"
	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

type assetFS struct {
	vfs.FileSystem
	rc *zip.ReadCloser
}

var fs vfs.FileSystem

// Fs returns the vfs.FileSystem for the assets contained in the binary.
// It always returns a non-nil FileSystem.  In case of an initialization
// error a "failing fs" is returned that returns errors for all operations.
func Fs() vfs.FileSystem {
	if fs == nil {
		var err error
		fs, err = open()
		if err != nil {
			fs = failfs(err)
		}
	}
	return fs
}

// find end-of-directory struct, near the end of the file.
// it specifies the size & offset of the central directory.
// we assume the central directory is located just before the end-of-central-directory.
// so that allows us to calculate the original size of the zip file.
// which in turn allows us to use godoc's zipfs to serve the zip file withend.
func open() (vfs.FileSystem, error) {
	f, err := os.Open(os.Args[0])
	if err != nil {
		return nil, err
	}
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}

	n := int64(65 * 1024)
	size := fi.Size()
	if size < n {
		n = size
	}
	buf := make([]byte, n)
	_, err = io.ReadAtLeast(io.NewSectionReader(f, size-n, n), buf, len(buf))
	if err != nil {
		return nil, err
	}
	o := int64(findSignatureInBlock(buf))
	if o < 0 {
		return nil, errors.New("could not locate zip file, no end-of-central-directory signature found")
	}
	cdirsize := int64(binary.LittleEndian.Uint32(buf[o+12:]))
	cdiroff := int64(binary.LittleEndian.Uint32(buf[o+16:]))
	zipsize := cdiroff + cdirsize + (int64(len(buf)) - o)

	rr := io.NewSectionReader(f, size-zipsize, zipsize)
	r, err := zip.NewReader(rr, zipsize)
	if err != nil {
		return nil, err
	}

	rc := &zip.ReadCloser{Reader: *r}
	return &assetFS{zipfs.New(rc, "<asset>"), rc}, nil
}

// Error returns a non-nil error if no asset could be found in the binary.
// For example when no zip file was appended to the binary.
func Error() error {
	switch fs := Fs().(type) {
	case *failFS:
		return fs.err
	}
	return nil
}

// Close the FileSystem, closing open files to the (zip file within) the binary.
func Close() {
	switch fs := Fs().(type) {
	case *assetFS:
		fs.rc.Close()
	}
	fs = nil
}
