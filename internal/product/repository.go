package product

import (
	"context"
	"shirt-shop/internal/models"
)

type RopoInterface interface {
	GetProductById(product_id int, product *models.Product) error
	GetProductAll(product *[]models.Product) error
	GetProductyBy(filter *models.ProductFilter, products *[]models.Product) error

	Size(size *[]models.Product_size) error
	Gender(gender *[]models.Product_gender) error
	Catagory(catagory *[]models.Product_catagory) error

	GetCache(ctx context.Context, key string, products *[]models.Product) error
	SetCache(ctx context.Context, key string, value *[]models.Product) error

	InsertProduct(product *models.ProductForm) error
	UpdateProduct(Product *models.Product) error
	PaddingProduct(Product *models.Product) error
}
