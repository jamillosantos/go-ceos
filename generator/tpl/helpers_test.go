package tpl

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers", func() {
	Describe("camelCase", func() {
		It("should format an empty string", func() {
			Expect(camelCase("")).To(BeEmpty())
		})

		It("should format a schema name", func() {
			Expect(camelCase("User")).To(Equal("user"))
		})

		It("should format a schema name with multiple words", func() {
			Expect(camelCase("UserGroup")).To(Equal("userGroup"))
		})
	})
})
