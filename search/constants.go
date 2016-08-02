package search

// These constants are the valid search query values for type.
const (
	NoteSearchType       string = "note"
	CollectionSearchType        = "collection"
	UserSearchType              = "user"
	TagSearchType               = "tag"
	AllSearchType               = "all"
)

// These constants are valid selectors
const (
	AuthorSelector string = "author"
	TagSelector           = "tag"
)

const queryFormat string = "searchtext @@ to_tsquery(?)"
