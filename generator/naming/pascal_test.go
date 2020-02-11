package naming_test

import (
	"github.com/jamillosantos/go-ceous/generator/naming"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pascal", func() {
	It("should format a name", func() {
		Expect(naming.PascalCase.Do("Complex Name")).To(Equal("ComplexName"))
		Expect(naming.PascalCase.Do("Complex_Name")).To(Equal("ComplexName"))
		Expect(naming.PascalCase.Do("complexName")).To(Equal("ComplexName"))
		Expect(naming.PascalCase.Do("ComplexName")).To(Equal("ComplexName"))
	})
})
