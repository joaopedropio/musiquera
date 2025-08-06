package utils

import (
	"crypto/rand"
	"fmt"
)

func GenerateMFACode() (string, error) {
	var max uint32 = 1000000 // 6 digits
	var num uint32

	// Read 4 random bytes
	b := make([]byte, 4)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	// Convert bytes to uint32
	num = (uint32(b[0]) << 24) | (uint32(b[1]) << 16) | (uint32(b[2]) << 8) | uint32(b[3])
	code := num % max

	return fmt.Sprintf("%06d", code), nil
}
