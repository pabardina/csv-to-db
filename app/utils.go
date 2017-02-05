package app

import (
	"fmt"
	"os"
	"strconv"
)

func GetRecordReturnValStruct(record []string) ValueStruct {

	number, err := strconv.ParseFloat(record[1], 64)

	if err != nil {
		ExitWithError(err)
	}

	value := ValueStruct{
		Date:   record[0],
		Number: number,
	}

	return value
}

func ExitWithError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
