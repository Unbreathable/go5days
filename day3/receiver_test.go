package main

import (
	"testing"
)

func resetReceiver() {
	CurrentReceiver = nil
}

func TestClaimReceiver(t *testing.T) {
	resetReceiver()

	// First claim should succeed
	r1, ok1 := ClaimReceiver()
	if !ok1 {
		t.Error("First claim should succeed")
	}
	if r1 == nil {
		t.Error("First claim should return a receiver")
		return
	}
	if len(r1.Token) != 12 {
		t.Errorf("Token length should be 12, got %d", len(r1.Token))
	}

	// Second claim should fail
	r2, ok2 := ClaimReceiver()
	if ok2 {
		t.Error("Second claim should fail")
	}
	if r2 != nil {
		t.Error("Second claim should return nil receiver")
	}
}

func TestMakeAttempt(t *testing.T) {
	resetReceiver()
	r, _ := ClaimReceiver()

	sender := r.NewSender("test_user")

	if sender == nil {
		t.Error("NewSender should return a non-nil attempt")
		return
	}
	if len(sender.Token) != 12 {
		t.Errorf("Attempt token length should be 12, got %d", len(sender.Token))
	}
	if sender.Name != "test_user" {
		t.Errorf("Attempt name should be test_user, got %s", sender.Name)
	}
	if len(sender.Challenge) != 6 {
		t.Errorf("Challenge length should be 6, got %d", len(sender.Challenge))
	}
	if r.CurrentAttempt != sender {
		t.Error("Sender should be registered in receiver")
	}
}

func TestCheckAttempt(t *testing.T) {
	resetReceiver()

	// Should fail when no receiver exists
	if CurrentReceiver.CheckAttempt("token", "123456") {
		t.Error("Check should fail when no receiver exists")
	}

	// Create receiver and attempt
	r, _ := ClaimReceiver()
	sender := r.NewSender("test_user")

	// Test valid attempt
	if !r.CheckAttempt(sender.Token, sender.Challenge) {
		t.Error("Check should succeed with correct token and challenge")
	}

	// Test invalid token
	if r.CheckAttempt("wrong_token", sender.Challenge) {
		t.Error("Check should fail with wrong token")
	}

	// Test invalid challenge
	if r.CheckAttempt(sender.Token, "wrong_challenge") {
		t.Error("Check should fail with wrong challenge")
	}
}
