package usecases

import (
	"shirt-shop/internal/models"
)

func (p *productUC) InsertProduct(productF *models.ProductForm) error {
	Repoerror := p.productRepo.InsertProduct(productF)
	if Repoerror != nil {
		return Repoerror
	}
	return nil
}
