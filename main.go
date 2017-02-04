package main

import (
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pabardina/csv-to-db/app"
)

func init() {

	flag.StringVar(&config.dbAddress, "db-address", "127.0.0.1", "Database url")
	flag.StringVar(&config.dbUsername, "db-username", "root", "Database user")
	flag.StringVar(&config.dbPassword, "db-password", "", "Database password")
	flag.StringVar(&config.dbName, "db-name", "test", "Database name")
	flag.StringVar(&config.dbTable, "db-table", "", "Table to create")
	flag.StringVar(&config.csvPath, "file", "", "CSV File")
	flag.IntVar(&config.bufferSQL, "buffer-sql", 1000, "Values to insert in DB")
	flag.Parse()

	if config.csvPath == "" {
		flag.PrintDefaults()
	}
}

func main() {

	var err error

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:3306)/%s",
		config.dbUsername, config.dbPassword, config.dbAddress, config.dbName))

	if err != nil {
		app.ExitWithError(err)
	}
	db.SetMaxIdleConns(0)
	defer db.Close()

	dbWriter := &app.DatabaseWriter{
		Buffer: config.bufferSQL,
		DB:     db,
		Table:  config.dbTable,
	}

	if err := dbWriter.CreateDBTable(); err != nil {
		app.ExitWithError(err)
	}

	if err := app.ParseCSV(config.csvPath, dbWriter); err != nil {
		app.ExitWithError(err)
	}

	fmt.Printf("%d inserted", app.InsertedValues)
}

var config struct {
	dbAddress  string
	dbUsername string
	dbPassword string
	dbName     string
	dbTable    string
	csvPath    string
	bufferSQL  int
}
