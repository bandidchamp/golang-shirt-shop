package configs

import (
	"github.com/joho/godotenv"
)

type Config struct {
	Fiber Fiber
	MySQL MySQL
	Jwt   Jwt
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

func ReadEnvironment() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	return nil
}
