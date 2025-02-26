package main

/*
// Route: /sender/create
func createSender(c *fiber.Ctx) error {

	// Parse the request
	var req struct {
		Name string `json:"name"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Make sure there is a receiver
	if currentReceiver == nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// Lock the mutex
	currentReceiver.Mutex.Lock()
	defer currentReceiver.Mutex.Unlock()

	// Make sure there is a receiver and no current sender
	if currentReceiver.currentSender != nil {
		return c.SendStatus(fiber.StatusConflict)
	}

	// Create a new attempt
	attempt := currentReceiver.MakeAttempt(req.Name)

	// Return a new attempt token
	return c.JSON(fiber.Map{
		"token": attempt.Token,
	})
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

	// Create a new sender
	attempt := &Sender{
		Mutex:     &sync.Mutex{}, // Initialize mutex
		Token:     generateToken(12),
		Name:      name,
		Challenge: generateNumbers(6),
	}

	// Register the sender in the receiver
	currentReceiver.currentSender = attempt

	return attempt
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

	// Lock the mutex
	currentReceiver.Mutex.Lock()
	defer currentReceiver.Mutex.Unlock()

	// Check if the attempt is valid
	if currentReceiver.currentSender == nil {
		return c.SendStatus(fiber.StatusNoContent)
	}

	sender := currentReceiver.currentSender
	sender.Mutex.Lock()
	defer sender.Mutex.Unlock()

	if sender.Token != req.Token || sender.Challenge != req.Code {
		return c.SendStatus(fiber.StatusForbidden)
	}

	// Change the status of the attempt to accepted
	sender.Accepted = true

	// Send that everything worked, they can now use the code to create a new connection
	return c.SendStatus(fiber.StatusOK)
}
*/
