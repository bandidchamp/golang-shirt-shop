package usecases

import (
	"errors"
	"os"
	"shirt-shop/internal/authen"
	"shirt-shop/internal/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type authenUC struct {
	authenRepo authen.RopoInterface
}

func NewAuthUseCase(authRepo authen.RopoInterface) authen.UCInterface {
	return &authenUC{authenRepo: authRepo}
}

func (p *authenUC) CheckAuthen(auth *models.AuthForm) (string, error) {
	if auth.Username == "" || auth.Password == "" {
		return "", errors.New("require username and password")
	}
	var user models.User
	Repoerror := p.authenRepo.BasicAuthen(auth.Username, &user)
	if Repoerror != nil {
		return "", Repoerror
	}

	// check login
	Bcrypt_err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(auth.Password))
	if Bcrypt_err != nil {
		// unsuccess.
		return "", Bcrypt_err
	}

	token := jwt.New(jwt.SigningMethodHS256)

	JWT_EXPIRES, err := strconv.Atoi(os.Getenv("APP_JWT_EXPIRES"))
	if err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["surname"] = user.Surname
	claims["username"] = user.Username
	claims["admin"] = user.Role == 1
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(JWT_EXPIRES)).Unix()

	tk, err := token.SignedString([]byte(os.Getenv("APP_JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tk, nil
}
