package authen

import "shirt-shop/internal/models"

type RopoInterface interface {
	BasicAuthen(username string, user *models.User) error
}
