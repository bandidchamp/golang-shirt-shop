package http

import (
	"fmt"
	"shirt-shop/interfaces"
	"shirt-shop/internal/models"
	"shirt-shop/internal/product"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productUC product.UCInterface
}

func NewProductHandler(productUC product.UCInterface) product.Handler {

	return &productHandler{
		productUC: productUC,
	}
}

func (ph *productHandler) GetProductOne(c *fiber.Ctx) error {
	pid := c.FormValue("id")
	var product *models.Product
	product, err := ph.productUC.CheckProductID(pid)
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
	return c.Status(200).JSON(fiber.Map{"product": product})
}

func (ph *productHandler) GetProductAll(c *fiber.Ctx) error {
	product, err := ph.productUC.CheckProductALL()
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
	return c.Status(200).JSON(fiber.Map{"productAll": product})
}

func (ph *productHandler) InsertProduct(c *fiber.Ctx) error {

	var params models.ProductForm
	params = models.ProductForm{
		Name:     c.FormValue("name"),
		Catagory: c.FormValue("catagory"),
		Size:     c.FormValue("size"),
		Gender:   c.FormValue("gender"),
		Price:    c.FormValue("price"),
		Quantiry: c.FormValue("quantity"),
	}

	err := ph.productUC.InsertProduct(&params)
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

func (ph *productHandler) Size(c *fiber.Ctx) error {
	var size []models.Product_size
	err := ph.productUC.Size(&size)
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
		Payload: size,
	}
	return c.Status(200).JSON(payload)
}
func (ph *productHandler) Gender(c *fiber.Ctx) error {
	var gender []models.Product_gender
	err := ph.productUC.Gender(&gender)
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
		Payload: gender,
	}
	return c.Status(200).JSON(payload)
}
func (ph *productHandler) Catagory(c *fiber.Ctx) error {
	var catagory []models.Product_catagory
	err := ph.productUC.Catagory(&catagory)
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
		Payload: catagory,
	}
	return c.Status(200).JSON(payload)
}
