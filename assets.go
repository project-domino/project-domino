package main

import (
	"archive/zip"
	"path"

	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
)

// NewAssetFilesystem creates a vfs.Filesystem for assets.
func NewAssetFilesystem(zipFilePath string) (vfs.FileSystem, error) {
	reader, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return nil, err
	}
	return zipfs.New(reader, path.Base(zipFilePath)), nil
}
