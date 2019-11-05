package ceous_test

import (
	"github.com/jamillosantos/go-ceous/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Connection", func() {
	BeforeEach(func() {
		tests.DBStart()
		tests.DBUsersCreate()
		tests.DBUsersInsertJoes()
	})

	AfterEach(func() {
		tests.DBStop()
	})

	Describe("Begin", func() {
		It("should start a transaction", func() {
			t, err := tests.Default.Begin()
			Expect(err).ToNot(HaveOccurred())
			Expect(t).ToNot(BeNil())
		})

		It("should commit a transaction", func() {
			t, err := tests.Default.Begin()
			Expect(err).ToNot(HaveOccurred())
			_, err = t.Exec("delete from users")
			Expect(t.Commit()).To(Succeed())
			Expect(err).ToNot(HaveOccurred())

			users, err := tests.Default.UserQuery().All()
			Expect(err).ToNot(HaveOccurred())
			Expect(users).To(HaveLen(0))
		})

		It("should rollback a transaction", func() {
			t, err := tests.Default.Begin()
			Expect(err).ToNot(HaveOccurred())
			_, err = t.Exec("delete from users")
			count, err := t.UserQuery().Count()
			Expect(err).ToNot(HaveOccurred())
			Expect(count).To(Equal(int64(0)))
			Expect(t.Rollback()).To(Succeed())
			Expect(err).ToNot(HaveOccurred())

			users, err := tests.Default.UserQuery().All()
			Expect(err).ToNot(HaveOccurred())
			Expect(users).To(HaveLen(4))
		})
	})
})
