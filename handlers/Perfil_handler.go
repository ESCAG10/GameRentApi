package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type PerfilHandler struct {
	Service *services.PerfilService
}

func NewPerfilHandler(service *services.PerfilService) *PerfilHandler {
	return &PerfilHandler{
		Service: service,
	}
}

// Crear perfil
func (h *PerfilHandler) CreatePerfil(c *fiber.Ctx) error {

	var perfil models.Perfil

	if err := c.BodyParser(&perfil); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.Service.Create(&perfil); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(perfil)
}

// Obtener todos
func (h *PerfilHandler) GetPerfil(c *fiber.Ctx) error {

	perfiles, err := h.Service.FindAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(perfiles)
}

// Obtener por ID
func (h *PerfilHandler) GetPerfilByID(c *fiber.Ctx) error {

	id := c.Params("id")

	perfil, err := h.Service.FindByID(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Perfil no encontrado",
		})
	}

	return c.JSON(perfil)
}

// Actualizar
func (h *PerfilHandler) UpdatePerfil(c *fiber.Ctx) error {

	id := c.Params("id")

	var perfil models.Perfil

	if err := c.BodyParser(&perfil); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.Service.Update(id, &perfil); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Perfil actualizado correctamente",
	})
}

// Eliminar
func (h *PerfilHandler) DeletePerfil(c *fiber.Ctx) error {

	id := c.Params("id")

	if err := h.Service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Perfil eliminado correctamente",
	})
}
