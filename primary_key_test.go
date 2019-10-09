package ceous_test

import (
	"github.com/jamillosantos/go-ceous"
	"github.com/jamillosantos/go-ceous/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrimaryKey", func() {
	Context("Wrapper", func() {
		It("should retrieve a field with prefix", func() {
			pk := &tests.CompositePK{}
			wpk := ceous.WrapPK("commenter_", pk)
			columnAddr, err := wpk.ColumnAddress("commenter_user_id")
			Expect(err).ToNot(HaveOccurred())
			Expect(columnAddr).To(Equal(&pk.UserID))
		})

		It("should fail retrieving a non existing field with prefix", func() {
			pk := &tests.CompositePK{}
			wpk := ceous.WrapPK("commenter_", pk)
			_, err := wpk.ColumnAddress("commenter_comment_id")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ceous.ErrFieldNotFound))
		})

		It("should fail retrieving a fix with wrong prefix", func() {
			pk := &tests.CompositePK{}
			wpk := ceous.WrapPK("commenter_", pk)
			_, err := wpk.ColumnAddress("prefix_comment_id")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ceous.ErrFieldNotFound))
		})

		It("should retrieve a field value with prefix", func() {
			pk := &tests.CompositePK{
				UserID: 1,
				PostID: 2,
			}
			wpk := ceous.WrapPK("commenter_", pk)
			columnValue, err := wpk.Value("commenter_user_id")
			Expect(err).ToNot(HaveOccurred())
			Expect(columnValue).To(Equal(1))
		})

		It("should fail retrieving the value of a non existing field with prefix", func() {
			pk := &tests.CompositePK{
				UserID: 1,
				PostID: 2,
			}
			wpk := ceous.WrapPK("commenter_", pk)
			_, err := wpk.ColumnAddress("commenter_comment_id")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ceous.ErrFieldNotFound))
		})

		It("should fail retrieving the value a fix with wrong prefix", func() {
			pk := &tests.CompositePK{
				UserID: 1,
				PostID: 2,
			}
			wpk := ceous.WrapPK("commenter_", pk)
			_, err := wpk.ColumnAddress("prefix_comment_id")
			Expect(err).To(HaveOccurred())
			Expect(err).To(Equal(ceous.ErrFieldNotFound))
		})
	})
})
