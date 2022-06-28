package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestPublicRoutes(t *testing.T) {
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}

	validID := "62bb3a5824a58182b4098967"

	tests := []struct {
		route         string
		description   string
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "get book by ID ",
			route:         "/api/v1/book/" + primitive.NewObjectID().Hex(),
			expectedError: false,
			expectedCode:  404,
		},
		{
			description:   "get book by valid ID ",
			route:         "/api/v1/book/" + validID,
			expectedError: false,
			expectedCode:  200,
		},
	}

	app := fiber.New()
	PublicRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest("GET", test.route, nil)
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
	}
}
