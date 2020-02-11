package helpers

func AppendStringIfNotEmpty(arr []string, data ...string) []string {
	for _, s := range data {
		if s != "" {
			arr = append(arr, s)
		}
	}
	return arr
}

const Pointer = "*"
