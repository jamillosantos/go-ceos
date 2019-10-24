package ceous_test

import (
	"github.com/jamillosantos/go-ceous"
	"github.com/jamillosantos/go-ceous/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store", func() {
	Describe("BaseStore", func() {
		Context("Insert", func() {
			BeforeEach(func() {
				tests.DBStart()
				tests.DBUsersCreate()
			})

			AfterEach(func() {
				tests.DBStop()
			})

			It("should insert a user specifying not fields", func() {
				db := ceous.WithDB(tests.DB)

				user := tests.User{
					Name:     "Snake Eyes",
					Password: "12345",
					Role:     "stealth",
				}

				userStore := tests.NewUserStore(db)
				Expect(userStore.Insert(&user)).To(Succeed())

				b, err := tests.NewUserQuery(db).Builder()
				Expect(err).ToNot(HaveOccurred())
				rs := tests.NewUserResultSet(b.RunWith(tests.DB).Query())
				Expect(rs.Next()).To(BeTrue())
				var userFound tests.User
				Expect(rs.ToModel(&userFound)).To(Succeed())
				Expect(userFound.Name).To(Equal("Snake Eyes"))
				Expect(userFound.Password).To(Equal("12345"))
				Expect(userFound.Role).To(Equal("stealth"))
				Expect(rs.Next()).To(BeFalse())
			})

			It("should insert a user specifying fields", func() {
				db := ceous.WithDB(tests.DB)

				user := tests.User{
					Name: "Snake Eyes",
				}

				userStore := tests.NewUserStore(db)
				Expect(userStore.Insert(&user, tests.Schema.User.Name)).To(Succeed())

				b, err := tests.NewUserQuery(db).Builder()
				Expect(err).ToNot(HaveOccurred())
				rs := tests.NewUserResultSet(b.RunWith(tests.DB).Query())
				Expect(rs.Next()).To(BeTrue())
				var userFound tests.User
				Expect(rs.ToModel(&userFound)).To(Succeed())
				Expect(userFound.Name).To(Equal("Snake Eyes"))
				Expect(userFound.Password).To(Equal(""))
				Expect(userFound.Role).To(Equal(""))
				Expect(rs.Next()).To(BeFalse())
			})
		})
	})
})
