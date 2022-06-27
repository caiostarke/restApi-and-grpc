package main

import (
	"github.com/caiostarke/restApi-and-grpc/pkg/routes"
	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()

	routes.PublicRoutes(app)

	app.Listen(":8080")
}
