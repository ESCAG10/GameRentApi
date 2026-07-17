package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type UsuarioHandler struct {
    UsuarioService *services.UsuarioService
}

func NewUsuarioHandler(service *services.UsuarioService) *UsuarioHandler {
    return &UsuarioHandler{
        UsuarioService: service,
    }
}

func (h *UsuarioHandler) CreateUsuario(c *fiber.Ctx) error {
    var usuario models.Usuario

    if err := c.BodyParser(&usuario); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Datos inválidos",
        })
    }

    if err := h.UsuarioService.Create(&usuario); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.Status(fiber.StatusCreated).JSON(usuario)
}

func (h *UsuarioHandler) GetUsuario(c *fiber.Ctx) error {
    usuario, err := h.UsuarioService.FindAll()
    if err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return c.JSON(usuario)
}

func (h *UsuarioHandler) GetUsuarioByID(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "ID inválido",
        })
    }

    usuario, err := h.UsuarioService.FindByID(id)
    if err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Usuario no encontrado",
        })
    }
    return c.JSON(usuario)
}

func (h *UsuarioHandler) UpdateUsuario(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "ID inválido",
        })
    }

    var usuario models.Usuario
    if err := c.BodyParser(&usuario); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Datos inválidos",
        })
    }

    if err := h.UsuarioService.Update(id, &usuario); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "message": "Usuario actualizado correctamente",
    })
}

func (h *UsuarioHandler) DeleteUsuario(c *fiber.Ctx) error {
    idHex := c.Params("id")
    id, err := bson.ObjectIDFromHex(idHex)
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "ID inválido",
        })
    }

    if err := h.UsuarioService.Delete(id); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": err.Error(),
        })
    }

    return c.JSON(fiber.Map{
        "message": "Usuario eliminado correctamente",
    })
}
