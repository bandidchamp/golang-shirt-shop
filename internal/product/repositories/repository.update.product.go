package repositories

import (
	"shirt-shop/internal/models"
)

func (db *ProductRepo) UpdateProduct(pid string, ProductForm *models.ProductForm, Product *models.Product) error {
	result := db.conn.Table(models.ProductTname).Find(&Product, "id = ?", pid).Updates(&ProductForm)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
