package http

import (
	"fmt"
	"shirt-shop/interfaces"
	"shirt-shop/internal/models"

	"github.com/gofiber/fiber/v2"
)

func (ph *productHandler) PaddingProduct(c *fiber.Ctx) error {

	var product models.Product
	var pid string = c.Query("pid")

	err := ph.productUC.PaddingProduct(pid, true, &product)
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
	return c.Status(200).JSON(fiber.Map{"result": "ok"})
}
