package routes

import (
	"gamerentapi/handlers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(
	app *fiber.App,
	usuarioHandler *handlers.UsuarioHandler,
	categoriaHandler *handlers.CategoriaHandler,
	videojuegoHandler *handlers.VideojuegoHandler,
	rentaHandler *handlers.RentaHandler,
	perfilHandler *handlers.PerfilHandler,
) {

	// ==========================
	// Ruta principal
	// ==========================
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"mensaje": "GameRent API funcionando",
		})
	})


	// ==========================
	//Login
	// ==========================

	app.Post("/login", usuarioHandler.Login)

	// ==========================
	// Usuario
	// ==========================

	app.Get("/usuario", usuarioHandler.GetUsuario)
	app.Post("/usuario", usuarioHandler.CreateUsuario)
	app.Get("/usuario/:id", usuarioHandler.GetUsuarioByID)
	app.Put("/usuario/:id", usuarioHandler.UpdateUsuario)
	app.Delete("/usuario/:id", usuarioHandler.DeleteUsuario)


	// ==========================
	// Categoría
	// ==========================
	app.Get("/categoria", categoriaHandler.GetCategoria)
	app.Post("/categoria", categoriaHandler.CreateCategoria)
	app.Get("/categoria/:id", categoriaHandler.GetCategoriaByID)
	app.Put("/categoria/:id", categoriaHandler.UpdateCategoria)
	app.Delete("/categoria/:id", categoriaHandler.DeleteCategoria)


	// ==========================
	// Renta
	// ==========================
	app.Get("/renta", rentaHandler.GetRenta)
	app.Post("/renta", rentaHandler.CreateRenta)
	// Obtener rentas por usuario
	app.Get("/renta/usuario/:id", rentaHandler.GetRentasByUsuarioID)
	// Obtener una renta por ID
	app.Get("/renta/:id", rentaHandler.GetRenta)
	app.Get("/renta/:id", rentaHandler.GetRentaByID)
	app.Put("/renta/:id", rentaHandler.UpdateRenta)
	// Eliminar una renta por ID
	app.Delete("/renta/:id", rentaHandler.DeleteRenta)


	// ==========================
	// Perfil
	// ==========================
	app.Get("/perfil", perfilHandler.GetPerfil)
	app.Post("/perfil", perfilHandler.CreatePerfil)
	app.Get("/perfil/:id", perfilHandler.GetPerfilByID)
	app.Put("/perfil/:id", perfilHandler.UpdatePerfil)
	app.Delete("/perfil/:id", perfilHandler.DeletePerfil)

	// ==========================
	// Videojuego
	// ==========================
	app.Get("/videojuego", videojuegoHandler.GetVideojuego)
	app.Post("/videojuego", videojuegoHandler.CreateVideojuego)
	app.Get("/videojuego/:id", videojuegoHandler.GetVideojuegoByID)
	app.Put("/videojuego/:id", videojuegoHandler.UpdateVideojuego)
	app.Delete("/videojuego/:id", videojuegoHandler.DeleteVideojuego)

}
