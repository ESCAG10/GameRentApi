package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type VideojuegoHandler struct {
	VideojuegoService *services.VideojuegoService
}

func NewVideojuegoHandler(service *services.VideojuegoService) *VideojuegoHandler {
	return &VideojuegoHandler{
		VideojuegoService: service,
	}
}

func (h *VideojuegoHandler) CreateVideojuego(c *fiber.Ctx) error {

	var videojuego models.Videojuego

	if err := c.BodyParser(&videojuego); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	err := h.VideojuegoService.Create(&videojuego)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Videojuego creado correctamente",
	})
}