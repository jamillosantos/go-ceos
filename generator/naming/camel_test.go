package naming_test

import (
	"github.com/jamillosantos/go-ceous/generator/naming"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Camel", func() {
	It("should format a name", func() {
		Expect(naming.CamelCase.Do("Complex Name")).To(Equal("complexName"))
		Expect(naming.CamelCase.Do("Complex_Name")).To(Equal("complexName"))
		Expect(naming.CamelCase.Do("complexName")).To(Equal("complexName"))
		Expect(naming.CamelCase.Do("ComplexName")).To(Equal("complexName"))
	})
})
