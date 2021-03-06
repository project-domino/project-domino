package vars

import "github.com/gin-gonic/gin"

// KeysList is a list of keys which are safe to send to the user. See Keys.
var KeysList = []string{
	"loggedIn",
	"user",
	"pageUser",
	"note",
	"noteJSON",
	"noteHTML",
	"collection",
	"collectionJSON",
	"query",
	"searchResult",
	"currentPage",
	"nextPage",
	"prevPage",
	"items",
	"collections",
	"collectionsJSON",
	"notes",
	"notesJSON",
	"featuredNotes",
	"featuredCollections",
}

// Keys returns the variables which have been set via gin.Context.Set. This is
// safer than simply passing c.Keys, as it ensures that no undesired values are
// accidentally passed to the user by using a whitelist, KeysList.
func Keys(c *gin.Context) Vars {
	vars := New()
	for _, key := range KeysList {
		vars[key], _ = c.Get(key)
	}
	return vars
}
