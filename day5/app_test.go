package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

// Helper function for JSON requests
func jsonRequest(method, url string, body interface{}) *http.Request {
	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest(method, url, strings.NewReader(string(jsonBody)))
	req.Header.Add("Content-Type", "application/json")
	return req
}

func TestReceiver(t *testing.T) {
	app := setupApp()

	// Test creating a new receiver
	t.Run("test creation of receiver", func(t *testing.T) {

		// Send the request
		resp, err := app.Test(httptest.NewRequest("POST", "/receiver/create", nil))
		if err != nil {
			t.Fatal("error during request:", err)
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("error during body reading:", err)
		}

		// Parse the response as a json
		var res struct {
			Token string `json:"token"`
		}
		if err := json.Unmarshal(body, &res); err != nil {
			t.Fatal(err)
		}

		// Make sure the token is valid
		if res.Token != currentReceiver.Token {
			t.Errorf("Expected token %s, got '%s'", currentReceiver.Token, res.Token)
		}
	})

	// Make sure another request fails
	t.Run("test creating two receivers failing", func(t *testing.T) {

		// Send the request
		resp, err := app.Test(httptest.NewRequest("POST", "/receiver/create", nil))
		if err != nil {
			t.Fatal("error during request:", err)
		}
		if resp.StatusCode == 200 {
			t.Error("Expected no response when there already is a receiver")
		}
	})

	// Make sure there is no sender yet
	t.Run("test check state returning correct stuff", func(t *testing.T) {

		// Check if the sender is there
		if currentReceiver.currentSender != nil {
			t.Error("Sender exists even though hasn't been created yet")
		}

		// Send the request
		req := httptest.NewRequest("POST", "/receiver/check_state", strings.NewReader(fmt.Sprintf(`{"token":"%s"}`, currentReceiver.Token)))
		req.Header.Add("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("error during request:", err)
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal("error during body reading:", err)
		}

		// Parse the response as a json
		var res struct {
			Exists bool `json:"exists"`
		}
		if err := json.Unmarshal(body, &res); err != nil {
			t.Fatal("error while parsing json:", err)
		}

		// Make the sender doesn't exist yet
		if res.Exists {
			t.Errorf("Expected exists to be false, got '%v'", res.Exists)
		}
	})

	// Make sure it doesn't work with an invalid token
	t.Run("test check state with invalid token", func(t *testing.T) {

		// Check if the sender is there
		if currentReceiver.currentSender != nil {
			t.Error("Sender exists even though hasn't been created yet")
		}

		// Send the request
		req := httptest.NewRequest("POST", "/receiver/check_state", strings.NewReader(`{"token":"123"}`))
		req.Header.Add("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("error during request:", err)
		}
		if resp.StatusCode == 200 {
			t.Error("Expected no response with invalid token")
		}
	})

	// Test sender creation
	t.Run("test sender creation", func(t *testing.T) {
		req := jsonRequest("POST", "/sender/create", map[string]string{
			"name": "TestSender",
		})
		resp, err := app.Test(req)
		if err != nil {
			t.Fatal("error during request:", err)
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}

		var res struct {
			Token string `json:"token"`
		}
		body, _ := io.ReadAll(resp.Body)
		if err := json.Unmarshal(body, &res); err != nil {
			t.Fatal("error parsing json:", err)
		}

		if res.Token == "" {
			t.Error("Expected token in response")
		}

		// Verify sender was created
		if currentReceiver.currentSender == nil {
			t.Error("Sender was not created in receiver")
		}
		if currentReceiver.currentSender.Name != "TestSender" {
			t.Errorf("Expected sender name TestSender, got %s", currentReceiver.currentSender.Name)
		}
	})

	// Test receiver state after sender creation
	t.Run("test receiver state with sender", func(t *testing.T) {
		req := jsonRequest("POST", "/receiver/check_state", map[string]string{
			"token": currentReceiver.Token,
		})
		resp, _ := app.Test(req)

		var state struct {
			Exists   bool   `json:"exists"`
			Name     string `json:"name"`
			Code     string `json:"code"`
			Accepted bool   `json:"accepted"`
		}
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &state)

		if !state.Exists {
			t.Error("Expected sender to exist in state")
		}
		if state.Name != "TestSender" {
			t.Errorf("Expected sender name TestSender, got %s", state.Name)
		}
		if state.Code == "" {
			t.Error("Expected challenge code in state")
		}
	})

	// Test sender attempt
	t.Run("test sender attempt", func(t *testing.T) {
		senderToken := currentReceiver.currentSender.Token
		challenge := currentReceiver.currentSender.Challenge

		// Test invalid code
		req := jsonRequest("POST", "/sender/attempt", map[string]string{
			"token": senderToken,
			"code":  "000000",
		})
		resp, _ := app.Test(req)
		if resp.StatusCode != fiber.StatusForbidden {
			t.Errorf("Expected status 403 for wrong code, got %d", resp.StatusCode)
		}

		// Test correct code
		req = jsonRequest("POST", "/sender/attempt", map[string]string{
			"token": senderToken,
			"code":  challenge,
		})
		resp, _ = app.Test(req)
		if resp.StatusCode != fiber.StatusOK {
			t.Errorf("Expected status 200 for correct code, got %d", resp.StatusCode)
		}

		// Verify state update
		req = jsonRequest("POST", "/receiver/check_state", map[string]string{
			"token": currentReceiver.Token,
		})
		resp, _ = app.Test(req)
		var state struct {
			Exists   bool `json:"exists"`
			Accepted bool `json:"accepted"`
		}
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &state)

		if !state.Accepted {
			t.Error("Expected sender to be accepted after correct code")
		}
	})
}
