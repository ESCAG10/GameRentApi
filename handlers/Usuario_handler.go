package handlers

import (
	"fmt"
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

	fmt.Println(string(c.Body()))

var usuario models.Usuario

if err := c.BodyParser(&usuario); err != nil {

    return c.Status(400).JSON(
        fiber.Map{
            "error": err.Error(),
        },
    )
}

fmt.Println("===== STRUCT USUARIO =====")
fmt.Printf("%+v\n", usuario)
fmt.Println("Nombre:", usuario.Nombre)
fmt.Println("Correo:", usuario.Correo)
fmt.Println("Password:", usuario.Password)
fmt.Println("Rol:", usuario.Rol)
fmt.Println("Activo:", usuario.Activo)
fmt.Println("==========================")

if err := h.Service.Create(&usuario); err != nil {

	return c.Status(500).JSON(
		fiber.Map{
			"error": err.Error(),
		},
	)
}

// No devolver la contraseña
usuario.Password = ""

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

	for i := range usuarios {
		usuarios[i].Password = ""
	}

	return c.Status(200).JSON(usuarios)
}



// Obtener usuario por ID
func (h *UsuarioHandler) GetUsuarioByID(c *fiber.Ctx,) error {


	id := c.Params("id")


	usuario, err := h.Service.FindByID(id)
	
	if err != nil {

		

		return c.Status(404).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	// No devolver la contraseña
	usuario.Password = ""

	return c.Status(200).JSON(usuario)
}



// Actualizar usuario
func (h *UsuarioHandler) UpdateUsuario(c *fiber.Ctx,) error {


	id := c.Params("id")


	var usuario models.Usuario


	if err := c.BodyParser(&usuario); err != nil {

    return c.Status(400).JSON(
        fiber.Map{
            "error": err.Error(),
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

func (h *UsuarioHandler) Login(c *fiber.Ctx) error {
	var request struct {
		Correo   string `json:"correo"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "Datos inválidos",
			},
		)
	}

	usuario, err := h.Service.Login(
		request.Correo,
		request.Password,
	)

	if err != nil {
		return c.Status(401).JSON(
			fiber.Map{
				"error": "Credenciales inválidas",
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"message": "Login exitoso",
			"usuario": fiber.Map{
				"id": usuario.ID.Hex(),
				"nombre": usuario.Nombre,
				"correo": usuario.Correo,
				"rol": usuario.Rol,
			},
		},
	)
}
