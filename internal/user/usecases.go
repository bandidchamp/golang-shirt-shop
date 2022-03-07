package user

import "shirt-shop/internal/models"

type UCInterface interface {
	GetUserById(id string, user *models.User) error
	GetUserAll(user *[]models.User) error
	InsertUser(user *models.UserForm) error
	UpdateUser(id string, user *models.UserForm) error
	DeleteUser(id string, user *models.User) error
	GetRole(role *[]models.Role) error
}
