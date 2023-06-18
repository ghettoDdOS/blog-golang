package models

type Article struct {
	BaseModel
	Title             string   `form:"Title"`
	Subject           string   `form:"Subject"`
	Keywords          []string `form:"Keywords[]"`
	Annotation        string   `form:"Annotation"`
	YearOfPublication int64    `form:"YearOfPublication"`
	SourceLink        string   `form:"SourceLink"`
}
