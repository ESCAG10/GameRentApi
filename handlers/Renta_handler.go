package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type RentaHandler struct {
	Service *services.RentaService
}

func NewRentaHandler(service *services.RentaService) *RentaHandler {
	return &RentaHandler{
		Service: service,
	}
}

func (h *RentaHandler) CreateRenta(c *fiber.Ctx) error {

	var renta models.Renta

	if err := c.BodyParser(&renta); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "datos inválidos",
		})
	}

	if err := h.Service.Create(&renta); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(renta)
}

func (h *RentaHandler) GetRentas(c *fiber.Ctx) error {

	rentas, err := h.Service.FindAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rentas)
}

func (h *RentaHandler) GetRenta(c *fiber.Ctx) error {

	id := c.Params("id")

	renta, err := h.Service.FindByID(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "renta no encontrada",
		})
	}

	return c.JSON(renta)
}

func (h *RentaHandler) UpdateRenta(c *fiber.Ctx) error {

	id := c.Params("id")

	var renta models.Renta

	if err := c.BodyParser(&renta); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "datos inválidos",
		})
	}

	if err := h.Service.Update(id, &renta); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "renta actualizada",
	})
}

func (h *RentaHandler) DeleteRenta(c *fiber.Ctx) error {

	id := c.Params("id")

	if err := h.Service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "renta eliminada",
	})
}