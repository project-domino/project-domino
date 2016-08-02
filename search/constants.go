package search

// These constants are the valid search query values for type.
const (
	NoteSearchType       string = "note"
	CollectionSearchType        = "collection"
	UserSearchType              = "user"
	TagSearchType               = "tag"
	AllSearchType               = "all"
)

const queryFormat string = "searchtext @@ to_tsquery(?)"
