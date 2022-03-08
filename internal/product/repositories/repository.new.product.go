package repositories

import (
	"shirt-shop/internal/product"

	"github.com/go-redis/redis/v8"
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
