package naming_test

import (
	"github.com/jamillosantos/go-ceous/generator/naming"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Snake", func() {
	It("should format a name", func() {
		Expect(naming.SnakeCase.Do("Complex Name")).To(Equal("complex_name"))
		Expect(naming.SnakeCase.Do("Complex_Name")).To(Equal("complex_name"))
		Expect(naming.SnakeCase.Do("complexName")).To(Equal("complex_name"))
		Expect(naming.SnakeCase.Do("ComplexName")).To(Equal("complex_name"))
	})
})
