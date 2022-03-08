package cache

import (
	"fmt"
	"shirt-shop/configs"
	"strconv"

	"github.com/go-redis/redis/v8"
)

//* NewRadisCache func for connect to Redis server.
func NewRadisCache(c *configs.Config) (*redis.Client, error) {
	//* Define Redis database number.
	dbNumber, _ := strconv.Atoi(c.Redis.DBNumber)

	redisConnURL := fmt.Sprintf(
		"%s:%s",
		c.Redis.Host,
		c.Redis.Port,
	)

	//* Set Redis options.
	options := &redis.Options{
		Addr:     redisConnURL,
		Password: c.Redis.Password,
		DB:       dbNumber,
	}

	return redis.NewClient(options), nil
	// return nil, nil
}
