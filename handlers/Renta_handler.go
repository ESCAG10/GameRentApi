package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type RentaHandler struct {
	RentaService *services.RentaService
}

func NewRentaHandler(service *services.RentaService) *RentaHandler {
	return &RentaHandler{
		RentaService: service,
	}
}

func (h *RentaHandler) CreateRenta(c *fiber.Ctx) error {

	var renta models.Renta

	if err := c.BodyParser(&renta); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	err := h.RentaService.Create(&renta)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Renta registrada correctamente",
	})
}