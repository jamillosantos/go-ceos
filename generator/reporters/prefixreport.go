package reporters

type prefixReporter struct {
	prefix   string
	reporter Reporter
}

// WithPrefix creates a new reporter with a prefix for all lines.
func WithPrefix(reporter Reporter, prefix string) Reporter {
	return &prefixReporter{
		prefix:   prefix,
		reporter: reporter,
	}
}

// Linef calls the reporter `Linef` implemenation with a prefix.
func (reporter *prefixReporter) Linef(fmt string, args ...interface{}) {
	reporter.reporter.Linef(reporter.prefix+fmt, args...)
}

// Line calls the reporter `Line` implemenation with a prefix.
func (reporter *prefixReporter) Line(args ...interface{}) {
	args = append([]interface{}{reporter.prefix}, args...)
	reporter.reporter.Line(args...)
}
