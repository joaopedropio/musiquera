package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/joaopedropio/musiquera/app/utils"
)

const InviteStatusPending = InviteStatus("pending")
const InviteStatusAccepted = InviteStatus("accepted")
const InviteStatusConfirmed = InviteStatus("confirmed")

type InviteStatus string

type Invite interface {
	ID() uuid.UUID
	UserID() *uuid.UUID
	Status() InviteStatus
	ConfirmationCode() string
	CreatedAt() time.Time
	Accept(userID uuid.UUID) (string, error)
	Confirm(confirmationCode string) error
}

func CreateInvite() Invite {
	return NewInvite(uuid.New(), nil, InviteStatusPending, "", time.Now())
}

func NewInvite(id uuid.UUID, userID *uuid.UUID, status InviteStatus, confirmationCode string, createdAt time.Time) Invite {
	return &invite{
		id,
		userID,
		status,
		confirmationCode,
		createdAt,
	}
}

type invite struct {
	id               uuid.UUID
	userID           *uuid.UUID
	status           InviteStatus
	confirmationCode string
	createdAt        time.Time
}

func (i *invite) ID() uuid.UUID {
	return i.id
}

func (i *invite) UserID() *uuid.UUID {
	return i.userID
}

func (i *invite) Status() InviteStatus {
	return i.status
}

func (i *invite) ConfirmationCode() string {
	return i.confirmationCode
}
func (i *invite) CreatedAt() time.Time {
	return i.createdAt
}

func (i *invite) Accept(userID uuid.UUID) (string, error) {
	if i.status != InviteStatusPending {
		return "", errors.New("can't set user id to a non pending invite")
	}
	code, err := utils.GenerateMFACode()
	if err != nil {
		return "", fmt.Errorf("unable to create mfa code: %w", err)
	}
	i.userID = &userID
	i.status = InviteStatusAccepted
	i.confirmationCode = code

	return code, nil
}

func (i *invite) Confirm(confirmationCode string) error {
	if confirmationCode != i.confirmationCode {
		return fmt.Errorf("confirmation code does not match: invite code: %s, code: %s", i.confirmationCode, confirmationCode)
	}
	i.status = InviteStatusConfirmed
	return nil
}
