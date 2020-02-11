package parser

import (
	"strings"

	"github.com/jamillosantos/go-ceous/generator/naming"
)

// columnPrefix serializes the prefix array into a string by concating the
// strings dividing the fields by _.
func columnPrefix(prefix []string) string {
	r := ""
	for _, s := range prefix {
		r = r + s + "_"
	}
	return r
}

// namePrefix serializes the prefix array into a string by concating the
// strings pascal casing all names.
func namePrefix(prefix []string) string {
	r := ""
	for _, s := range prefix {
		r = r + naming.PascalCase.Do(s)
	}
	return r
}

func memberAccess(path ...string) string {
	return strings.Join(path, ".")
}
