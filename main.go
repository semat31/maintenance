package main

import "fmt"

func CreatCode(divide string, parts ...string) (code string, err bool) {

	fmt.Println("divide   ", divide)
	fmt.Println("len(parts)   ", len(parts))

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

	return code, err
}

func main() {
	//CreatCode("29", "VA", "2131")

	fmt.Println(CreatCode("-", "29", "VA", "2151"))
}
