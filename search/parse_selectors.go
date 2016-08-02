package search

import "github.com/remexre/go-parcom"

var selectors = []parcom.Parser{
	parcom.Tag(AuthorSelector),
	parcom.Tag(TagSelector),
}
