package mysql

import (
	"fmt"
	"shirt-shop/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabase(config *configs.Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MySQL.User,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.Database,
	)), &gorm.Config{})

	return db, err
}
