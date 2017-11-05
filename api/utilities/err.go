package utilities

import "fmt"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
