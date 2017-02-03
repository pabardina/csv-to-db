package main_test

import (
	"database/sql"
	"encoding/csv"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	_ "github.com/mattn/go-sqlite3"
)

var _ = Describe("Main", func() {

	Describe("initDatabase", func() {

		var err error
		var db *sql.DB

		BeforeSuite(func() {
			db, _ = sql.Open("sqlite3", "./test.db")
		})

		AfterEach(func() {
			err = nil
		})

		Context("when everything is ok", func() {

			db, _ = sql.Open("sqlite3", "./test.db")

			It("should ping the DB", func() {
				err = db.Ping()
				Expect(err).To(BeNil())
			})

			It("should have the table in DB", func() {
			})

		})

		AfterSuite(func() {
			os.Remove("./test.db")
			db = nil
		})

		Describe("parseCSV", func() {

			var err error

			AfterEach(func() {
				err = nil
			})

			Context("when everything is ok", func() {

				var file *os.File

				It("should be possible to open the csv file", func() {
					file, err = os.Open("static/data.csv")
					Expect(file).NotTo(BeNil())
					Expect(err).To(BeNil())
				})

				It("should have data", func() {
					reader := csv.NewReader(file)
					entries, err := reader.ReadAll()

					Expect(len(entries)).To(Equal(1000))
					Expect(err).To(BeNil())
				})

			})

		})

		Describe("parseCSV", func() {

			var err error

			AfterEach(func() {
				err = nil
			})

			Context("when everything is ok", func() {

				var file *os.File

				It("should be possible to open the csv file", func() {
					file, err = os.Open("static/data.csv")
					Expect(file).NotTo(BeNil())
					Expect(err).To(BeNil())
				})

				It("should have data", func() {
					reader := csv.NewReader(file)
					entries, err := reader.ReadAll()

					Expect(len(entries)).To(Equal(1000))
					Expect(err).To(BeNil())
				})

			})

		})
	})
})
