package usecases

import (
	"shirt-shop/internal/product"
)

type productUC struct {
	productRepo product.RopoInterface
}

func NewProductUseCase(productRepo product.RopoInterface) product.UCInterface {
	return &productUC{productRepo: productRepo}
}
