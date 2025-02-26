package main

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Route: /receiver/create
func createReceiver(c *fiber.Ctx) error {

	// TODO: Implement

	return c.SendStatus(fiber.StatusOK)
}

type Receiver struct {
	Mutex         *sync.Mutex // For not having concurrent writes
	Token         string
	currentSender *Sender // The current attempt
}

// The current receiver
var currentReceiver *Receiver = nil

// Returns whether or not the receiver has been claimed
func claimReciever() (*Receiver, bool) {

	// TODO: Port old method

	return currentReceiver, true
}

// Endpoint for the receiver to check for a code
// Route: /receiver/check_state
func checkReceiverState(c *fiber.Ctx) error {

	// Parse the request
	var req struct {
		Token string `json:"token"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// TODO: Return correct stuff

	// If there isn't one, return that nothing exists
	return c.JSON(fiber.Map{
		"exists": false,
	})
}
