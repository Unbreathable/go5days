package main

/*
type Receiver struct {
	Token          string
	CurrentSender *Sender // The current attempt
}

// The current receiver
var CurrentReceiver *Receiver = nil

// Returns a new receiver or false when the receiver has already been set
func ClaimReceiver() (*Receiver, bool) {

	// Check if the receiver has already been claimed
	if CurrentReceiver != nil {
		return nil, false
	}

	// Create a new receiver
	CurrentReceiver = &Receiver{
		Token: generateToken(12),
	}

	return CurrentReceiver, true
}

type Sender struct {
	Token     string // Token to identify the attempt
	Name      string // Name of the sender
	Challenge string // Code the sender has to enter
}

func (r *Receiver) NewSender(name string) *Sender {

	// Create a new attempt
	sender := &Sender{
		Token:     generateToken(12),
		Name:      name,
		Challenge: generateNumbers(6),
	}

	// Register the attempt in the receiver
	CurrentReceiver.CurrentSender = sender

	return attempt
}

func (r *Receiver) CheckAttempt(token string, code string) bool {
	if CurrentReceiver == nil {
		return false
	}
	if CurrentReceiver.CurrentAttempt == nil {
		return false
	}
	if CurrentReceiver.CurrentAttempt.Token != token || CurrentReceiver.CurrentAttempt.Challenge != code {
		return false
	}
	return true
}

*/
