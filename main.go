package main

import (
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

var config struct {
	DBAddress  string
	DBUsername string
	DBPassword string
	DBName     string
	DBTable    string
	CsvPath    string
}

func initDatabase() {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:3306)/%s",
		config.DBUsername, config.DBPassword, config.DBAddress, config.DBName))
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(fmt.Sprintf(
		"CREATE TABLE if not exists %s (id MEDIUMINT NOT NULL AUTO_INCREMENT, date datetime, city varchar(255), PRIMARY KEY (id) )",
		config.DBTable))
	if err != nil {
		log.Fatal(err)
	}
}

func parseCSV() [][]string {
	file, err := os.Open(config.CsvPath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	entries, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return entries
}

func saveData(entries [][]string) int64 {

	sqlString := "INSERT INTO test(date, city) VALUES "
	data := []interface{}{}

	for _, v := range entries {
		sqlString += "(?, ?),"
		data = append(data, v[0], v[1])
	}

	sqlString = strings.TrimSuffix(sqlString, ",")
	query, err := db.Prepare(sqlString)
	if err != nil {
		log.Fatal(err)
	}

	results, err := query.Exec(data...)
	if err != nil {
		log.Fatal(err)
	}

	newEntries, err := results.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return newEntries
}

func init() {

	flag.StringVar(&config.DBAddress, "db-address", "127.0.0.1", "Database url")
	flag.StringVar(&config.DBUsername, "db-username", "root", "Database user")
	flag.StringVar(&config.DBPassword, "db-password", "", "Database password")
	flag.StringVar(&config.DBName, "db-name", "test", "Database name")
	flag.StringVar(&config.DBTable, "db-table", "", "Table to create")
	flag.StringVar(&config.CsvPath, "file", "", "CSV File")
	flag.Parse()

	if config.CsvPath == "" {
		flag.PrintDefaults()
	}

}

func main() {

	initDatabase()
	defer db.Close()
	entriesCSV := parseCSV()
	results := saveData(entriesCSV)
	fmt.Printf("%d inserted", results)
}
