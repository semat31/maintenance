package main

import (
	"fmt"

	"github.com/semat31/maintenance/code"
)

func main() {

	code1 := "2V1"
	fmt.Println(code1)

	isCode := code.IsCode(code1)
	fmt.Println(isCode)

	sliceCode, err := code.SplitCode(code1)
	fmt.Println(sliceCode, err)

}
