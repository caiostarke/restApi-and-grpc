package main

import (
	"github.com/caiostarke/restApi-and-grpc/pkg/configs"
	"github.com/caiostarke/restApi-and-grpc/pkg/middleware"
	"github.com/caiostarke/restApi-and-grpc/pkg/routes"
	"github.com/gofiber/fiber/v2"

	_ "github.com/caiostarke/restApi-and-grpc/docs"
	_ "github.com/joho/godotenv/autoload"
)

// @Library API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	config := configs.FiberConfig()

	app := fiber.New(config)
	middleware.FiberMiddleware(app)

	routes.PublicRoutes(app)
	routes.PrivateRoutes(app)
	routes.SwaggerRoute(app)

	app.Listen(":8080")
}
