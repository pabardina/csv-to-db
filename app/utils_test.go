package app

import "testing"

func TestGetRecordReturnValStruct(t *testing.T) {
	t.Log("Should return a valStruct")
	data := []string{"2015-01-01", "212.23"}
	valStruct := GetRecordReturnValStruct(data)

	if valStruct.Date == "" {
		t.Error("Excepted to have a date in struct")
	}

	if valStruct.Number == 0 {
		t.Error("Excepted to have a number in struct")
	}
}
