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
			// tests.DBStart()
			// tests.DBUsersCreate()
			// tests.DBUsersInsertJoes()
		})

		AfterEach(func() {
			// tests.DBStop()
		})

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

				It("should select specified fields", func() {
					q := tests.NewUserQuery().Select(tests.Schema.User.Name, tests.Schema.User.CreatedAt)
					builder, err := q.BaseQuery.Builder()
					Expect(err).ToNot(HaveOccurred())
					sql, args, err := builder.ToSql()
					Expect(err).ToNot(HaveOccurred())
					Expect(sql).To(Equal("SELECT name, created_at FROM users"))
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
			})
		})
	})
})
