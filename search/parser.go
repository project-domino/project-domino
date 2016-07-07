package search

import (
	"errors"

	"github.com/remexre/go-parcom"
)

var parser = parcom.Map(
	parcom.Many0(parcom.Map(parcom.Chain(
		optionalWS,
		parcom.Alt(
			selector,
			text,
		),
		optionalWS,
	), func(_, i, _ interface{}) interface{} {
		return i
	})),
	func(in []interface{}) *Query {
		q := &Query{make(map[string]string), nil}
		for _, component := range in {
			switch c := component.(type) {
			case string:
				q.Text = append(q.Text, c)
			case []string:
				q.Selectors[c[0]] = c[1]
			}
		}
		return q
	},
)

// ParseQuery parses a Query.
func ParseQuery(in string) (*Query, error) {
	rem, qIface, ok := parser(in)
	if !ok || len(rem) > 0 {
		return nil, errors.New("failed to parse")
	}
	q, ok := qIface.(*Query)
	if !ok {
		return nil, errors.New("parser returned non-Query value")
	}
	return q, nil
}
