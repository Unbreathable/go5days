package main

/*
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

	// Endpoint for getting the value of a counter
	app.Get("/:id", func(c *fiber.Ctx) error {

		// Get the id from the URL
		id := c.Params("id", "")

		// Return the value of the counter
		ct := getCounter(id)
		return c.SendString(fmt.Sprintf("%d", ct.GetValue()))
	})

	// Endpoint for incrementing a counter
	app.Post("/:id", func(c *fiber.Ctx) error {

		// Get the id from the URL
		id := c.Params("id", "")

		// Return the value of the counter and increment
		ct := getCounter(id)
		ct.Increment()
		return c.SendString(fmt.Sprintf("%d", ct.GetValue()))
	})

	// Return the app
	return app
}
*/
