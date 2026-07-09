package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
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
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	err := h.UsuarioService.Create(&usuario)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "Usuario creado correctamente",
	})
}

func (h *UsuarioHandler) GetUsuarios(c *fiber.Ctx) error {

	usuarios, err := h.UsuarioService.FindAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(usuarios)
}

func (h *UsuarioHandler) GetUsuarioByID(c *fiber.Ctx) error {
	id := c.Params("id")

	usuario, err := h.UsuarioService.FindByID(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(usuario)
}

func (h *UsuarioHandler) UpdateUsuario(c *fiber.Ctx) error {
	id := c.Params("id")

	var usuario models.Usuario

	if err := c.BodyParser(&usuario); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	err := h.UsuarioService.Update(id, &usuario)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Usuario actualizado correctamente",
	})
}

func (h *UsuarioHandler) DeleteUsuario(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.UsuarioService.Delete(id)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Usuario eliminado correctamente",
	})
}