package main

import (
	"archive/zip"
	"html/template"
	"io/ioutil"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"golang.org/x/tools/godoc/vfs"
	"golang.org/x/tools/godoc/vfs/httpfs"
	"golang.org/x/tools/godoc/vfs/zipfs"
)

// SetupAssets detects the location assets and templates should be loaded from,
// then loads them into a router.
func SetupAssets(r *gin.Engine) error {
	assetFS, err := GetAssetFileSystem()
	if err != nil {
		return err
	}

	templates, err := GetTemplates(assetFS)
	if err != nil {
		return err
	}

	r.SetHTMLTemplate(templates)
	r.StaticFS("/assets/", httpfs.New(assetFS))
	return nil
}

// GetTemplates loads templates from a vfs.FileSystem.
func GetTemplates(fs vfs.FileSystem) (*template.Template, error) {
	allFiles, err := fs.ReadDir("/")
	if err != nil {
		return nil, err
	}
	t := template.New("")
	for _, file := range allFiles {
		if file.IsDir() || path.Ext(file.Name()) != ".html" {
			continue
		}

		reader, err := fs.Open("/" + file.Name())
		if err != nil {
			return nil, err
		}

		src, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		t.New(file.Name()).Parse(string(src))
	}
	return t, nil
}

// GetAssetFileSystem returns a vfs.FileSystem containing the assets and
// templates.
func GetAssetFileSystem() (vfs.FileSystem, error) {
	if viper.GetBool("assets.dev") {
		return vfs.OS("assets/dist"), nil
	}
	return NewZipFileSystem(viper.GetString("assets.path"))
}

// NewZipFileSystem creates a vfs.FileSystem for assets from a .zip file.
func NewZipFileSystem(filePath string) (vfs.FileSystem, error) {
	reader, err := zip.OpenReader(filePath)
	if err != nil {
		return nil, err
	}
	return zipfs.New(reader, path.Base(filePath)), nil
}
