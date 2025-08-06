package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joaopedropio/musiquera/app/utils"
)

func TestSecurity_HashPassword(t *testing.T) {
	password := "12345"

	hashed, err := utils.HashPassword(password)
	assert.NoError(t, err)
	verified, err := utils.VerifyPassword(password, hashed)
	assert.NoError(t, err)
	assert.True(t, verified)
}
