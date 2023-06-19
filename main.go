package main

import (
	"fmt"

	"github.com/semat31/maintenance/code"
)

func main() {
	d := code.NewDelimiter("")

	code1, _ := code.New(d, "29", "VA", "2131")
	fmt.Println(code1)

	code2, _ := code.New(d, "29", "VA", "2131")
	fmt.Println(code2)
}

// func (c code.Code) Splitter() []string {
// 	return []string{}
// }
