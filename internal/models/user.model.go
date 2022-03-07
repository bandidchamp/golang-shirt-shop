package models

var UserTname string = "user"
var RoleTname string = "role"

type User struct {
	Id       int    `json:"id" gorm:"size:64;primaryKey;autoIncrement:true;"`
	Name     string `json:"name" gorm:"size:128"`
	Surname  string `json:"surname" gorm:"size:128"`
	Username string `json:"username" gorm:"size:128" validate:"required"`
	Hash     string `json:"hash" gorm:"size:2048" validate:"required"`
	Role     int    `json:"role" gorm:"size:64" validate:"required"`
}

type Role struct {
	Id   int    `json:"id" gorm:"primaryKey;autoIncrement:true;"`
	Name string `json:"name"`
}

type AuthForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserForm struct {
	Name     string `json:"name" gorm:"size:128" validate:"required"`
	Surname  string `json:"surname" gorm:"size:128" validate:"required"`
	Username string `json:"username" gorm:"size:128" validate:"required"`
	Hash     string `json:"hash" gorm:"size:2048" validate:"required"`
	Role     string `json:"role" gorm:"size:64" validate:"required"`
}

func (user *User) TableName() string {
	return "user"
}
