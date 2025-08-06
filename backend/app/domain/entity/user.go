package domain

import (
	"errors"
	"fmt"
	"net/mail"
	"regexp"
	"time"

	"github.com/google/uuid"

	"github.com/joaopedropio/musiquera/app/utils"
)

type User interface {
	ID() uuid.UUID
	Email() string
	Username() string
	Name() string
	Password() string
	CreatedAt() time.Time
}

type user struct {
	id uuid.UUID
	email string
	username string
	name     string
	password string
	createdAt time.Time
}

func CreateUser(name, username, email, password string) (User, error) {
	if !utils.IsValidPassword(password) {
		return nil, errors.New("invalid password")
	}
	if !isValidUsername(username) {
		return nil, errors.New("invalid username")
	}
	if email != "" {
		_, err :=mail.ParseAddress(email)
		if err != nil {
			return nil, fmt.Errorf("unable to parse email '%s': %w", email, err)
		}
	}
	return NewUser(
		uuid.New(), 
		email,
		username,
		name,
		password,
		time.Now(),
	), nil
}

func isValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_-]{3,}$`)
	return re.MatchString(username)
}

func NewUser(id uuid.UUID, email, username, name, password string, createdAt time.Time) User {
	return &user{
		id: id,
		email: email,
		username: username,
		name:     name,
		password: password,
		createdAt: createdAt,
	}
}

func (u *user) ID() uuid.UUID {
	return u.id
}

func (u *user) Email() string {
	return u.email
}

func (u *user) Username() string {
	return u.username
}

func (u *user) Name() string {
	return u.name
}

func (u *user) Password() string {
	return u.password
}

func (u *user) CreatedAt() time.Time {
	return u.createdAt
}
