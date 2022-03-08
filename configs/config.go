package configs

import (
	"github.com/joho/godotenv"
)

type Config struct {
	Fiber Fiber
	MySQL MySQL
	Jwt   Jwt
	Redis Redis
}

type Fiber struct {
	Host    string
	Port    string
	TimeOut string
}

type MySQL struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

type Jwt struct {
	Secret  string
	Expires string
}

type Redis struct {
	Host     string
	Port     string
	Password string
	DBNumber string
}

func ReadEnvironment() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}
