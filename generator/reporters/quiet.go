package reporters

// Quiet is a reporter that do not output anything.
type Quiet struct {
}

func (q *Quiet) Linef(fmt string, args ...interface{}) {
}

func (q *Quiet) Line(args ...interface{}) {
}
