package http

import (
	"fmt"
	"os"
	"shirt-shop/interfaces"
	"shirt-shop/internal/authen"
	"shirt-shop/internal/models"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type authenHandler struct {
	authen authen.UCInterface
}

func NewAuthenHandler(authenUC authen.UCInterface) authen.Handler {

	return &authenHandler{
		authen: authenUC,
	}
}

func (ah *authenHandler) Login(c *fiber.Ctx) error {

	var authenData *models.AuthForm
	authenData = &models.AuthForm{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}

	token, err := ah.authen.CheckAuthen(authenData)
	if err != nil {
		var payload interfaces.ResponsePayload
		payload = interfaces.ResponsePayload{
			Code:    fiber.StatusBadGateway,
			Status:  true,
			Message: fmt.Sprintf("%s", err),
			Payload: nil,
		}
		return c.Status(fiber.StatusBadGateway).JSON(payload)
	}

	COOKIE_EXPIRES, err := strconv.Atoi(os.Getenv("APP_COOKIE_EXPIRES"))
	if err != nil {
		var payload interfaces.ResponsePayload
		payload = interfaces.ResponsePayload{
			Code:    fiber.StatusBadGateway,
			Status:  true,
			Message: fmt.Sprintf("%s", err),
			Payload: nil,
		}
		return c.Status(fiber.StatusBadGateway).JSON(payload)
	}

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * time.Duration(COOKIE_EXPIRES)),
		HTTPOnly: true,
	})

	var payload interfaces.ResponsePayload
	payload = interfaces.ResponsePayload{
		Status:  true,
		Code:    200,
		Message: "ok",
		Payload: token,
	}
	return c.Status(200).JSON(payload)
}

func (ah *authenHandler) HandleGetUserFromToken(c *fiber.Ctx) error {
	users := c.Locals("user").(*jwt.Token)
	claims := users.Claims.(jwt.MapClaims)

	var payload interfaces.ResponsePayload
	payload = interfaces.ResponsePayload{
		Status:  true,
		Code:    200,
		Message: "ok",
		Payload: claims,
	}
	return c.Status(200).JSON(payload)
}
