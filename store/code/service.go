package code

import (
	"unicode"
)

// **
func IsCode(code string) (isCode bool) {

	return
}

// **
func SplitCode(code string) ([]string, bool) {

	var index1 int
	var index2 int

	codeSlice := make([]string, 4)

	// ***
	for i, runeI := range code {
		if unicode.IsLetter(runeI) {
			index1 = i
			break
		}

	}
	// ***
	for i, runeI := range code {
		if unicode.IsDigit(runeI) && i > index1 {

			index2 = i
			break
		}

	}

	codeSlice[0] = code
	codeSlice[1] = code[:index1]
	codeSlice[2] = code[index1:index2]
	codeSlice[3] = code[index2:]

	return codeSlice, true
}

// **
