package models

type Article struct {
	BaseModel
	Title             string
	Subject           string
	Keywords          []string
	Annotation        string
	YearOfPublication int
	SourceLink        string
}
