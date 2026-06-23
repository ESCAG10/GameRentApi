package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type PerfilHandler struct {
	PerfilService *services.PerfilService
}

func NewPerfilHandler(service *services.PerfilService) *PerfilHandler {
	return &PerfilHandler{
		PerfilService: service,
	}
}

func (h *PerfilHandler) CreatePerfil(c *fiber.Ctx) error {

	var perfil models.Perfil

	if err := c.BodyParser(&perfil); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	err := h.PerfilService.Create(&perfil)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Perfil creado correctamente",
	})
}