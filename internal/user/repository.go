package user

import "shirt-shop/internal/models"

// import "shirt-shop/internal/models"

// type RopoInterface interface {
// 	GetProductById(product_id int, product *models.Product) error
// 	GetProductAll(product *[]models.Product) error
// 	InsertProduct(product *models.ProductForm) error
// }
type RopoInterface interface {
	GetUserById(id string, user *models.User) error
	GetUserAll(user *[]models.User) error
	CreateUser(UserData *models.UserForm) error
	UpdateUser(id string, UserData *models.UserForm) error
	DeleteUser(id string, User *models.User) error
	GetRole(role *[]models.Role) error
}
