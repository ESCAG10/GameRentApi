package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type PerfilHandler struct {
    PerfilService *services.PerfilService
}

func NewPerfilHandler(service *services.PerfilService) *PerfilHandler {
    return &PerfilHandler{PerfilService: service}
}

func (h *PerfilHandler) CreatePerfil(c *fiber.Ctx) error {
    var perfil models.Perfil
    if err := c.BodyParser(&perfil); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
    }
    if err := h.PerfilService.Create(&perfil); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(fiber.StatusCreated).JSON(perfil)
}

func (h *PerfilHandler) GetPerfil(c *fiber.Ctx) error {
    perfil, err := h.PerfilService.FindAll()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(perfil)
}

func (h *PerfilHandler) GetPerfilByID(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    perfil, err := h.PerfilService.FindByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Perfil no encontrado"})
    }
    return c.JSON(perfil)
}

func (h *PerfilHandler) UpdatePerfil(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    var perfil models.Perfil
    if err := c.BodyParser(&perfil); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Datos inválidos"})
    }
    if err := h.PerfilService.Update(id, &perfil); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Perfil actualizado correctamente"})
}

func (h *PerfilHandler) DeletePerfil(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "ID inválido"})
    }
    if err := h.PerfilService.Delete(id); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(fiber.Map{"message": "Perfil eliminado correctamente"})
}
