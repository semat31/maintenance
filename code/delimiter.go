package code

type Delimiter struct {
	delimiter string
}

func NewDelimiter(d string) *Delimiter {
	return &Delimiter{d}
}
