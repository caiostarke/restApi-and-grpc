package routes

import (
	"github.com/caiostarke/restApi-and-grpc/app/controllers"
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	// Routes for GET method: Books
	route.Get("/books", controllers.GetBooks)             // get list of all books
	route.Get("/book/:id", controllers.GetBook)           // get one book by ID
	route.Get("/token/new", controllers.GetNewAcessToken) // create a new access tokens

	// Routes for Users
	route.Post("/signup", controllers.CreateUser) // Create a new User
	route.Post("/login", controllers.Login)       // Login
}
