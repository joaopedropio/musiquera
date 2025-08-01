package infra_test

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joaopedropio/musiquera/app/infra"
)

func TestPasswordService_Hash(t *testing.T) {
	password := "12345"

	service := infra.NewPasswordService(rand.Text())
	hashed, err :=service.HashPassword(password)
	assert.Nil(t, err)
	verified, err := service.VerifyPassword(password, hashed)
	assert.Nil(t, err)
	assert.True(t, verified)
}
