package repositories

import (
	"shirt-shop/internal/models"
)

func (db *ProductRepo) UpdateProduct(Product *models.Product) error {
	result := db.conn.Table(models.ProductTname).Updates(&Product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
