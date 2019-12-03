package ceous_test

import (
	"github.com/jamillosantos/go-ceous"
	"github.com/jamillosantos/go-ceous/tests"
	"github.com/jamillosantos/go-ceous/tests/db"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Store", func() {
	Describe("BaseStore", func() {
		Context("Insert", func() {
			BeforeEach(func() {
				db.DBStart()
				db.DBUsersCreate()
				db.DBUserGroupsCreate()
			})

			AfterEach(func() {
				db.DBStop()
			})

			It("should insert a user specifying no fields", func() {
				user := tests.User{
					Name:     "Snake Eyes",
					Password: "12345",
					Role:     "stealth",
				}

				userStore := db.Default.UserStore()
				Expect(userStore.Insert(&user)).To(Succeed())

				b, err := db.Default.UserQuery().Builder()
				Expect(err).ToNot(HaveOccurred())
				rs := db.NewUserResultSet(b.RunWith(db.DB).Query())
				Expect(rs.Next()).To(BeTrue())
				var userFound tests.User
				Expect(rs.ToModel(&userFound)).To(Succeed())
				Expect(userFound.Name).To(Equal("Snake Eyes"))
				Expect(userFound.Password).To(Equal("12345"))
				Expect(userFound.Role).To(Equal("stealth"))
				Expect(rs.Next()).To(BeFalse())
			})

			It("should insert multiple users returning the PK", func() {
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

				userStore := db.Default.UserStore()
				Expect(userStore.Insert(&user1)).To(Succeed())
				Expect(user1.ID).To(Equal(1))

				Expect(userStore.Insert(&user2)).To(Succeed())
				Expect(user2.ID).To(Equal(2))

				b, err := db.Default.UserQuery().Builder()
				Expect(err).ToNot(HaveOccurred())
				rs := db.NewUserResultSet(b.RunWith(db.DB).Query())
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
				user := tests.User{
					Name: "Snake Eyes",
				}

				userStore := db.Default.UserStore()
				Expect(userStore.Insert(&user, tests.Schema.User.Name)).To(Succeed())

				b, err := db.Default.UserQuery().Builder()
				Expect(err).ToNot(HaveOccurred())
				rs := db.NewUserResultSet(b.RunWith(db.DB).Query())
				Expect(rs.Next()).To(BeTrue())
				var userFound tests.User
				Expect(rs.ToModel(&userFound)).To(Succeed())
				Expect(userFound.Name).To(Equal("Snake Eyes"))
				Expect(userFound.Password).To(Equal(""))
				Expect(userFound.Role).To(Equal(""))
				Expect(rs.Next()).To(BeFalse())
			})

			It("should insert a model with a composite PK", func() {
				userGroup := tests.UserGroup{
					ID: tests.UserGroupPK{
						UserID:  1,
						GroupID: 2,
					},
					Admin: true,
				}

				userGroupStore := db.Default.UserGroupStore()
				Expect(userGroupStore.Insert(&userGroup)).To(Succeed())

				b, err := db.Default.UserGroupQuery().Builder()
				Expect(err).ToNot(HaveOccurred())
				rs := db.NewUserGroupResultSet(b.RunWith(db.DB).Query())
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
				db.DBStart()
				db.DBUsersCreate()
				db.DBUsersInsertJoes()
				db.DBUserGroupsCreate()
				db.DBUserGroupsInsert()
			})

			AfterEach(func() {
				db.DBStop()
			})

			It("should update a user not specifying fields", func() {
				user, err := db.Default.UserQuery().ByID(1).One()
				Expect(err).ToNot(HaveOccurred())

				store := db.Default.UserStore()
				user.Name = "Snake Eyes 02"
				user.Password = "67890"
				user.Role = "stealth 02"
				n, err := store.Update(&user)
				Expect(err).ToNot(HaveOccurred())
				Expect(n).To(Equal(int64(1)))

				userFound, err := db.Default.UserQuery().ByID(1).One()
				Expect(err).ToNot(HaveOccurred())
				Expect(userFound.Name).To(Equal("Snake Eyes 02"))
				Expect(userFound.Password).To(Equal("67890"))
				Expect(userFound.Role).To(Equal("stealth 02"))
			})

			It("should update a user specifying fields", func() {
				user, err := db.Default.UserQuery().ByID(1).One()
				Expect(err).ToNot(HaveOccurred())

				store := db.Default.UserStore()
				user.Name = "Snake Eyes 02"
				user.Password = "67890"
				user.Role = "stealth 02"
				n, err := store.Update(&user, tests.Schema.User.Name)
				Expect(err).ToNot(HaveOccurred())
				Expect(n).To(Equal(int64(1)))

				userFound, err := db.Default.UserQuery().ByID(1).One()
				Expect(err).ToNot(HaveOccurred())
				Expect(userFound.Name).To(Equal("Snake Eyes 02"))
				Expect(userFound.Password).To(Equal(""))
				Expect(userFound.Role).To(Equal(""))
			})

			PIt("should update a model with composite PK", func() {
				/*
					pk := tests.UserGroupPK{
						UserID:  1,
						GroupID: 2,
					}

					userGroup, err := db.Default.UserGroupQuery().ByID(pk).One()
					Expect(err).ToNot(HaveOccurred())

					store := db.Default.UserGroupStore()
					userGroup.Admin = true
					n, err := store.Update(&userGroup, tests.Schema.UserGroup.Admin)
					Expect(err).ToNot(HaveOccurred())
					Expect(n).To(Equal(int64(1)))

					userGroupFound, err := db.Default.UserGroupQuery().ByID(pk).One()
					Expect(err).ToNot(HaveOccurred())
					Expect(userGroupFound.ID.UserID).To(Equal(1))
					Expect(userGroupFound.ID.GroupID).To(Equal(2))
					Expect(userGroupFound.Admin).To(BeTrue())
				*/
			})

			PIt("should fail updating a non existing model")
		})

		Context("Delete", func() {
			BeforeEach(func() {
				db.DBStart()
				db.DBUsersCreate()
				db.DBUsersInsertJoes()
				db.DBUserGroupsCreate()
				db.DBUserGroupsInsert()
			})

			AfterEach(func() {
				db.DBStop()
			})

			It("should delete a user", func() {
				user, err := db.Default.UserQuery().ByID(1).One()
				Expect(err).ToNot(HaveOccurred())

				store := db.Default.UserStore()
				Expect(store.Delete(&user)).To(Succeed())

				_, err = db.Default.UserQuery().ByID(1).One()
				Expect(err).To(Equal(ceous.ErrNotFound))
			})

			PIt("should delete a model with composite PK", func() {
				/*
					userGroup, err := db.Default.UserGroupQuery().ByIDUserID(1).ByIDGroupID(2).One()
					Expect(err).ToNot(HaveOccurred())

					store := db.Default.UserGroupStore()
					Expect(store.Delete(&userGroup)).To(Succeed())

					_, err = db.Default.UserGroupQuery().ByIDUserID(1).ByIDGroupID(2).One()
					Expect(err).To(Equal(ceous.ErrNotFound))
				*/
			})

			PIt("should fail deleting a non existing model")
		})
	})
})
