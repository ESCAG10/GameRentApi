package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type CategoriaHandler struct {
    CategoriaService *services.CategoriaService
}

func NewCategoriaHandler(service *services.CategoriaService) *CategoriaHandler {
    return &CategoriaHandler{CategoriaService: service}
}

func (h *CategoriaHandler) CreateCategoria(c *fiber.Ctx) error {
    var categoria models.Categoria
    if err := c.BodyParser(&categoria); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
    }
    if err := h.CategoriaService.Create(&categoria); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusCreated).JSON(categoria)
}

func (h *CategoriaHandler) GetCategoria(c *fiber.Ctx) error {
    categoria, err := h.CategoriaService.FindAll()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(categoria)
}

func (h *CategoriaHandler) GetCategoriaByID(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    categoria, err := h.CategoriaService.FindByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Categoría no encontrada"})
    }
    return c.JSON(categoria)
}

func (h *CategoriaHandler) UpdateCategoria(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    var categoria models.Categoria
    if err := c.BodyParser(&categoria); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
    }
    if err := h.CategoriaService.Update(id, &categoria); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Categoría actualizada correctamente"})
}

func (h *CategoriaHandler) DeleteCategoria(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    if err := h.CategoriaService.Delete(id); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Categoría eliminada correctamente"})
}
