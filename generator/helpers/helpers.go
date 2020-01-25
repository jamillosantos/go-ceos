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

func AppendStringIfNotEmpty(arr []string, data ...string) []string {
	for _, s := range data {
		if s != "" {
			arr = append(arr, s)
		}
	}
	return arr
}

const Pointer = "*"
