package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

const InviteStatusPending = InviteStatus("pending")
const InviteStatusAccepted = InviteStatus("accepted")


type InviteStatus string

type Invite interface {
	ID() uuid.UUID
	UserID() uuid.UUID
	Status() InviteStatus
	CreatedAt() time.Time

	SetUserID(userID uuid.UUID) error
}

func NewInvite(id, userID uuid.UUID, status InviteStatus, createdAt time.Time) Invite {
	return &invite{
		id,
		userID,
		status,
		createdAt,
	}
}

type invite struct {
	id        uuid.UUID
	userID    uuid.UUID
	status    InviteStatus
	createdAt time.Time
}

func (i *invite) ID() uuid.UUID {
	return i.id
}

func (i *invite) UserID() uuid.UUID {
	return i.userID
}

func (i *invite) Status() InviteStatus {
	return i.status
}

func (i *invite) CreatedAt() time.Time {
	return i.createdAt
}

func (i *invite) SetUserID(userID uuid.UUID) error {
	if i.status != InviteStatusPending {
		return errors.New("can't set user id to a non pending invite")
	}
	i.userID = userID
	i.status = InviteStatusAccepted
	return nil
}
