package main

import (
	"crypto/rand"
	"math/big"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// Generate a random string with the specified length
func generateToken(tkLength int32) string {

	s := make([]rune, tkLength)

	length := big.NewInt(int64(len(letters)))

	for i := range s {

		number, _ := rand.Int(rand.Reader, length)
		s[i] = letters[number.Int64()]
	}

	return string(s)
}

var numbers = []rune("0123456789")

// Generate a random string of numbers with the specified length
func generateNumbers(tkLength int32) string {

	s := make([]rune, tkLength)

	length := big.NewInt(int64(len(numbers)))

	for i := range s {

		number, _ := rand.Int(rand.Reader, length)
		s[i] = numbers[number.Int64()]
	}

	return string(s)
}
