package search

// These constants are the valid search query values for type.
const (
	NoteSearchType       string = "note"
	CollectionSearchType        = "collection"
	UserSearchType              = "user"
	TagSearchType               = "tag"
	AllSearchType               = "all"
)

// MaxItems is the maximum number of items that can be returned
const MaxItems uint = 100

// MaxPage is the maximum page number which can be returned
const MaxPage uint = 10

const queryFormat string = "searchtext @@ to_tsquery(?)"
