package http

import (
	"fmt"
	"shirt-shop/interfaces"
	"shirt-shop/internal/models"

	"github.com/gofiber/fiber/v2"
)

func (ph *productHandler) UpdateProduct(c *fiber.Ctx) error {

	var params models.ProductForm
	params = models.ProductForm{
		Name:      c.FormValue("name"),
		Catagory:  c.FormValue("catagory"),
		Size:      c.FormValue("size"),
		Gender:    c.FormValue("gender"),
		Price:     c.FormValue("price"),
		Quantiry:  c.FormValue("quantity"),
		Ispadding: "0",
	}

	var product models.Product
	var pid string = c.Query("pid")

	err := ph.productUC.UpdateProduct(pid, &params, &product)
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
