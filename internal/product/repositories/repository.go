package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"shirt-shop/internal/models"
	"shirt-shop/internal/product"
	"strings"

	"github.com/go-redis/redis/v8"
	"go.elastic.co/apm"
	"gorm.io/gorm"
)

type ProductRepo struct {
	conn  *gorm.DB
	cache *redis.Client
}

func NewProductRepository(db *gorm.DB, cache *redis.Client) product.RopoInterface {
	return &ProductRepo{
		conn:  db,
		cache: cache,
	}
}
func (db *ProductRepo) GetProductById(pid int, product *models.Product) error {
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

func (db *ProductRepo) GetProductyBy(filter *models.ProductFilter, products *[]models.Product) error {
	var sliceCon []string

	var query string

	if filter.Name != "" {
		sliceCon = append(sliceCon, `pd.name LIKE "%`+filter.Name+`%"`)
	}

	if filter.Gender != "" {
		sliceCon = append(sliceCon, `pd.gender = `+filter.Gender+``)
	}

	if filter.Catagory != "" {
		sliceCon = append(sliceCon, `pd.catagory = `+filter.Catagory+``)
	}

	if filter.Size != "" {
		sliceCon = append(sliceCon, `pd.name = `+filter.Size+``)
	}

	// sliceCon concat ADN logic here.
	condition := strings.Join(sliceCon, " AND ")

	if filter.OrderBy != "" {
		query = query + ` ORDER BY pd.id ` + filter.OrderBy + ` `
	}

	if filter.Limit != "" && filter.Offset != "" {
		query = query + ` LIMIT ` + filter.Limit + ` OFFSET ` + filter.Offset + ` `
	}
	result := db.conn.Table(models.ProductTname).Raw(`
		SELECT 
		* 
		FROM product AS pd
		LEFT JOIN product_catagory AS pdc
			ON pdc.product_catagory_id = pd.catagory
		LEFT JOIN product_gender AS pdg
			ON pdg.product_gender_id = pd.gender
		LEFT JOIN product_size AS pds
			ON pds.product_size_id = pd.size
		WHERE ` + condition + ` ` + query + `
	`).Scan(&products)
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

func (db *ProductRepo) GetCache(ctx context.Context, key string, products *[]models.Product) error {

	span, ctx := apm.StartSpan(ctx, "ProductRepo.GetCache", "repo")
	defer span.End()

	val, err := db.cache.Get(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to get cache with error: %w", err)
	}

	if err := json.Unmarshal([]byte(val), &products); err != nil {
		return fmt.Errorf("failed to unmarshal with error: %w", err)
	}

	return nil

}

func (db *ProductRepo) SetCache(ctx context.Context, key string, value *[]models.Product) error {
	span, ctx := apm.StartSpan(ctx, "ProductRepo.GetCache", "repo")
	defer span.End()

	v, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal with error: %w", err)
	}

	if err := db.cache.Set(ctx, key, v, 0).Err(); err != nil {
		return fmt.Errorf("failed to set cache with error: %w", err)
	}

	return nil
}
