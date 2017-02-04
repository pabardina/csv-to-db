package app_test

import (
	"database/sql"
	"os"

	. "github.com/pabardina/csv-to-db/app"

	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/mattn/go-sqlite3"
)

var _ = Describe("App", func() {

	var dbWriter *DatabaseWriter
	var err error
	var db *sql.DB
	var csvPath string

	Context("When the application is launched", func() {

		BeforeSuite(func() {
			db, _ = sql.Open("sqlite3", "./test.db")
			csvPath = "../static/data.csv"
			dbWriter = &DatabaseWriter{
				DB:     db,
				Buffer: 300,
				Table:  "test",
			}
		})

		It("Should init the database without error", func() {
			err = dbWriter.CreateDBTable()
			Expect(err).To(BeNil())
		})

		It("Should successfully parse the CSV file", func() {
			err = ParseCSV(csvPath, dbWriter)
			Expect(err).To(BeNil())

		})

		It("Should have inserted data in the database", func() {
			results, _ := dbWriter.DB.Query(fmt.Sprintf("Select count(*) from %s", dbWriter.Table))
			count := 0
			for results.Next() {
				_ = results.Scan(&count)
			}
			Expect(count).NotTo(Equal(0))

		})

		AfterSuite(func() {
			os.Remove("./test.db")
		})

	})

})
