package helpers

import "strings"

func CamelCase(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToLower(name[0:1]) + name[1:]
}
