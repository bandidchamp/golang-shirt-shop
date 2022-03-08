package product

import (
	"context"
	"shirt-shop/internal/models"
)

type UCInterface interface {
	CheckProductID(product_id string) (*models.Product, error)
	CheckProductALL() (*[]models.Product, error)
	InsertProduct(*models.ProductForm) error
	Size(size *[]models.Product_size) error
	Gender(gender *[]models.Product_gender) error
	Catagory(catagory *[]models.Product_catagory) error
	GetProduct(ctx context.Context, filter *models.ProductFilter, products *[]models.Product) error
}
