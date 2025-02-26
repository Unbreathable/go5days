package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Whether or not the API should be tested, turn this to true once finished with API
const APITests = false

// Start the server
func main() {
	app := setupApp(false)
	app.Listen("127.0.0.1:3000")
}

// This function needs to be separate because we want to test the app
func setupApp(testing bool) *fiber.App {
	app := fiber.New()
	if !testing {
		app.Use(logger.New())
	}

	// TODO: Implement your endpoints here

	// Return the app
	return app
}
