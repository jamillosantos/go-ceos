package ceous_test

import (
	"github.com/jamillosantos/go-ceous"
	"github.com/jamillosantos/go-ceous/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Operators", func() {
	Describe("Eq", func() {
		It("should generate the sql for field = value", func() {
			d := ceous.Eq(tests.Schema.User.Name, "Snake Eyes")
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).ToNot(HaveOccurred())
			Expect(args).To(ConsistOf("Snake Eyes"))
			Expect(sql).To(Equal("name = ?"))
		})

		It("should generate the sql for field = NULL", func() {
			d := ceous.Eq(tests.Schema.User.Name, nil)
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).ToNot(HaveOccurred())
			Expect(args).To(BeEmpty())
			Expect(sql).To(Equal("name IS NULL"))
		})

		It("should generate the sql for field <> value (driver.Valuer)", func() {
			d := ceous.Eq(tests.Schema.User.Name, tests.ErrorString("this is a test"))
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("valuer with error"))
			Expect(args).To(BeEmpty())
			Expect(sql).To(BeEmpty())
		})

		It("should generate the sql for field = value (driver.Valuer)", func() {
			d := ceous.Eq(tests.Schema.User.Name, tests.CapitalString("this is a test"))
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).ToNot(HaveOccurred())
			Expect(args).To(ConsistOf("THIS IS A TEST"))
			Expect(sql).To(Equal("name = ?"))
		})

		It("should generate the sql for field IN values", func() {
			d := ceous.Eq(tests.Schema.User.ID, []interface{}{1, 2, 3})
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).ToNot(HaveOccurred())
			Expect(args).To(ConsistOf(1, 2, 3))
			Expect(sql).To(Equal("id IN (?,?,?)"))
		})

		It("should generate the sql for an empty list", func() {
			d := ceous.Eq(tests.Schema.User.ID, []interface{}{})
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).ToNot(HaveOccurred())
			Expect(args).To(BeEmpty())
			Expect(sql).To(Equal("(1=0)"))
		})
	})

	Describe("Ne", func() {
		It("should generate the sql for field <> value", func() {
			d := ceous.Ne(tests.Schema.User.Name, "Snake Eyes")
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).ToNot(HaveOccurred())
			Expect(args).To(ConsistOf("Snake Eyes"))
			Expect(sql).To(Equal("name <> ?"))
		})

		It("should generate the sql for field <> NULL", func() {
			d := ceous.Ne(tests.Schema.User.Name, nil)
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).ToNot(HaveOccurred())
			Expect(args).To(BeEmpty())
			Expect(sql).To(Equal("name IS NOT NULL"))
		})

		It("should generate the sql for field <> value (driver.Valuer)", func() {
			d := ceous.Ne(tests.Schema.User.Name, tests.ErrorString("this is a test"))
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("valuer with error"))
			Expect(args).To(BeEmpty())
			Expect(sql).To(BeEmpty())
		})

		It("should generate the sql for field !IN values", func() {
			d := ceous.Ne(tests.Schema.User.ID, []interface{}{1, 2, 3})
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).ToNot(HaveOccurred())
			Expect(args).To(ConsistOf(1, 2, 3))
			Expect(sql).To(Equal("id NOT IN (?,?,?)"))
		})

		It("should generate the sql for an empty list", func() {
			d := ceous.Ne(tests.Schema.User.ID, []interface{}{})
			sql, args, err := d(tests.Schema.User).ToSql()
			Expect(err).ToNot(HaveOccurred())
			Expect(args).To(BeEmpty())
			Expect(sql).To(Equal("(1=1)"))
		})
	})
})
