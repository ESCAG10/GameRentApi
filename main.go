package main

import (
	"log"

	"gamerentapi/config"
	"gamerentapi/handlers"
	"gamerentapi/repositories"
	"gamerentapi/services"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error cargando archivo .env")
	}

	if err := config.ConnectMongo(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	
	// Usuario
	usuarioCollection := config.Database.Collection("usuarios")
	
	usuarioRepository := repositories.NewUsuarioRepository(usuarioCollection,)
	
	usuarioService := services.NewUsuarioService(usuarioRepository,)
	
	_ = handlers.NewUsuarioHandler(usuarioService,)
	
	// Categoria
	categoriaCollection := config.Database.Collection("categorias")
	
	categoriaRepository := repositories.NewCategoriaRepository(categoriaCollection,)
	
	categoriaService := services.NewCategoriaService(categoriaRepository,)
	
	_ = handlers.NewCategoriaHandler(categoriaService,)
	
	// Videojuego
	videojuegoCollection := config.Database.Collection("videojuegos")
	
	videojuegoRepository := repositories.NewVideojuegoRepository(videojuegoCollection,)
	
	videojuegoService := services.NewVideojuegoService(videojuegoRepository,)
	
	_ = handlers.NewVideojuegoHandler(videojuegoService,)
	
	// Renta
	rentaCollection := config.Database.Collection("rentas")
	
	rentaRepository := repositories.NewRentaRepository(rentaCollection,)
	
	rentaService := services.NewRentaService(rentaRepository,)
	
	_ = handlers.NewRentaHandler(rentaService,)

	//Perfil
	perfilCollection := config.Database.Collection("perfiles")
	
	perfilRepository := repositories.NewPerfilRepository(perfilCollection,)
	
	perfilService := services.NewPerfilService(perfilRepository,)
	
	_ = handlers.NewPerfilHandler(perfilService,)


	log.Println("Servidor iniciado en http://localhost:3000")

	log.Fatal(app.Listen(":3000"))


}