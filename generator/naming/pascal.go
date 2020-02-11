package naming

import (
	"regexp"
	"strings"

	"github.com/jamillosantos/go-ceous/generator/helpers"
)

type pascalCase struct{}

// PascalCase format a given name as Pascal case.
var PascalCase pascalCase

var pascalSplitRule = regexp.MustCompile("[a-z]+|[A-Z][a-z]+|[A-Z]*")

// Do implements pascal casing a given `name`.
func (*pascalCase) Do(name string) string {
	toks := splitRule.FindAllString(name, -1)
	var s strings.Builder
	for _, tok := range toks {
		s.WriteString(helpers.PascalCase(tok))
	}
	return s.String()
}
