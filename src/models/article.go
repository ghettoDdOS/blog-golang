package models

type Article struct {
	Title             string
	Subject           string
	Keywords          []string
	Annotation        string
	YearOfPublication int
	SourceLink        string
	baseModel
}
