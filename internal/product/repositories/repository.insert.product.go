package repositories

import (
	"shirt-shop/internal/models"
)

func (db *ProductRepo) InsertProduct(product *models.ProductForm) error {
	result := db.conn.Table(models.ProductTname).Select(
		"Name",
		"Catagory",
		"Size",
		"Gender",
		"Price",
		"Quantiry",
		"Ispadding",
	).Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
