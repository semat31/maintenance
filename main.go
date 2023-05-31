package main

import "fmt"

func CreatCode(divide string, parts ...string) (code string, err bool) {
	if len(parts) == 0 {
		code = "none"
		err = false
	} else {
		for _, part := range parts {
			code = code + part + divide
		}
		err = true
	}

	return code, err
}

func main() {
	//CreatCode("29", "VA", "2131")

	fmt.Println(CreatCode("", "29", "VA", "2151"))
}
