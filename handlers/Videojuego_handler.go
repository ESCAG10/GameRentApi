package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)

type VideojuegoHandler struct {
	Service *services.VideojuegoService
}

func NewVideojuegoHandler(service *services.VideojuegoService) *VideojuegoHandler {
	return &VideojuegoHandler{
		Service: service,
	}
}

// Crear videojuego
func (h *VideojuegoHandler) CreateVideojuego(c *fiber.Ctx) error {

	var videojuego models.Videojuego

	if err := c.BodyParser(&videojuego); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.Service.Create(&videojuego); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(videojuego)
}

// Obtener todos los videojuegos
func (h *VideojuegoHandler) GetVideojuego(c *fiber.Ctx) error {

	videojuego, err := h.Service.FindAll()

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(videojuego)
}

func (h *VideojuegoHandler) GetVideojuegoActivos(c *fiber.Ctx) error {

	videojuego, err := h.Service.FindActivos()

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(videojuego)
}

// Obtener videojuego por ID
func (h *VideojuegoHandler) GetVideojuegoByID(c *fiber.Ctx) error {

	id := c.Params("id")

	videojuego, err := h.Service.FindByID(id)

	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "Videojuego no encontrado",
		})
	}

	return c.JSON(videojuego)
}

// Actualizar videojuego
func (h *VideojuegoHandler) UpdateVideojuego(c *fiber.Ctx) error {

	id := c.Params("id")

	var videojuego models.Videojuego

	if err := c.BodyParser(&videojuego); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Datos inválidos",
		})
	}

	if err := h.Service.Update(id, &videojuego); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Videojuego actualizado correctamente",
	})
}

// Eliminar videojuego
func (h *VideojuegoHandler) DeleteVideojuego(c *fiber.Ctx) error {

	id := c.Params("id")

	if err := h.Service.Delete(id); err != nil {
	return c.Status(404).JSON(
		fiber.Map{
			"error": "Videojuego no encontrado",
		},
	)
}

	if err := h.Service.Delete(id); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Videojuego eliminado correctamente",
	})

}