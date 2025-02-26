package main

/*
// Route: /receiver/create
func createReceiver(c *fiber.Ctx) error {

	// Claim the receiver
	receiver, valid := claimReciever()
	if !valid {
		return c.SendStatus(fiber.StatusConflict)
	}

	return c.JSON(fiber.Map{
		"token": receiver.Token,
	})
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

	// Check if the receiver has already been claimed
	if currentReceiver != nil {
		return nil, false
	}

	// Create a new receiver
	currentReceiver = &Receiver{
		Mutex: &sync.Mutex{},
		Token: generateToken(12),
	}

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

	// Check if the token provided is correct
	if currentReceiver == nil {
		return c.SendStatus(fiber.StatusBadGateway)
	}

	// Make sure there are no concurrent errors
	currentReceiver.Mutex.Lock()
	defer currentReceiver.Mutex.Unlock()

	// Check the token
	if currentReceiver.Token != req.Token {
		return c.SendStatus(fiber.StatusForbidden)
	}

	// Check if there is a current attempt
	if currentReceiver.currentSender != nil {
		return c.JSON(fiber.Map{
			"exists":    true,
			"completed": currentReceiver.currentSender.Connected,
			"accepted":  currentReceiver.currentSender.Accepted,
			"name":      currentReceiver.currentSender.Name,
			"code":      currentReceiver.currentSender.Challenge,
		})
	}

	// If there isn't one, return that nothing exists
	return c.JSON(fiber.Map{
		"exists": false,
	})
}

*/
