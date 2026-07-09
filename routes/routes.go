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

	app.Get("/usuarios", usuarioHandler.GetUsuarios)

	app.Post("/usuarios", usuarioHandler.CreateUsuario)
	app.Get("/usuarios/:id", usuarioHandler.GetUsuarioByID)
	app.Put("/usuarios/:id", usuarioHandler.UpdateUsuario)
	app.Delete("/usuarios/:id", usuarioHandler.DeleteUsuario)

	app.Get("/rentas", rentaHandler.GetRentas)
	app.Post("/rentas", rentaHandler.CreateRenta)
	app.Get("/rentas/:id", rentaHandler.GetRenta)
	app.Put("/rentas/:id", rentaHandler.UpdateRenta)
	app.Delete("/rentas/:id", rentaHandler.DeleteRenta)

	
	app.Get("/categorias", categoriaHandler.GetCategorias)
	app.Post("/categorias", categoriaHandler.CreateCategoria)
	app.Get("/categorias/:id", categoriaHandler.GetCategoria)
	app.Put("/categorias/:id", categoriaHandler.UpdateCategoria)
	app.Delete("/categorias/:id", categoriaHandler.DeleteCategoria)

	app.Post("/perfiles", perfilHandler.CreatePerfil)
	app.Get("/perfiles", perfilHandler.GetPerfiles)
	app.Get("/perfiles/:id", perfilHandler.GetPerfil)
	app.Put("/perfiles/:id", perfilHandler.UpdatePerfil)
	app.Delete("/perfiles/:id", perfilHandler.DeletePerfil)

	app.Get("/videojuegos", videojuegoHandler.GetVideojuegos)
	app.Post("/videojuegos", videojuegoHandler.CreateVideojuego)
	app.Get("/videojuegos/:id", videojuegoHandler.GetVideojuego)
	app.Put("/videojuegos/:id", videojuegoHandler.UpdateVideojuego)
	app.Delete("/videojuegos/:id", videojuegoHandler.DeleteVideojuego)


}
