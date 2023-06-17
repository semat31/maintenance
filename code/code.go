// ***
package code

import "fmt"

// ***
type Code string

// ***
func New(divide Delimiter, parts ...string) (Code, bool) {

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
				code = code + string(divide)
			}
		}
		err = true
	}

	return Code(code), err
}
