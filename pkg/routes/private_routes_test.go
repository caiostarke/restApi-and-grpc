package routes

import (
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/caiostarke/restApi-and-grpc/app/models"
	"github.com/caiostarke/restApi-and-grpc/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestPrivateRoutes(t *testing.T) {
	if err := godotenv.Load("../../.env.test"); err != nil {
		panic(err)
	}
	userTest := models.UserResponse{}

	dataString := `{
		"id": "62bc54a2ff197412ae97d03d",
		"username":"random",
		"email": "random.rand.rrr"
	}`

	token, err := utils.GenerateNewAccessToken(userTest)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		description   string
		route         string // input route
		method        string // input method
		tokenString   string // input token
		body          io.Reader
		expectedError bool
		expectedCode  int
	}{
		{
			description:   "delete book without JWT and body",
			route:         "/api/v1/book",
			method:        "DELETE",
			tokenString:   "",
			body:          nil,
			expectedError: false,
			expectedCode:  400,
		},
		{
			description:   "delete book without right credentials",
			route:         "/api/v1/book",
			method:        "DELETE",
			tokenString:   "Bearer" + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  401,
		},
		{
			description:   "delete book with credentials",
			route:         "/api/v1/book",
			method:        "DELETE",
			tokenString:   "Bearer " + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  404,
		},
		{
			description:   "delete user without credentials",
			route:         "/api/v1/user",
			method:        "DELETE",
			tokenString:   "Bearer " + token,
			body:          strings.NewReader(dataString),
			expectedError: false,
			expectedCode:  401,
		},
	}

	app := fiber.New()

	PrivateRoutes(app)

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Authorization", test.tokenString)
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}
}
