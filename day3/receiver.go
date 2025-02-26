package main

type Receiver struct {
	Token          string
	CurrentAttempt *Sender // The current attempt
}

// The current receiver
var CurrentReceiver *Receiver = nil

// Returns a new receiver or false when the receiver has already been set
func ClaimReceiver() (*Receiver, bool) {
	// TODO: Implement
	return nil, true
}

type Sender struct {
	Token     string // Token to identify the attempt
	Name      string // Name of the sender
	Challenge string // Code the sender has to enter
}

func (r *Receiver) NewSender(name string) *Sender {
	// TODO: Implement
	return nil
}

func (r *Receiver) CheckAttempt(token string, code string) bool {
	// TODO: Implement
	return true
}
