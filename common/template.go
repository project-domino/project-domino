package common

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"golang.org/x/tools/godoc/vfs"

	"github.com/gorilla/context"
)

// Views is a template containing all views.
var Views = template.New("Views").Funcs(template.FuncMap{
	"json": func(obj interface{}) (string, error) {
		b, err := json.Marshal(obj)
		return string(b), err
	},
})

// LoadTemplates loads the Views global with templates from the archive.
func LoadTemplates(fs vfs.FileSystem) error {
	files, err := fs.ReadDir("templates")
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		reader, err := fs.Open(path.Join("templates", file.Name()))
		if err != nil {
			return err
		}

		src, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}

		Views.Parse(string(src))
	}
	return nil
}

// ExecuteTemplate executes the given template from Views, with built-in error
// handling. The input variables are pulled from the context.
func ExecuteTemplate(w http.ResponseWriter, r *http.Request, template string) {
	if err := Views.ExecuteTemplate(w, template, context.GetAll(r)); err != nil {
		HandleError(w, err, http.StatusInternalServerError)
	}
}

// HandleError gives the standard response to an error.
func HandleError(w http.ResponseWriter, err error, code int) {
	log.Println(err)
	http.Error(w, err.Error(), code)
}

// TemplateHandler is a handler that renders a template from Views, with error
// handling. The input variables are pulled from the context.
func TemplateHandler(template string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ExecuteTemplate(w, r, template)
	}
}
