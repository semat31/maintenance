package main

import "fmt"

type Code string

func NewCode(divide string, parts ...string) (Code, bool) {

	fmt.Println("divide   ", divide)
	fmt.Println("parts   ", parts)
	fmt.Println("len(parts)   ", len(parts))

	var code string
	var err bool

	if len(parts) == 0 {
		code = "none"
		err = false
	} else {
		for i, part := range parts {
			code = code + part
			if i < len(parts)-1 {
				code = code + divide
			}
		}
		err = true
	}

	return Code(code), err
}

func main() {
	//
	code1, _ := NewCode("/", "29", "VA", "2131")
	fmt.Println(code1)

	code2 := "29VA2132"
	fmt.Println(code2)

}
