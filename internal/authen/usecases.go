package authen

import "shirt-shop/internal/models"

type UCInterface interface {
	CheckAuthen(auth *models.AuthForm) (string, error)
}
