package app

import (
	"fmt"
	"os"
)

func ExitWithError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}
