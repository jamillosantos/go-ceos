package naming

import (
	"regexp"
	"strings"
)

// SnakeCase format a given name as snake case.
type snakeCase struct{}

// SnakeCase is the engine for formating names in snake_case format.
var SnakeCase snakeCase

var splitRule = regexp.MustCompile("[a-z]+|[A-Z][a-z]+|[A-Z]+")

// Do implements snake casing a given `name`.
func (*snakeCase) Do(name string) string {
	toks := splitRule.FindAllString(name, -1)
	var s strings.Builder
	for i, tok := range toks {
		if i > 0 {
			s.WriteString("_")
		}
		s.WriteString(strings.ToLower(tok))
	}
	return s.String()
}
