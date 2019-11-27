package reporters

import (
	"fmt"
	"os"
)

// Verbose enable the output logging.
type Verbose struct {
}

func (q *Verbose) Linef(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}

func (q *Verbose) Line(args ...interface{}) {
	for _, arg := range args {
		fmt.Fprint(os.Stderr, arg)
	}
	fmt.Fprintln(os.Stderr)
}
