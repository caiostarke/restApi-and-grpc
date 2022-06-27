package controllers

import (
	"time"

	"github.com/caiostarke/restApi-and-grpc/app/models"
	"github.com/caiostarke/restApi-and-grpc/platform/database"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get all books.
	books, err := db.GetBooks()
	if err != nil {
		// Return, if books not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "books were not found",
			"count": 0,
			"books": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(books),
		"books": books,
	})

}

func CreateBook(c *fiber.Ctx) error {
	// Get now time.
	// now := time.Now().Unix()

	// // Get claims from JWT.
	// claims, err := utils.ExtractTokenMetadata(c)
	// if err != nil {
	// 	// Return status 500 and JWT parse error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// // Set expiration time from JWT data of current book.
	// expires := claims.Expires

	// // Checking, if now time greather than expiration from JWT.
	// if now > expires {
	// 	// Return status 401 and unauthorized error message.
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "unauthorized, check expiration time of your token",
	// 	})
	// }

	// Create new Book struct
	book := &models.Book{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(book); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set initialized default data for book:
	book.CreatedAt = time.Now()
	book.BookStatus = 1 // 0 == draft, 1 == active

	// Delete book by given ID.
	if err := db.CreateBook(book); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})
}

func GetBook(c *fiber.Ctx) error {
	// Catch book ID from URL.
	id := c.Params("id")

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Get book by ID.
	book, err := db.GetBook(id)
	if err != nil {
		// Return, if book not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "book with the given ID is not found",
			"book":  nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"book":  book,
	})

}

func UpdateBook(c *fiber.Ctx) error {
	// Get now time.
	// now := time.Now().Unix()

	// // Get claims from JWT.
	// claims, err := utils.ExtractTokenMetadata(c)
	// if err != nil {
	// 	// Return status 500 and JWT parse error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// // Set expiration time from JWT data of current book.
	// expires := claims.Expires

	// // Checking, if now time greather than expiration from JWT.
	// if now > expires {
	// 	// Return status 401 and unauthorized error message.
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "unauthorized, check expiration time of your token",
	// 	})
	// }

	// Create new Book struct
	book := &models.Book{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(book); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if book with given ID is exists.
	foundedBook, err := db.GetBook(book.ID.Hex())
	if err != nil {
		// Return status 404 and book not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "book with this ID not found",
		})
	}

	// Set initialized default data for book:
	book.UpdatedAt = time.Now()

	// Create a new validator for a Book model.
	// validate := utils.NewValidator()

	// Validate book fields.
	// if err := validate.Struct(book); err != nil {
	// 	// Return, if some fields are not valid.
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   utils.ValidatorErrors(err),
	// 	})
	// }

	// Update book by given ID.
	if err := db.UpdateBook(foundedBook.ID, book); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201.
	return c.SendStatus(fiber.StatusCreated)

}

func DeleteBook(c *fiber.Ctx) error {
	// Get now time.
	// now := time.Now().Unix()

	// // Get claims from JWT.
	// claims, err := utils.ExtractTokenMetadata(c)
	// if err != nil {
	// 	// Return status 500 and JWT parse error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// // Set expiration time from JWT data of current book.
	// expires := claims.Expires

	// // Checking, if now time greather than expiration from JWT.
	// if now > expires {
	// 	// Return status 401 and unauthorized error message.
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   "unauthorized, check expiration time of your token",
	// 	})
	// }

	// Create new Book struct
	book := &models.Book{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(book); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Create a new validator for a Book model.
	// validate := utils.NewValidator()

	// Validate only one book field ID.
	// if err := validate.StructPartial(book, "id"); err != nil {
	// 	// Return, if some fields are not valid.
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   utils.ValidatorErrors(err),
	// 	})
	// }

	// Create database connection.
	db, err := database.OpenDBConnection()
	if err != nil {
		// Return status 500 and database connection error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Checking, if book with given ID is exists.
	foundedBook, err := db.GetBook(book.ID.Hex())
	if err != nil {
		// Return status 404 and book not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "book with this ID not found",
		})
	}

	// Delete book by given ID.
	if err := db.DeleteBook(foundedBook.ID); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}
