package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type CategoriaHandler struct {
	CategoriaService *services.CategoriaService
}

func NewCategoriaHandler(service *services.CategoriaService) *CategoriaHandler {
	return &CategoriaHandler{
		CategoriaService: service,
	}
}

func (h *CategoriaHandler) CreateCategoria(c *fiber.Ctx) error {

	var categoria models.Categoria

	if err := c.BodyParser(&categoria); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	err := h.CategoriaService.Create(&categoria)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Categoría creada correctamente",
	})
}