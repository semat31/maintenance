package code

// **
type SplitFunc func(d string, c Code) []string

// **
func (c Code) Splitter(delimiter *Delimiter, splitter SplitFunc) []string {
	return splitter(delimiter.delimiter, c)
}