package reporters

type Reporter interface {
	Linef(fmt string, args ...interface{})
	Line(args ...interface{})
}
