package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type CategoriaHandler struct {
	Service *services.CategoriaService
}

func NewCategoriaHandler(service *services.CategoriaService,) *CategoriaHandler {
	return &CategoriaHandler{
		Service: service,
	}
}

func (h *CategoriaHandler) CreateCategoria(c *fiber.Ctx,) error {

	var categoria models.Categoria

	if err := c.BodyParser(&categoria); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "datos inválidos",
		})
	}

	err := h.Service.Create(&categoria)

	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.Status(201).JSON(categoria)
}

func (h *CategoriaHandler) GetCategorias(c *fiber.Ctx,) error {

	categorias, err := h.Service.FindAll()

	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(categorias)
}

func (h *CategoriaHandler) GetCategoria(c *fiber.Ctx,) error {

	id := c.Params("id")

	categoria, err := h.Service.FindByID(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "categoria no encontrada",
		})
	}

	return c.JSON(categoria)
}

func (h *CategoriaHandler) UpdateCategoria(c *fiber.Ctx,) error {

	id := c.Params("id")

	var categoria models.Categoria

	if err := c.BodyParser(&categoria); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "datos inválidos",
		})
	}

	err := h.Service.Update(id,&categoria,)

	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(fiber.Map{
		"message": "categoria actualizada",
	})
}

func (h *CategoriaHandler) DeleteCategoria(c *fiber.Ctx,) error {

	id := c.Params("id")

	err := h.Service.Delete(id)

	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(fiber.Map{
		"message": "categoria eliminada",
	})
}