package usecases

import (
	"shirt-shop/internal/models"
	"strconv"
)

func (p *productUC) PaddingProduct(pid string, padding bool, product *models.Product) error {
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

	var productModels models.Product
	productModels = models.Product{
		Id:        pidInt,
		Ispadding: padding,
	}

	Repoerror := p.productRepo.UpdateProduct(&productModels)
	if Repoerror != nil {
		return Repoerror
	}
	return nil
}
