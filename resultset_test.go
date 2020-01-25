package ceous_test

import (
	"github.com/jamillosantos/go-ceous/tests/db"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("ResultSet", func() {
	Describe("RecordResultSet", func() {
		BeforeEach(func() {
			db.DBStart()
			db.DBUsersCreate()
			db.DBUsersInsertJoes()
		})

		AfterEach(func() {
			db.DBStop()
		})
	})
})
