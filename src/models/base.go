package models

type BaseModel struct {
	Id int64 `form:"Id"`
}

func InitConstraints() {
	setUniqueEmailConstraint()
}
