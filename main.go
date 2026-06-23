package main

import (
	"log"

	"gamerentapi/config"
	"gamerentapi/handlers"
	"gamerentapi/routes"

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

	// create a UsuarioHandler instance (adjust initialization if your handler requires dependencies)
	usuarioHandler := &handlers.UsuarioHandler{}
	routes.SetupRoutes(app, usuarioHandler)

	log.Fatal(app.Listen(":3000"))
}