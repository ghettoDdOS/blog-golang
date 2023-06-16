package models

type Article struct {
	Title             string
	Subject           string
	Keywords          []*Keyword
	Annotation        string
	YearOfPublication int
	SourceLink        string
	baseModel
}
