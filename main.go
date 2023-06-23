package main

import (
	"fmt"

	"github.com/semat31/maintenance/code"
)

func main() {
	d := code.NewDelimiter("-")

	var sp code.SplitFunc

	code1, _ := code.New(d, "29", "VA", "2131")
	fmt.Println(code1)

	fmt.Println(code1.Splitter(*d,sp))
}

