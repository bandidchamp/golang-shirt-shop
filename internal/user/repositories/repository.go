package repositories

import (
	"errors"
	"shirt-shop/internal/models"
	"shirt-shop/internal/user"

	"gorm.io/gorm"
)

type UserRepo struct {
	conn *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.RopoInterface {
	return &UserRepo{
		conn: db,
	}
}

func (db *UserRepo) GetUserById(id string, user *models.User) error {
	result := db.conn.Table(models.UserTname).Find(&user, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (db *UserRepo) GetUserAll(user *[]models.User) error {
	result := db.conn.Table(models.UserTname).Find(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *UserRepo) CreateUser(UserData *models.UserForm) error {
	result := db.conn.Table(models.UserTname).Select(
		"name",
		"surname",
		"username",
		"hash",
		"role",
	).Create(&UserData)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("cann't insert user.")
	}
	return nil
}
func (db *UserRepo) UpdateUser(id string, UserData *models.UserForm) error {

	result := db.conn.Table(models.UserTname).Where("id = ?", id).Updates(map[string]interface{}{
		"name":    UserData.Name,
		"surname": UserData.Surname,
		"role":    UserData.Role,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (db *UserRepo) DeleteUser(id string, User *models.User) error {
	result := db.conn.Table(models.UserTname).Where("id = ?", id).Delete(&User)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *UserRepo) GetRole(role *[]models.Role) error {
	result := db.conn.Table(models.RoleTname).Find(&role)
	if result != nil {
		return result.Error
	}
	return nil
}
