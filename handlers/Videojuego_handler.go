package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type VideojuegoHandler struct {
    VideojuegoService *services.VideojuegoService
}

func NewVideojuegoHandler(service *services.VideojuegoService) *VideojuegoHandler {
    return &VideojuegoHandler{VideojuegoService: service}
}

func (h *VideojuegoHandler) CreateVideojuego(c *fiber.Ctx) error {
    var videojuego models.Videojuego
    if err := c.BodyParser(&videojuego); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
    }
    if err := h.VideojuegoService.Create(&videojuego); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusCreated).JSON(videojuego)
}

func (h *VideojuegoHandler) GetVideojuego(c *fiber.Ctx) error {
    videojuego, err := h.VideojuegoService.FindAll()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(videojuego)
}

func (h *VideojuegoHandler) GetVideojuegoByID(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    videojuego, err := h.VideojuegoService.FindByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Videojuego no encontrado"})
    }
    return c.JSON(videojuego)
}

func (h *VideojuegoHandler) UpdateVideojuego(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    var videojuego models.Videojuego
    if err := c.BodyParser(&videojuego); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
    }
    if err := h.VideojuegoService.Update(id, &videojuego); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Videojuego actualizado correctamente"})
}

func (h *VideojuegoHandler) DeleteVideojuego(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    if err := h.VideojuegoService.Delete(id); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Videojuego eliminado correctamente"})
}
