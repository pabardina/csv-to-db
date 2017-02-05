package app

import (
	"database/sql"
	"fmt"
	"testing"

	"os"

	_ "github.com/mattn/go-sqlite3"
)

func TestAllTheWorkflow(t *testing.T) {
	t.Log("Test the application (expected values in db)")
	db, _ := sql.Open("sqlite3", "./test.db")
	defer db.Close()

	csvPath := "../static/data.csv"

	dbWriter := &DatabaseWriter{
		DB:     db,
		Buffer: 200,
		Table:  "test",
	}

	if err := dbWriter.CreateDBTable(); err != nil {
		t.Errorf("Excepted no error, but : %s", err)
	}

	if err := ParseCSV(csvPath, dbWriter); err != nil {
		t.Errorf("Excepted no error, but : %s", err)
	}

	results, _ := dbWriter.DB.Query(fmt.Sprintf("Select count(*) from %s", dbWriter.Table))
	count := 0
	for results.Next() {
		_ = results.Scan(&count)
	}

	fmt.Print(count)

	if count == 0 {
		t.Error("Excepted to have values in database, but it's empty")
	}

	os.Remove("./test.db")

}

func TestGetRecordReturnValStruct(t *testing.T) {
	data := []string{"2015-01-01", "212.23"}
	valStruct := GetRecordReturnValStruct(data)

	if valStruct.Date == "" {
		t.Error("Excepted to have a date in struct")
	}

	if valStruct.Number == 0 {
		t.Error("Excepted to have a number in struct")
	}
}
