package usecases

import (
	"shirt-shop/internal/models"
)

func (p *productUC) UpdateProduct(pid string, productF *models.ProductForm, product *models.Product) error {
	Repoerror := p.productRepo.UpdateProduct(pid, productF, product)
	if Repoerror != nil {
		return Repoerror
	}
	return nil
}
