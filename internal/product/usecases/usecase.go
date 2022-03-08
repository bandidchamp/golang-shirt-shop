package usecases

import (
	"context"
	"errors"
	"fmt"
	"shirt-shop/internal/models"
	"shirt-shop/internal/product"
	"strconv"
)

type productUC struct {
	productRepo product.RopoInterface
}

func NewProductUseCase(productRepo product.RopoInterface) product.UCInterface {
	return &productUC{productRepo: productRepo}
}

func (p *productUC) CheckProductID(ctxId string) (*models.Product, error) {
	id, err := strconv.Atoi(ctxId)
	// logic check id
	if id == 0 && id < 0 {
		return nil, errors.New("id have error")
	}
	if err != nil {
		return nil, err
	}
	var product models.Product
	Repoerror := p.productRepo.GetProductById(id, &product)
	if Repoerror != nil {
		return nil, Repoerror
	}
	return &product, nil
}

func (p *productUC) CheckProductALL() (*[]models.Product, error) {
	var product []models.Product
	Repoerror := p.productRepo.GetProductAll(&product)
	if Repoerror != nil {
		return nil, Repoerror
	}
	return &product, nil
}

func (p *productUC) InsertProduct(productF *models.ProductForm) error {
	Repoerror := p.productRepo.InsertProduct(productF)
	if Repoerror != nil {
		return Repoerror
	}
	return nil
}

func (p *productUC) Size(size *[]models.Product_size) error {
	Repoerror := p.productRepo.Size(size)
	if Repoerror != nil {
		return Repoerror
	}
	return nil
}
func (p *productUC) Gender(gender *[]models.Product_gender) error {
	Repoerror := p.productRepo.Gender(gender)
	if Repoerror != nil {
		return Repoerror
	}
	return nil
}
func (p *productUC) Catagory(catagory *[]models.Product_catagory) error {
	Repoerror := p.productRepo.Catagory(catagory)
	if Repoerror != nil {
		return Repoerror
	}
	return nil
}

func (p *productUC) GetProduct(ctx context.Context, filter *models.ProductFilter, products *[]models.Product) error {
	// logic switch db or cache.
	key := fmt.Sprintf("product_key_%s_%s_%s_%s_%s_%s_%s", filter.Name, filter.Gender, filter.Catagory, filter.Size, filter.Limit, filter.Offset, filter.OrderBy)
	// check cache
	err := p.productRepo.GetCache(ctx, key, products)

	// if cache have error.
	if err != nil && errors.Unwrap(err).Error() != "redis: nil" {
		return fmt.Errorf("failed to get cache product with error: %w", err)
	}

	// notfound product cache = "query db" Then "save cache".
	if err != nil && errors.Unwrap(err).Error() == "redis: nil" {
		fmt.Println("Get data from db.")
		// query filter and save cache.
		rperr := p.productRepo.GetProductyBy(filter, products)
		if rperr != nil {
			return rperr
		}
		// set product to cache.
		SCerr := p.productRepo.SetCache(ctx, key, products)
		if SCerr != nil {
			return SCerr
		}
	} else {
		fmt.Println("Get data from Cache")
	}

	return nil
}
