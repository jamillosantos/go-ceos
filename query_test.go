package ceous_test

import (
	sq "github.com/elgris/sqrl"
	"github.com/jamillosantos/go-ceous"
	"github.com/jamillosantos/go-ceous/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Query", func() {
	Describe("BaseQuery", func() {
		Context("SQL Generation", func() {
			Context("Select Fields", func() {
				It("should select all fields", func() {
					q := tests.NewUserQuery()
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, name, password, role, created_at, updated_at FROM users"))
					Expect(args).To(BeEmpty())
				})

				It("should not change the select fields when not providing fields", func() {
					q := tests.NewUserQuery().Select( /* no fields specified. */ )
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, name, password, role, created_at, updated_at FROM users"))
					Expect(args).To(BeEmpty())
				})

				It("should select specified fields", func() {
					q := tests.NewUserQuery(ceous.WithDB(tests.DB)).Select(tests.Schema.User.Name, tests.Schema.User.CreatedAt)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT name, created_at FROM users"))
					Expect(args).To(BeEmpty())
				})

				It("should select specified fields calling Select multiple times", func() {
					q := tests.NewUserQuery().Select(tests.Schema.User.ID).Select(tests.Schema.User.Name).Select(tests.Schema.User.CreatedAt)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, name, created_at FROM users"))
					Expect(args).To(BeEmpty())
				})

				It("should not select excluded fields", func() {
					q := tests.NewUserQuery().ExcludeFields(tests.Schema.User.Name, tests.Schema.User.CreatedAt)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, password, role, updated_at FROM users"))
					Expect(args).To(BeEmpty())
				})

				It("should not select excluded fields calling ExcludeFields muiltiple times", func() {
					q := tests.NewUserQuery().ExcludeFields(tests.Schema.User.Name).ExcludeFields(tests.Schema.User.CreatedAt)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, password, role, updated_at FROM users"))
					Expect(args).To(BeEmpty())
				})

				It("should not change select when calling ExcludeFields with no fields defined", func() {
					q := tests.NewUserQuery().ExcludeFields()
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, name, password, role, created_at, updated_at FROM users"))
					Expect(args).To(BeEmpty())
				})

				It("should select specific fields excluding others", func() {
					q := tests.NewUserQuery().
						Select(tests.Schema.User.ID, tests.Schema.User.Name, tests.Schema.User.Password).
						ExcludeFields(tests.Schema.User.Password)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, name FROM users"))
					Expect(args).To(BeEmpty())
				})
			})

			Context("Alias", func() {
				It("should select a field with a different alias", func() {
					userAlias := tests.Schema.User.As("usr")
					u := ceous.FieldAlias(userAlias)
					q := tests.NewUserQuery().Select(u(tests.Schema.User.ID), tests.Schema.User.Name)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT usr.id, name FROM users"))
					Expect(args).To(BeEmpty())
				})

				It("should select a field with a different alias", func() {
					userAlias := tests.Schema.User.As("usr")
					u := ceous.FieldAlias(userAlias)
					q := tests.NewUserQuery().Select(u(tests.Schema.User.ID), tests.Schema.User.Name)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT usr.id, name FROM users"))
					Expect(args).To(BeEmpty())
				})
			})

			Context("Limit + Offset", func() {
				It("should limit a query", func() {
					q := tests.NewUserQuery().Limit(1)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, name, password, role, created_at, updated_at FROM users LIMIT 1"))
					Expect(args).To(BeEmpty())
				})

				It("should offset a query", func() {
					q := tests.NewUserQuery().Offset(2)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, name, password, role, created_at, updated_at FROM users OFFSET 2"))
					Expect(args).To(BeEmpty())
				})

				It("should limit and offset a query", func() {
					q := tests.NewUserQuery().Limit(1).Offset(2)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, name, password, role, created_at, updated_at FROM users LIMIT 1 OFFSET 2"))
					Expect(args).To(BeEmpty())
				})

				It("should change limit and offset after building", func() {
					q := tests.NewUserQuery().Limit(1).Offset(2)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())

					// Change the query afterwards.
					q.Limit(3).Offset(4)

					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id, name, password, role, created_at, updated_at FROM users LIMIT 3 OFFSET 4"))
					Expect(args).To(BeEmpty())
				})
			})

			Context("Where", func() {
				It("should generate a where one condition", func() {
					q := tests.NewUserQuery().Select(tests.Schema.User.ID).ByID(1)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id FROM users WHERE id = ?"))
					Expect(args).To(ConsistOf(1))
				})

				It("should generate a where with multiple conditions", func() {
					q := tests.NewUserQuery().Select(tests.Schema.User.ID).ByID(1).ByName("Snake Eyes")
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id FROM users WHERE id = ? AND name = ?"))
					Expect(args).To(ConsistOf(1, "Snake Eyes"))
				})

				It("should generate a where with string conditions", func() {
					q := tests.NewUserQuery().Select(tests.Schema.User.ID).Where("LENGTH(password) < 6")
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id FROM users WHERE LENGTH(password) < 6"))
					Expect(args).To(BeEmpty())
				})

				It("should generate a where with string conditions with args", func() {
					q := tests.NewUserQuery().Select(tests.Schema.User.ID).Where("LENGTH(password) < ?", 6)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id FROM users WHERE LENGTH(password) < ?"))
					Expect(args).To(ConsistOf(6))
				})

				It("should generate a where with string pointer conditions with args", func() {
					str := "LENGTH(password) < ?"
					q := tests.NewUserQuery().Select(tests.Schema.User.ID).Where(&str, 6)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id FROM users WHERE LENGTH(password) < ?"))
					Expect(args).To(ConsistOf(6))
				})

				It("should generate a where with Sqlizer conditions", func() {
					q := tests.NewUserQuery().Select(tests.Schema.User.ID).Where(sq.And{
						sq.Eq{"id": 1},
						ceous.OpNot(sq.Eq{"password": "12345"}),
					})
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT id FROM users WHERE (id = ? AND NOT (password = ?))"))
					Expect(args).To(ConsistOf(1, "12345"))
				})
			})
		})
	})

	Context("Count", func() {
		BeforeEach(func() {
			tests.DBStart()
			tests.DBUsersCreate()
			tests.DBUsersInsertJoes()
		})

		AfterEach(func() {
			tests.DBStop()
		})

		It("should count using one condition", func() {
			n, err := tests.NewUserQuery(ceous.WithDB(tests.DB)).ByID(1).Count()
			Expect(err).ToNot(HaveOccurred())
			Expect(n).To(BeEquivalentTo(1))
		})

		It("should count using a where", func() {
			n, err := tests.NewUserQuery(ceous.WithDB(tests.DB)).Where(ceous.Ne(tests.Schema.User.ID, 1)).Count()
			Expect(err).ToNot(HaveOccurred())
			Expect(n).To(BeEquivalentTo(3))
		})

		It("should count not matching anything", func() {
			n, err := tests.NewUserQuery(ceous.WithDB(tests.DB)).ByID(50).Count()
			Expect(err).ToNot(HaveOccurred())
			Expect(n).To(BeEquivalentTo(0))
		})

		It("should not take limit into consideration", func() {
			n, err := tests.NewUserQuery(ceous.WithDB(tests.DB)).Limit(3).Count()
			Expect(err).ToNot(HaveOccurred())
			Expect(n).To(BeEquivalentTo(4))
		})

		It("should count using alias", func() {
			n, err := tests.NewUserQuery(ceous.WithDB(tests.DB), ceous.WithSchema(tests.Schema.User.As("u"))).Count()
			Expect(err).ToNot(HaveOccurred())
			Expect(n).To(BeEquivalentTo(4))
		})
	})

	Context("Select", func() {
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

		It("should retrieve a user", func() {
			user, err := tests.NewUserQuery(ceous.WithDB(tests.DB)).ByID(1).One()
			Expect(err).ToNot(HaveOccurred())
			Expect(user.ID).To(Equal(1))
			Expect(user.Name).To(Equal("Snake Eyes"))
		})

		It("should retrieve a model with relation", func() {
			userGroup, err := tests.NewUserGroupQuery(ceous.WithDB(tests.DB)).ByID(tests.UserGroupPK{
				UserID:  1,
				GroupID: 2,
			}).WithUser().One()
			Expect(err).ToNot(HaveOccurred())
			Expect(userGroup.User.ID).To(Equal(1))
			Expect(userGroup.User.Name).To(Equal("Snake Eyes"))
		})

		It("should retrieve a model with relation", func() {
			userGroup, err := tests.NewUserGroupQuery(ceous.WithDB(tests.DB)).WithUser().ByID(tests.UserGroupPK{
				UserID:  1,
				GroupID: 2,
			}).One()
			Expect(err).ToNot(HaveOccurred())
			Expect(userGroup.User).ToNot(BeNil())
			Expect(userGroup.User.ID).To(Equal(1))
			Expect(userGroup.User.Name).To(Equal("Snake Eyes"))
		})

		It("should retrieve models with relation", func() {
			userGroups, err := tests.NewUserGroupQuery(ceous.WithDB(tests.DB)).WithUser().OrderBy(tests.Schema.UserGroup.ID.UserID, tests.Schema.UserGroup.ID.GroupID).All()
			Expect(err).ToNot(HaveOccurred())
			Expect(userGroups).To(HaveLen(4))
			Expect(userGroups[0].User).ToNot(BeNil())
			Expect(userGroups[0].User.ID).To(Equal(1))
			Expect(userGroups[0].User.Name).To(Equal("Snake Eyes"))
			Expect(userGroups[1].User.ID).To(Equal(1))
			Expect(userGroups[1].User.Name).To(Equal("Snake Eyes"))
			Expect(userGroups[2].User.ID).To(Equal(2))
			Expect(userGroups[2].User.Name).To(Equal("Scarlet"))
			Expect(userGroups[3].User.ID).To(Equal(4))
			Expect(userGroups[3].User.Name).To(Equal("Duke"))
		})
	})
})
