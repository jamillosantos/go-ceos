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
				tests.DBUserGroupsCreate()
			})

			AfterEach(func() {
				tests.DBStop()
			})

			It("should insert a user specifying no fields", func() {
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

			It("should insert multiple users returning the PK", func() {
				db := ceous.WithDB(tests.DB)

				user1 := tests.User{
					Name:     "Snake Eyes",
					Password: "12345",
					Role:     "stealth",
				}
				user2 := tests.User{
					Name:     "Scarlet",
					Password: "54321",
					Role:     "intelligence",
				}

				userStore := tests.NewUserStore(db)
				Expect(userStore.Insert(&user1)).To(Succeed())
				Expect(user1.ID).To(Equal(1))

				Expect(userStore.Insert(&user2)).To(Succeed())
				Expect(user2.ID).To(Equal(2))

				b, err := tests.NewUserQuery(db).Builder()
				Expect(err).ToNot(HaveOccurred())
				rs := tests.NewUserResultSet(b.RunWith(tests.DB).Query())
				Expect(rs.Next()).To(BeTrue())
				var userFound tests.User
				Expect(rs.ToModel(&userFound)).To(Succeed())
				Expect(userFound.Name).To(Equal("Snake Eyes"))
				Expect(userFound.Password).To(Equal("12345"))
				Expect(userFound.Role).To(Equal("stealth"))
				Expect(rs.Next()).To(BeTrue())
				Expect(rs.ToModel(&userFound)).To(Succeed())
				Expect(userFound.Name).To(Equal("Scarlet"))
				Expect(userFound.Password).To(Equal("54321"))
				Expect(userFound.Role).To(Equal("intelligence"))
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

			FIt("should insert a model with a composite PK", func() {
				db := ceous.WithDB(tests.DB)

				userGroup := tests.UserGroup{
					ID: tests.UserGroupPK{
						UserID:  1,
						GroupID: 2,
					},
					Admin: true,
				}

				userGroupStore := tests.NewUserGroupStore(db)
				Expect(userGroupStore.Insert(&userGroup)).To(Succeed())

				b, err := tests.NewUserGroupQuery(db).Builder()
				Expect(err).ToNot(HaveOccurred())
				rs := tests.NewUserGroupResultSet(b.RunWith(tests.DB).Query())
				Expect(rs.Next()).To(BeTrue())
				var userGroupFound tests.UserGroup
				Expect(rs.ToModel(&userGroupFound)).To(Succeed())
				Expect(userGroupFound.ID.UserID).To(Equal(1))
				Expect(userGroupFound.ID.GroupID).To(Equal(2))
				Expect(userGroupFound.Admin).To(BeTrue())
				Expect(rs.Next()).To(BeFalse())
			})
		})

		Context("Update", func() {
			BeforeEach(func() {
				tests.DBStart()
				tests.DBUsersCreate()
				tests.DBUsersInsertJoes()
				tests.DBUserGroupsCreate()
				tests.DBUserGroupsInsert()
			})

			AfterEach(func() {
				tests.DBStop()
			})

			It("should update a user not specifying fields", func() {
				db := ceous.WithDB(tests.DB)

				user, err := tests.NewUserQuery(db).ByID(1).One()
				Expect(err).ToNot(HaveOccurred())

				store := tests.NewUserStore(db)
				user.Name = "Snake Eyes 02"
				user.Password = "67890"
				user.Role = "stealth 02"
				n, err := store.Update(&user)
				Expect(err).ToNot(HaveOccurred())
				Expect(n).To(Equal(int64(1)))

				userFound, err := tests.NewUserQuery(db).ByID(1).One()
				Expect(err).ToNot(HaveOccurred())
				Expect(userFound.Name).To(Equal("Snake Eyes 02"))
				Expect(userFound.Password).To(Equal("67890"))
				Expect(userFound.Role).To(Equal("stealth 02"))
			})

			It("should update a user specifying fields", func() {
				db := ceous.WithDB(tests.DB)

				user, err := tests.NewUserQuery(db).ByID(1).One()
				Expect(err).ToNot(HaveOccurred())

				store := tests.NewUserStore(db)
				user.Name = "Snake Eyes 02"
				user.Password = "67890"
				user.Role = "stealth 02"
				n, err := store.Update(&user, tests.Schema.User.Name)
				Expect(err).ToNot(HaveOccurred())
				Expect(n).To(Equal(int64(1)))

				userFound, err := tests.NewUserQuery(db).ByID(1).One()
				Expect(err).ToNot(HaveOccurred())
				Expect(userFound.Name).To(Equal("Snake Eyes 02"))
				Expect(userFound.Password).To(Equal(""))
				Expect(userFound.Role).To(Equal(""))
			})

			It("should update a model with composite PK", func() {
				db := ceous.WithDB(tests.DB)

				pk := tests.UserGroupPK{
					UserID:  1,
					GroupID: 2,
				}

				userGroup, err := tests.NewUserGroupQuery(db).ByID(pk).One()
				Expect(err).ToNot(HaveOccurred())

				store := tests.NewUserGroupStore(db)
				userGroup.Admin = true
				n, err := store.Update(&userGroup, tests.Schema.UserGroup.Admin)
				Expect(err).ToNot(HaveOccurred())
				Expect(n).To(Equal(int64(1)))

				userGroupFound, err := tests.NewUserGroupQuery(db).ByID(pk).One()
				Expect(err).ToNot(HaveOccurred())
				Expect(userGroupFound.ID.UserID).To(Equal(1))
				Expect(userGroupFound.ID.GroupID).To(Equal(2))
				Expect(userGroupFound.Admin).To(BeTrue())
			})

			PIt("should fail updating a non existing model")
		})

		Context("Delete", func() {
			BeforeEach(func() {
				tests.DBStart()
				tests.DBUsersCreate()
				tests.DBUsersInsertJoes()
				tests.DBUserGroupsCreate()
				tests.DBUserGroupsInsert()
			})

			AfterEach(func() {
				tests.DBStop()
			})

			It("should delete a user", func() {
				db := ceous.WithDB(tests.DB)

				user, err := tests.NewUserQuery(db).ByID(1).One()
				Expect(err).ToNot(HaveOccurred())

				store := tests.NewUserStore(db)
				Expect(store.Delete(&user)).To(Succeed())

				_, err = tests.NewUserQuery(db).ByID(1).One()
				Expect(err).To(Equal(ceous.ErrNotFound))
			})

			It("should delete a model with composite PK", func() {
				db := ceous.WithDB(tests.DB)

				pk := tests.UserGroupPK{
					UserID:  1,
					GroupID: 2,
				}

				userGroup, err := tests.NewUserGroupQuery(db).ByID(pk).One()
				Expect(err).ToNot(HaveOccurred())

				store := tests.NewUserGroupStore(db)
				Expect(store.Delete(&userGroup)).To(Succeed())

				_, err = tests.NewUserGroupQuery(db).ByID(pk).One()
				Expect(err).To(Equal(ceous.ErrNotFound))
			})

			PIt("should fail deleting a non existing model")
		})
	})
})
