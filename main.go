package main

import (
	"github.com/caiostarke/restApi-and-grpc/pkg/configs"
	"github.com/caiostarke/restApi-and-grpc/pkg/middleware"
	"github.com/caiostarke/restApi-and-grpc/pkg/routes"
	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)
	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)

	app.Listen(":8080")
}
