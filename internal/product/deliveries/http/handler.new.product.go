package http

import (
	"shirt-shop/internal/product"
)

type productHandler struct {
	productUC product.UCInterface
}

func NewProductHandler(productUC product.UCInterface) product.Handler {

	return &productHandler{
		productUC: productUC,
	}
}
