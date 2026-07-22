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

// Crear renta
func (h *RentaHandler) CreateRenta(c *fiber.Ctx) error {

	var renta models.Renta

	if err := c.BodyParser(&renta); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.Service.Create(&renta); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(renta)
}

// Obtener todas las rentas
func (h *RentaHandler) GetRenta(c *fiber.Ctx) error {

	renta, err := h.Service.FindAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(renta)
}

// Obtener renta por ID
func (h *RentaHandler) GetRentaByID(c *fiber.Ctx) error {

	id := c.Params("id")

	renta, err := h.Service.FindByID(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Renta no encontrada",
		})
	}

	return c.JSON(renta)
}

// Actualizar renta
func (h *RentaHandler) UpdateRenta(c *fiber.Ctx) error {

	id := c.Params("id")

	var renta models.Renta

	if err := c.BodyParser(&renta); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.Service.Update(id, &renta); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Renta actualizada correctamente",
	})
}

// Eliminar renta
func (h *RentaHandler) DeleteRenta(c *fiber.Ctx) error {

	id := c.Params("id")

	if err := h.Service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Renta eliminada correctamente",
	})
}