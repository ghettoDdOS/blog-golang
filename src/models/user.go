package models

import "blog/src/config"

type User struct {
	BaseModel
	FirstName  string `form:"FirstName"`
	LastName   string `form:"LastName"`
	Patronymic string `form:"Patronymic"`
	Job        string `form:"Job"`
	Email      string `form:"Email"`
	Password   string `form:"Password"`
}

func setUniqueEmailConstraint() {
	query := `CREATE CONSTRAINT unique_email
	FOR (u:User)
	REQUIRE u.email IS UNIQUE
	`
	db := config.GetDatabaseConnection()
	conn := db.NewExecuter()
	conn.DoQuery(query)
	conn.Close()
}
