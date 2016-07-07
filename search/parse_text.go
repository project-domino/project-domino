package search

import "github.com/remexre/go-parcom"

var text = parcom.AnyOfFunc(func(b byte) bool {
	return b != ' ' && b != '\t' && b != '\r' && b != '\n'
})
