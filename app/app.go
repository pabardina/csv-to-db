package app

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

var InsertedValues int64
var allInsertsDone = sync.WaitGroup{}

type DatabaseWriter struct {
	Buffer int
	DB     *sql.DB
	Table  string
}

type ValueStruct struct {
	Date   string
	Number float64
}

type ValueStructWriter interface {
	Write(values []ValueStruct) error
	GetBufferSQL() int
}

func (d *DatabaseWriter) CreateDBTable() error {
	var err error

	err = d.DB.Ping()
	if err != nil {
		return err
	}

	_, err = d.DB.Exec(fmt.Sprintf(
		"CREATE TABLE if not exists %s (date date, number float)",
		d.Table))
	if err != nil {
		return err
	}
	return nil
}

func (d *DatabaseWriter) GetBufferSQL() int {
	return d.Buffer
}

func (d *DatabaseWriter) Write(values []ValueStruct) error {
	allInsertsDone.Add(1)
	defer allInsertsDone.Add(-1)
	sqlString := fmt.Sprintf("INSERT INTO %s(date, number) VALUES ", d.Table)
	data := []interface{}{}

	for _, v := range values {
		sqlString += "(?, ?),"
		data = append(data, v.Date, v.Number)
	}

	sqlString = strings.TrimSuffix(sqlString, ",")
	query, err := d.DB.Prepare(sqlString)
	if err != nil {
		return err
	}

	results, err := query.Exec(data...)
	if err != nil {
		return err
	}

	newEntries, err := results.RowsAffected()
	if err != nil {
		return err
	}
	InsertedValues += newEntries
	return nil
}

func ParseCSV(csvPath string, v ValueStructWriter) error {
	var values []ValueStruct

	file, err := os.Open(csvPath)

	if err != nil {
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2

	_, err = reader.Read()

	if err != nil {
		return err
	}

	bufferSQL := v.GetBufferSQL()

	for {
		record, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				// Add last records in DB
				if errWrite := v.Write(values); errWrite != nil {
					return err
				}
				break
			}
			return err
		}

		number, err := strconv.ParseFloat(record[1], 64)

		if err != nil {
			return err
		}

		val := ValueStruct{
			Date:   record[0],
			Number: number,
		}
		values = append(values, val)

		if len(values) > bufferSQL {
			go func(values []ValueStruct) {
				if err := v.Write(values); err != nil {
					ExitWithError(err)
				}
			}(values)
			values = []ValueStruct{}
		}
	}

	if err != nil {
		return err
	}
	return nil
}
