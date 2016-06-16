package main

import (
	"archive/zip"
	"path"

	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
)

// NewZipFileSystem creates a vfs.FileSystem for assets from a .zip file.
func NewZipFileSystem(filePath string) (vfs.FileSystem, error) {
	reader, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}
	return zipfs.New(reader, path.Base(filePath)), nil
}
