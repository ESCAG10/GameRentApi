package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type CategoriaHandler struct {
	Service *services.CategoriaService
}

func NewCategoriaHandler(service *services.CategoriaService) *CategoriaHandler {
	return &CategoriaHandler{
		Service: service,
	}
}

// Crear categoría
func (h *CategoriaHandler) CreateCategoria(c *fiber.Ctx) error {

	var categoria models.Categoria

	if err := c.BodyParser(&categoria); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.Service.Create(&categoria); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(categoria)
}

// Obtener todas
func (h *CategoriaHandler) GetCategoria(c *fiber.Ctx) error {

	categoria, err := h.Service.FindAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(categoria)
}

// Obtener por ID
func (h *CategoriaHandler) GetCategoriaByID(c *fiber.Ctx) error {

	id := c.Params("id")

	categoria, err := h.Service.FindByID(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Categoría no encontrada",
		})
	}

	return c.JSON(categoria)
}

// Actualizar categoría
func (h *CategoriaHandler) UpdateCategoria(c *fiber.Ctx) error {

	id := c.Params("id")

	var categoria models.Categoria

	if err := c.BodyParser(&categoria); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.Service.Update(id, &categoria); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Categoría actualizada correctamente",
	})
}

// Eliminar categoría
func (h *CategoriaHandler) DeleteCategoria(c *fiber.Ctx) error {

	id := c.Params("id")

	if err := h.Service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Categoría eliminada correctamente",
	})
}