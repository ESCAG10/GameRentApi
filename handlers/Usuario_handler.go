package handlers

import (
	"gamerentapi/models"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
)


type UsuarioHandler struct {
	Service *services.UsuarioService
}


func NewUsuarioHandler(service *services.UsuarioService,) *UsuarioHandler {

	return &UsuarioHandler{
		Service: service,
	}
}


// Crear usuario
func (h *UsuarioHandler) CreateUsuario(c *fiber.Ctx,) error {


	var usuario models.Usuario


	if err := c.BodyParser(&usuario); err != nil {

		return c.Status(400).JSON(
			fiber.Map{
				"error": "Datos inválidos",
			},
		)
	}


	if err := h.Service.Create(&usuario); err != nil {

		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}


	return c.Status(201).JSON(usuario)
}



// Obtener todos los usuarios
func (h *UsuarioHandler) GetUsuario(c *fiber.Ctx,) error {


	usuarios, err := h.Service.FindAll()

	if err != nil {

		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}


	return c.JSON(usuarios)
}



// Obtener usuario por ID
func (h *UsuarioHandler) GetUsuarioByID(c *fiber.Ctx,) error {


	id := c.Params("id")


	usuario, err := h.Service.FindByID(id)


	if err != nil {

		return c.Status(404).JSON(
			fiber.Map{
				"error": "Usuario no encontrado",
			},
		)
	}


	return c.JSON(usuario)
}



// Actualizar usuario
func (h *UsuarioHandler) UpdateUsuario(c *fiber.Ctx,) error {


	id := c.Params("id")


	var usuario models.Usuario


	if err := c.BodyParser(&usuario); err != nil {

		return c.Status(400).JSON(
			fiber.Map{
				"error": "Datos inválidos",
			},
		)
	}


	if err := h.Service.Update(
		id,
		&usuario,
	); err != nil {


		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}


	return c.JSON(
		fiber.Map{
			"message": "Usuario actualizado correctamente",
		},
	)
}



// Eliminar usuario
func (h *UsuarioHandler) DeleteUsuario(c *fiber.Ctx,) error {


	id := c.Params("id")


	if err := h.Service.Delete(id); err != nil {


		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}


	return c.JSON(
		fiber.Map{
			"message": "Usuario eliminado correctamente",
		},
	)
}
