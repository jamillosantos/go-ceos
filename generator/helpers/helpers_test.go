package helpers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers", func() {
	Describe("camelCase", func() {
		It("should format an empty string", func() {
			Expect(CamelCase("")).To(BeEmpty())
		})

		It("should format a schema name", func() {
			Expect(CamelCase("User")).To(Equal("user"))
		})

		It("should format a schema name with multiple words", func() {
			Expect(CamelCase("UserGroup")).To(Equal("userGroup"))
		})
	})
})
