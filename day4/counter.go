package main

import "sync"

// ID -> *Counter
var counterMap *sync.Map = &sync.Map{}

// Get a counter by its id or create a new one starting at 0
func getCounter(id string) *Counter {
	// TODO: Implement
	return &Counter{}
}

type Counter struct {
	// Add all your state here
}

// Increment the counter by one
func (c *Counter) Increment() {
	// TODO: Implement
}

func (c *Counter) GetValue() int64 {
	// TODO: Implement
	return 0
}
