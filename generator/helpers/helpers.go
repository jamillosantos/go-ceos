package helpers

import "strings"

func CamelCase(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToLower(name[0:1]) + name[1:]
}

func PascalCase(name string) string {
	if name == "" {
		return ""
	}
	return strings.ToUpper(name[0:1]) + name[1:]
}

const Pointer = "*"
