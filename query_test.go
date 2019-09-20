package ceous_test

import (
	"github.com/jamillosantos/go-ceous"
	"github.com/jamillosantos/go-ceous/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Query", func() {
	Describe("BaseQuery", func() {
		BeforeEach(func() {
			tests.DBStart()
			tests.DBUsersCreate()
			tests.DBUsersInsertJoes()
		})

		AfterEach(func() {
			tests.DBStop()
		})

		It("it should retrieve a table", func() {
			rs, err := tests.DB.Query("SELECT * FROM users")
			Expect(err).ToNot(HaveOccurred())
			defer rs.Close()
			resultSet := ceous.NewRecordResultSet(rs)
			var u UserTestModel
			Expect(resultSet.Next()).To(BeTrue())
			Expect(resultSet.ToModel(&u)).ToNot(HaveOccurred())
			Expect(u.ID).To(Equal(1))
			Expect(u.Name).To(Equal("Snake Eyes"))
			Expect(resultSet.Next()).To(BeTrue())
			Expect(resultSet.ToModel(&u)).ToNot(HaveOccurred())
			Expect(u.ID).To(Equal(2))
			Expect(u.Name).To(Equal("Scarlet"))
			Expect(resultSet.Next()).To(BeTrue())
			Expect(resultSet.ToModel(&u)).ToNot(HaveOccurred())
			Expect(u.ID).To(Equal(3))
			Expect(u.Name).To(Equal("Tank"))
			Expect(resultSet.Next()).To(BeTrue())
			Expect(resultSet.ToModel(&u)).ToNot(HaveOccurred())
			Expect(u.ID).To(Equal(4))
			Expect(u.Name).To(Equal("Duke"))
			Expect(resultSet.Next()).To(BeFalse())
		})
	})
})
