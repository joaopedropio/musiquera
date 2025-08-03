package infra_test

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/joaopedropio/musiquera/app/domain/entity"
	"github.com/joaopedropio/musiquera/app/infra"
)

type UserRepoMock struct {
	GetUserByUsernameFunc func(username string) (domain.User, error)
	AddUserFunc           func(user domain.User) error
	CreateInviteFunc      func() (uuid.UUID, error)
	SaveInviteFunc func(invite domain.Invite) error
	GetInviteByIDFunc func(id uuid.UUID) (domain.Invite, error)
}

func (r *UserRepoMock) GetInviteByID (id uuid.UUID) (domain.Invite, error) {
	if r.GetInviteByIDFunc == nil {
		panic("GetInviteByIDFunc not implemented")
	}
	return r.GetInviteByIDFunc(id)
}

func (r *UserRepoMock) SaveInvite(invite domain.Invite) error {
	if r.SaveInviteFunc == nil {
		panic("SaveInvite not implemented")
	}
	return r.SaveInviteFunc(invite)
}

func (r *UserRepoMock) GetUserByUsername(username string) (domain.User, error) {
	if r.GetUserByUsernameFunc == nil {
		panic("GetUserByUsername not implemented")
	}
	return r.GetUserByUsernameFunc(username)
}

func (r *UserRepoMock) AddUser(user domain.User) error {
	if r.AddUserFunc == nil {
		panic("AddUser not implemented")
	}
	return r.AddUserFunc(user)
}

func (r *UserRepoMock) CreateInvite() (uuid.UUID, error) {
	if r.CreateInviteFunc == nil {
		panic("CreateInvite not implemented")
	}
	return r.CreateInviteFunc()
}

func TestLoginService_ShouldNotLogin_WhenUserExistsButPasswordDoesNotMatch(t *testing.T) {
	username := "username"
	password := "12345"
	email := "example@mail.com"
	wrongPasswornd := "invalid password"
	passService := infra.NewPasswordService(rand.Text())
	hash, err := passService.HashPassword(password)
	assert.Nil(t, err)

	userRepoMock := &UserRepoMock{}
	userRepoMock.GetUserByUsernameFunc = func(username string) (domain.User, error) {
		return domain.NewUser(uuid.New(), email, username, "User", hash, time.Now()), nil
	}

	loginService := infra.NewLoginService(passService, userRepoMock)
	token, err := loginService.Login(username, wrongPasswornd)
	assert.Empty(t, token)
	assert.NotNil(t, err)
	assert.Equal(t, "password does not match", err.Error())
}

func TestLoginService_ShouldNotLogin_WhenUserDoesNotExists(t *testing.T) {
	username := "username"
	passService := infra.NewPasswordService(rand.Text())

	userRepoMock := &UserRepoMock{}
	userRepoMock.GetUserByUsernameFunc = func(username string) (domain.User, error) {
		return nil, fmt.Errorf("user with username %s not found", username)
	}

	loginService := infra.NewLoginService(passService, userRepoMock)
	token, err := loginService.Login(username, "12345")
	assert.NotNil(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "unable to get user by username: user with username username not found", err.Error())
}

func TestLoginService_ShouldLogin_WhenUserExistsAndPasswordMatches(t *testing.T) {
	username := "username"
	password := "12345"
	email := "example@mail.com"
	passService := infra.NewPasswordService(rand.Text())
	hash, err := passService.HashPassword(password)
	assert.Nil(t, err)

	userRepoMock := &UserRepoMock{}
	userRepoMock.GetUserByUsernameFunc = func(username string) (domain.User, error) {
		return domain.NewUser(uuid.New(), email, username, "User", hash, time.Now()), nil
	}

	loginService := infra.NewLoginService(passService, userRepoMock)
	_, err = loginService.Login(username, password)
	assert.Nil(t, err)
}

func TestLoginService_ShouldBeLogged_WhenTokenIsValidAndContainsUsername(t *testing.T) {
	username := "username"
	password := "12345"
	email := "example@mail.com"
	passService := infra.NewPasswordService(rand.Text())
	hash, err := passService.HashPassword(password)
	assert.Nil(t, err)

	userRepoMock := &UserRepoMock{}
	userRepoMock.GetUserByUsernameFunc = func(username string) (domain.User, error) {
		return domain.NewUser(uuid.New(), email, username, "User", hash, time.Now()), nil
	}

	loginService := infra.NewLoginService(passService, userRepoMock)
	tkn, err := loginService.Login(username, password)
	assert.Nil(t, err)

	isLogged, err := loginService.IsLogged(tkn)
	assert.Nil(t, err)
	assert.True(t, isLogged)
}

func TestLoginService_ShouldNotBeLogged_WhenTokenIsEncodedOnADifferentSecret(t *testing.T) {
	// correct secret => "a-string-secret-at-least-256-bits-long"
	jwtSecret := "wrong secret"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXJuYW1lIn0.WmRTFgvHxI9HIWEu0hWFaKBoi6ssP2eUZoZHZRiR08w"
	passService := infra.NewPasswordService(jwtSecret)
	loginService := infra.NewLoginService(passService, nil)
	isLogged, err := loginService.IsLogged(token)
	assert.NotNil(t, err)
	assert.False(t, isLogged)
	assert.Equal(t, "unable to decode jwt token: could not verify message using any of the signatures or keys", err.Error())
}

func TestLoginService_ShouldNotBeLogged_WhenTokenIsEmpty(t *testing.T) {
	jwtSecret := "secret"
	token := ""
	passService := infra.NewPasswordService(jwtSecret)
	loginService := infra.NewLoginService(passService, nil)
	isLogged, err := loginService.IsLogged(token)

	assert.NotNil(t, err)
	assert.False(t, isLogged)
	assert.Equal(t, "unable to decode jwt token: failed to parse jws: invalid byte sequence", err.Error())
}

func TestLoginService_ShouldNotBeLogged_WhenTokenDoesNotHaveUsernameField(t *testing.T) {
	secret := "a-string-secret-at-least-256-bits-long"
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
	passService := infra.NewPasswordService(secret)
	loginService := infra.NewLoginService(passService, nil)

	isLogged, err := loginService.IsLogged(token)
	assert.NotNil(t, err)
	assert.False(t, isLogged)
	assert.Equal(t, "jwt token (besides valid) does not have username filed", err.Error())
}
