package search

import (
	"bytes"
	"strconv"
	"strings"
)

// Query represents a search query.
type Query struct {
	Selectors map[string][]string
	Text      []string
}

func (q *Query) String() string {
	buf := new(bytes.Buffer)
	for selector, vals := range q.Selectors {
		for _, val := range vals {
			buf.WriteString(selector)
			buf.WriteRune(':')
			if strings.Contains(val, " ") {
				buf.WriteString(strconv.Quote(val))
			} else {
				buf.WriteString(val)
			}
			buf.WriteRune(' ')
		}
	}
	buf.WriteString(strings.Join(q.Text, " "))
	return buf.String()
}
