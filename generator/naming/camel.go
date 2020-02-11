package naming

import (
	"strings"
)

type camelCase struct{}

// CamelCase format a given name as Camel case.
var CamelCase camelCase

// Do implements camcel casing a given `name`.
func (*camelCase) Do(name string) string {
	toks := splitRule.FindAllString(name, -1)
	var s strings.Builder
	var fnc func(string) string = strings.ToLower
	for i, tok := range toks {
		if i > 0 {
			fnc = strings.ToUpper
		}
		s.WriteString(fnc(tok[0:1]) + tok[1:])
	}
	return s.String()
}
