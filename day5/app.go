package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := setupApp()
	app.Listen("127.0.0.1:3000")
}

func setupApp() *fiber.App {
	app := fiber.New()

	// Configure the basic shit
	app.Use(logger.New())

	// Register all the endpoints
	app.Post("/receiver/create", createReceiver)
	app.Post("/receiver/check_state", checkReceiverState)
	app.Post("/sender/create", createSender)
	app.Post("/sender/attempt", checkAttempt)

	return app
}
