package search

import "github.com/remexre/go-parcom"

var whitespace = parcom.AnyOf(" \t\r\n")
var optionalWS = parcom.Opt(whitespace, "")
