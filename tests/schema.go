package tests

import (
	sq "github.com/elgris/sqrl"
)

func init() {
	// TODO(jota): This should be gone in the when finishing the library.
	sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
}
