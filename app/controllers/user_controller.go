package controllers

import (
	"fmt"
	"time"

	"github.com/caiostarke/restApi-and-grpc/app/models"
	"github.com/caiostarke/restApi-and-grpc/pkg/utils"
	"github.com/caiostarke/restApi-and-grpc/platform/database"
	"github.com/gofiber/fiber/v2"
)

// GetUser func get an user by giver id.
// @Description Get an user by giver id.
// @Summary Get an user by giver id
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.UserResponse
// @Router /v1/user [get]
func GetUser(c *fiber.Ctx) error {
	// Get id param
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

	// Get all books.
	user, err := db.GetUser(id)
	if err != nil {
		// Return, if books not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user was not found",
			"count": 0,
			"books": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}

// UpdateUser func updates an user.
// @Description UpdateUser func updates an user..
// @Summary UpdateUser func updates an user.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.UserResponse
// @Router /v1/user [put]
func UpdateUser(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires
	user := claims.Username

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new User struct
	userFromUser := &models.UserResponse{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(userFromUser); err != nil {
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

	// Checking, if user with given ID is exists.
	foundedUser, err := db.GetUser(userFromUser.ID)
	if err != nil {
		// Return status 404 and user not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with this ID not found",
		})
	}

	fmt.Println(foundedUser.Username)
	fmt.Println(user)

	if user != foundedUser.Username {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, you cant update other user",
		})
	}

	foundedUser.Email = userFromUser.Email
	foundedUser.Username = userFromUser.Username

	// Update user by given ID.
	if err := db.UpdateUser(foundedUser.ID, &foundedUser); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 201.
	return c.SendStatus(fiber.StatusCreated)
}

// DeleteUser func deletes an user.
// @Description DeleteUser func deletes an user..
// @Summary DeleteUser func deletes an user.
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {array} models.UserResponse
// @Router /v1/user [delete]
func DeleteUser(c *fiber.Ctx) error {
	// Get now time.
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		// Return status 500 and JWT parse error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Set expiration time from JWT data of current book.
	expires := claims.Expires
	usernameFromJWT := claims.Username
	roleFromJWT := claims.Role
	fmt.Println(roleFromJWT)
	fmt.Println(usernameFromJWT)

	// Checking, if now time greather than expiration from JWT.
	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	// Create new User struct
	user := &models.UserResponse{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
		// Return status 400 and error message.
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if usernameFromJWT != user.Username && roleFromJWT != "admin" {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, not user or admin",
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

	// Checking, if user with given ID is exists.
	foundedUser, err := db.GetUser(user.ID)
	if err != nil {
		// Return status 404 and user not found error.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "user with this ID not found",
		})
	}

	// Delete user by given ID.
	if err := db.DeleteUser(foundedUser.ID.Hex()); err != nil {
		// Return status 500 and error message.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 204 no content.
	return c.SendStatus(fiber.StatusNoContent)
}

// CreateUser func for creates a new user.
// @Description Create a new user.
// @Summary create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Param email body string true "Email"
// @Success 200 {object} models.UserResponse
// @Security ApiKeyAuth
// @Router /v1/user [post]
func CreateUser(c *fiber.Ctx) error {
	// Create new User struct
	user := &models.SignUpRequest{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
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

	if err := db.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	userFiltered := models.UserResponse{}
	userFiltered.Email = user.Email
	userFiltered.Username = user.Username

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  userFiltered,
	})
}

func Login(c *fiber.Ctx) error {
	// Create new User struct
	user := &models.LoginRequest{}

	// Check, if received JSON data is valid.
	if err := c.BodyParser(user); err != nil {
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

	userRes, err := db.Login(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   "Login, ID or Password are wrong",
		})
	}

	token, err := utils.GenerateNewAccessToken(userRes)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  userRes,
		"token": token,
	})
}
