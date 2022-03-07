package repositories

import (
	"shirt-shop/internal/authen"
	"shirt-shop/internal/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	conn *gorm.DB
}

func NewAuthRepository(db *gorm.DB) authen.RopoInterface {
	return &UserRepo{
		conn: db,
	}
}

func (db *UserRepo) BasicAuthen(username string, user *models.User) error {
	result := db.conn.Table("user").Find(&user, "username = ?", username)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
