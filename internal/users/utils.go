package users

import (
	"math/rand"
	"time"
)

const (
	awailableLetters = "1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	keyLength = 32
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generateUserID generates random user ID
func generateUserID() string {
	id := make([]byte, keyLength)
	for idx := range id {
		id[idx] = awailableLetters[rand.Int63()%keyLength]
	}

	return string(id)
}
