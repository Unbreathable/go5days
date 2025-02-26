package main

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Route: /sender/create
func createSender(c *fiber.Ctx) error {

	// Parse the request
	var req struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// TODO: Implement

	// Return a new attempt token
	return c.SendStatus(fiber.StatusOK)
}

type Sender struct {
	Mutex     *sync.Mutex // For not having concurrent writes
	Token     string      // Token to identify the attempt
	Name      string      // Name of the sender
	Challenge string      // Code the sender has to enter
	Accepted  bool        // Whether the attempt has been accepted
	Connected bool        // Whether the sender is connected or not
}

func (r *Receiver) MakeAttempt(name string) *Sender {

	// TODO: Port old function

	return nil
}

// Route: /sender/attempt
func checkAttempt(c *fiber.Ctx) error {

	// Parse the request
	var req struct {
		Token string `json:"token"`
		Code  string `json:"code"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Make sure there is a receiver
	if currentReceiver == nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// TODO: Implement

	// Send that everything worked, they can now use the code to create a new connection
	return c.SendStatus(fiber.StatusOK)
}
