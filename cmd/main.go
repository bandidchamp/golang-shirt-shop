package main

import (
	"fmt"
	"os"
	"shirt-shop/configs"
	"shirt-shop/internal/models"
	"shirt-shop/internal/server"
	"shirt-shop/pkg/database/mysql"
)

func main() {
	fmt.Println("Start . . .")

	// Read Evironment.
	err := configs.ReadEnvironment()
	if err != nil {
		fmt.Println("ReadEnvironment:", err)
	}

	// Defind config
	config := &configs.Config{}
	config.MySQL.Host = os.Getenv("MYSQL_HOST")
	config.MySQL.Database = os.Getenv("MYSQL_DATABASE_NAME")
	config.MySQL.User = os.Getenv("MYSQL_USERNAME")
	config.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
	config.MySQL.Port = os.Getenv("MYSQL_PORT")
	config.Fiber.Host = os.Getenv("APP_HOST")
	config.Fiber.Port = os.Getenv("APP_PORT")
	config.Fiber.TimeOut = os.Getenv("APP_SERVER_TIMEOUT")
	config.Jwt.Secret = os.Getenv("APP_JWT_SECRET")
	config.Jwt.Expires = os.Getenv("APP_JWT_EXPIRES")

	// Defind Database
	db, err := mysql.SetupDatabase(config)
	if err != nil {
		fmt.Println("Cann't connect Database")
	}

	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Role{})

	Server := server.Setup(db, config)
	serr := Server.StartServer()
	fmt.Println("started server")
	fmt.Println(serr)
	if serr != nil {
		fmt.Println("cann't start server.")
	}

}
