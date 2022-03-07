package http

import (
	"fmt"
	"shirt-shop/interfaces"
	"shirt-shop/internal/models"
	"shirt-shop/internal/user"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userUC user.UCInterface
}

func NewUserHandler(userUC user.UCInterface) user.Handler {
	return &userHandler{
		userUC: userUC,
	}
}

func (u *userHandler) GetUserById(c *fiber.Ctx) error {
	uid := c.FormValue("id")
	var user models.User
	err := u.userUC.GetUserById(uid, &user)
	if err != nil {
		var payload interfaces.ResponsePayload
		payload = interfaces.ResponsePayload{
			Code:    fiber.StatusBadGateway,
			Status:  false,
			Message: fmt.Sprint(err),
			Payload: nil,
		}
		return c.Status(fiber.StatusBadGateway).JSON(payload)
	}
	var payload interfaces.ResponsePayload
	payload = interfaces.ResponsePayload{
		Status:  true,
		Code:    200,
		Message: "ok",
		Payload: user,
	}
	return c.Status(200).JSON(payload)
}
func (u *userHandler) GetUserAll(c *fiber.Ctx) error {
	var user []models.User
	err := u.userUC.GetUserAll(&user)
	if err != nil {
		var payload interfaces.ResponsePayload
		payload = interfaces.ResponsePayload{
			Code:    fiber.StatusBadGateway,
			Status:  false,
			Message: fmt.Sprint(err),
			Payload: nil,
		}
		return c.Status(fiber.StatusBadGateway).JSON(payload)
	}
	var payload interfaces.ResponsePayload
	payload = interfaces.ResponsePayload{
		Status:  true,
		Code:    200,
		Message: "ok",
		Payload: user,
	}
	return c.Status(200).JSON(payload)
}
func (u *userHandler) InsertUser(c *fiber.Ctx) error {
	var user models.UserForm
	user = models.UserForm{
		Name:     c.FormValue("name"),
		Surname:  c.FormValue("surname"),
		Username: c.FormValue("username"),
		Hash:     c.FormValue("password"),
		Role:     c.FormValue("role"),
	}
	err := u.userUC.InsertUser(&user)
	if err != nil {
		var payload interfaces.ResponsePayload
		payload = interfaces.ResponsePayload{
			Code:    fiber.StatusBadGateway,
			Status:  false,
			Message: fmt.Sprint(err),
			Payload: nil,
		}
		return c.Status(fiber.StatusBadGateway).JSON(payload)
	}
	var payload interfaces.ResponsePayload
	payload = interfaces.ResponsePayload{
		Status:  true,
		Code:    200,
		Message: "ok",
		Payload: nil,
	}
	return c.Status(200).JSON(payload)
}
func (u *userHandler) UpdateUser(c *fiber.Ctx) error {
	return nil
}
func (u *userHandler) DeleteUser(c *fiber.Ctx) error {
	return nil
}

func (u *userHandler) GetRole(c *fiber.Ctx) error {
	var role []models.Role
	err := u.userUC.GetRole(&role)
	if err != nil {
		var payload interfaces.ResponsePayload
		payload = interfaces.ResponsePayload{
			Code:    fiber.StatusBadGateway,
			Status:  false,
			Message: fmt.Sprint(err),
			Payload: nil,
		}
		return c.Status(fiber.StatusBadGateway).JSON(payload)
	}
	var payload interfaces.ResponsePayload
	payload = interfaces.ResponsePayload{
		Status:  true,
		Code:    200,
		Message: "ok",
		Payload: role,
	}
	return c.Status(200).JSON(payload)
}
