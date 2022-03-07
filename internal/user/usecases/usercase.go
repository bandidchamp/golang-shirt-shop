package usecases

import (
	"errors"
	"regexp"
	"shirt-shop/internal/models"
	"shirt-shop/internal/user"
	"strconv"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type userUC struct {
	userRepo user.RopoInterface
}

func NewUserUseCase(userRepo user.RopoInterface) user.UCInterface {
	return &userUC{userRepo: userRepo}
}

func (u *userUC) GetUserById(id string, user *models.User) error {
	intVar, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if intVar < 0 {
		return errors.New("id must not be less than 0.")
	}
	rerr := u.userRepo.GetUserById(id, user)
	if rerr != nil {
		return err
	}
	return nil
}

func (u *userUC) GetUserAll(user *[]models.User) error {
	rerr := u.userRepo.GetUserAll(user)
	if rerr != nil {
		return rerr
	}
	return nil
}

// insert
func (u *userUC) InsertUser(user *models.UserForm) error {
	err := validator.New().Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return validationErrors
	}

	if match, err := regexp.MatchString("([0-9]+)", user.Name); match || err != nil {
		return errors.New("name must string")
	}

	if match, err := regexp.MatchString("([0-9]+)", user.Surname); match || err != nil {
		return errors.New("surname must string")
	}

	Hash, err := bcrypt.GenerateFromPassword([]byte(user.Hash), 8)
	if err != nil {
		return errors.New("can not genarate hash.")
	}

	user.Hash = string(Hash)

	rerr := u.userRepo.CreateUser(user)
	if rerr != nil {
		return rerr
	}
	return nil
}

// update
func (u *userUC) UpdateUser(id string, user *models.UserForm) error {
	intVar, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if intVar < 0 {
		return errors.New("id must not be less than 0.")
	}
	rerr := u.userRepo.UpdateUser(id, user)
	if rerr != nil {
		return rerr
	}
	return nil
}

// delete
func (u *userUC) DeleteUser(id string, user *models.User) error {
	intVar, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if intVar < 0 {
		return errors.New("id must not be less than 0.")
	}
	rerr := u.userRepo.DeleteUser(id, user)
	if rerr != nil {
		return rerr
	}
	return nil
}

// get role.
func (u *userUC) GetRole(role *[]models.Role) error {
	err := u.userRepo.GetRole(role)
	if err != nil {
		return err
	}
	return nil
}
