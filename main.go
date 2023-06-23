package main

import (
	"fmt"

	"github.com/semat31/maintenance/code"
)

func splitter(d string, c code.Code) []string {
	return []string{d, string(c), "c"}
}

func main() {
	d := code.NewDelimiter("-")

	code1, _ := code.New(d, "29", "VA", "2131")
	fmt.Println(code1)

	fmt.Println(code1.Splitter(d, splitter))
}
