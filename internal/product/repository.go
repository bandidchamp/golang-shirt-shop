package product

import "shirt-shop/internal/models"

type RopoInterface interface {
	GetProductOne(product_id int, product *models.Product) error
	GetProductAll(product *[]models.Product) error
	InsertProduct(product *models.ProductForm) error
	Size(size *[]models.Product_size) error
	Gender(gender *[]models.Product_gender) error
	Catagory(catagory *[]models.Product_catagory) error
}
