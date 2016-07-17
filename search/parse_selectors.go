package search

import "github.com/remexre/go-parcom"

var selectors = []parcom.Parser{
	parcom.Tag("author"),
	parcom.Tag("type"),
	parcom.Tag("university"),
	parcom.Tag("tag"),
}
