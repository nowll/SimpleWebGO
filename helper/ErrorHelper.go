package helper

import "fmt"

func ErrorHelper(err error)  {
	if err != nil {
	fmt.Println(err)
	return
	}
}