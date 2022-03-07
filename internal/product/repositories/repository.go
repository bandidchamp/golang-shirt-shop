package repositories

import (
	"shirt-shop/internal/models"
	"shirt-shop/internal/product"

	"gorm.io/gorm"
)

type ProductRepo struct {
	conn *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.RopoInterface {
	return &ProductRepo{
		conn: db,
	}
}
func (db *ProductRepo) GetProductOne(pid int, product *models.Product) error {
	// find one product
	result := db.conn.Table(models.ProductTname).Find(&product, "id = ?", pid)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *ProductRepo) GetProductAll(product *[]models.Product) error {
	// find all product
	result := db.conn.Table(models.ProductTname).Find(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *ProductRepo) InsertProduct(product *models.ProductForm) error {
	result := db.conn.Table(models.ProductTname).Select(
		"Name",
		"Catagory",
		"Size",
		"Gender",
		"Price",
		"Quantiry",
	).Create(&product)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *ProductRepo) Size(size *[]models.Product_size) error {
	result := db.conn.Table(models.SizeTname).Find(&size)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (db *ProductRepo) Gender(gender *[]models.Product_gender) error {
	result := db.conn.Table(models.GenderTname).Find(&gender)
	if result != nil {
		return result.Error
	}
	return nil
}
func (db *ProductRepo) Catagory(catagory *[]models.Product_catagory) error {
	result := db.conn.Table(models.CatagoryTname).Find(&catagory)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
