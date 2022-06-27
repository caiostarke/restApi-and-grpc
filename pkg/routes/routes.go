package routes

import (
	"github.com/caiostarke/restApi-and-grpc/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/books", controllers.GetBooks)
	route.Post("/book", controllers.CreateBook) // create a new book
	route.Get("/book/:id", controllers.GetBook)
	route.Put("/book", controllers.UpdateBook)    // update one book by ID
	route.Delete("/book", controllers.DeleteBook) // delete one book by ID

	// route.Get("/token/new", controllers.GetNewAcessToken)
}
