package domain

import (
	"time"

	"github.com/google/uuid"
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
