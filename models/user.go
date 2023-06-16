package models

type User struct {
	FirstName  string `form:"FirstName"`
	LastName   string `form:"LastName"`
	Patronymic string `form:"Patronymic"`
	Job        string `form:"Job"`
	Email      string `form:"Email"`
	Password   string `form:"Password"`
	baseModel
}
