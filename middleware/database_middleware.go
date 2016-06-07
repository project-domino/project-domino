package middleware

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/context"
	"github.com/jinzhu/gorm"
)

// DatabaseMiddleware adds the provided database handle to all requests.
func DatabaseMiddleware(db *gorm.DB) negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		context.Set(r, "db", db)
		next(w, r)
	}
}
