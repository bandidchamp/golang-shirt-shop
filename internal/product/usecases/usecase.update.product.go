package usecases

import (
	"shirt-shop/internal/models"
	"strconv"
)

func (p *productUC) UpdateProduct(pid string, productF *models.ProductForm, product *models.Product) error {
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		return err
	}

	// Id        int     `json:"id" gorm:"size:64"`
	// Name      string  `json:"name" gorm:"size:4096" validate:"required"`
	// Catagory  int     `json:"catagory" gorm:"size:64" validate:"required"`
	// Size      int     `json:"size" gorm:"size:64" validate:"required"`
	// Gender    int     `json:"gender" gorm:"size:64" validate:"required"`
	// Price     float32 `json:"price" gorm:"size:64" validate:"required"`
	// Quantiry  float32 `json:"quantity" gorm:"size:64" validate:"required"`
	// Ispadding bool    `json:"ispadding" gorm:"size:1"`

	Catagory, Catagory_err := strconv.Atoi(productF.Catagory)
	if Catagory_err != nil {
		return Catagory_err
	}
	Size, Size_err := strconv.Atoi(productF.Size)
	if Size_err != nil {
		return Size_err
	}
	Gender, Gender_err := strconv.Atoi(productF.Gender)
	if Gender_err != nil {
		return Gender_err
	}

	Price, Price_err := strconv.ParseFloat(productF.Price, 32)
	if Price_err != nil {
		return Price_err
	}
	Quantiry, Quantiry_err := strconv.ParseFloat(productF.Quantiry, 32)
	if Quantiry_err != nil {
		return Quantiry_err
	}

	var productModels models.Product
	productModels = models.Product{
		Id:        pidInt,
		Name:      productF.Name,
		Catagory:  Catagory,
		Size:      Size,
		Gender:    Gender,
		Price:     Price,
		Quantiry:  Quantiry,
		Ispadding: false,
	}

	Repoerror := p.productRepo.UpdateProduct(&productModels)
	if Repoerror != nil {
		return Repoerror
	}
	return nil
}
