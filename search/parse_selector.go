package search

import "github.com/remexre/go-parcom"

var selector = parcom.Map(parcom.Chain(
	parcom.Alt(selectors...),
	parcom.Tag(":"),
	parcom.Alt(
		parcom.Map(parcom.Chain(
			parcom.Tag("\""),
			parcom.Opt(parcom.AnyOfFunc(func(b byte) bool {
				return b != '"'
			}), ""),
			parcom.Tag("\""),
		), func(_, value, _ string) string {
			return value
		}),
		text,
	),
), func(selector, _, value string) []string {
	return []string{selector, value}
})
