package tpl

import "strings"

func camelCase(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToLower(name[0:1]) + name[1:]
}
