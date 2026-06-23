package routes

import (
	"gamerentapi/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	usuarioHandler *handlers.UsuarioHandler,
) {

	app.Get("/usuarios", usuarioHandler.GetUsuarios)

	app.Post("/usuarios", usuarioHandler.CreateUsuario)
}