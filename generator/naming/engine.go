package naming

type Engine interface {
	Do(name string) string
}
