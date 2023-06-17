package code

type Delimiter string

func NewDelimiter(d string) Delimiter {
	return Delimiter(d)
}
