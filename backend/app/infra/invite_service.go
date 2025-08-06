package infra

import (
	"fmt"

	"github.com/google/uuid"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
)

const InviteRouteURLFormat = "https://%s/api/invite/%s"

type InviteService interface {
	GetInvite(inviteID uuid.UUID) (domain.Invite, error)
	CreateInvite() (domain.Invite, string, error)
	AcceptInvite(inviteID uuid.UUID, name, username, password, email string, useEmail bool) error
	ConfirmInvite(inviteID uuid.UUID, username, confirmationCode string) error
}

func NewInviteService(appURL string, userRepo UserRepo) InviteService {
	return &inviteService{
		appURL:   appURL,
		userRepo: userRepo,
	}
}

type inviteService struct {
	userRepo UserRepo
	appURL   string
}

func (s *inviteService) GetInvite(inviteID uuid.UUID) (domain.Invite, error) {
	return s.userRepo.GetInviteByID(inviteID)
}

func (s *inviteService) CreateInvite() (domain.Invite, string, error) {
	invite := domain.CreateInvite()
	if err := s.userRepo.SaveInvite(invite); err != nil {
		return nil, "", fmt.Errorf("unable to save invite: %w", err)
	}
	return invite, fmt.Sprintf(InviteRouteURLFormat, s.appURL, invite.ID()), nil
}

func (s *inviteService) AcceptInvite(inviteID uuid.UUID, name, username, password, email string, useEmail bool) error {
	invite, err := s.userRepo.GetInviteByID(inviteID)
	if err != nil {
		return fmt.Errorf("unable to get invite: %w", err)
	}
	if invite.Status() != domain.InviteStatusPending {
		return fmt.Errorf("cant accept invite with status %s", invite.Status())
	}
	if len(username) < 3 {
		return fmt.Errorf("username must be at least 3 characters long: len(username): %d", len(username))
	}
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("unable to check is username exists: %w", err)
	}
	if user.Username() == username {
		return fmt.Errorf("username %s is already in use", username)
	}
	if useEmail {
		user, err := s.userRepo.GetUserByEmail(email)
		if err != nil {
			return fmt.Errorf("unable to get user by email: %w", err)
		}
		if user.Email() == email {
			return fmt.Errorf("email %s is already in use", email)
		}
	}

	if !useEmail {
		email = ""
	}
	newUser, err := domain.CreateUser(name, username, email, password)
	if err != nil {
		return fmt.Errorf("unable to create user: %w", err)
	}
	if err := s.userRepo.AddUser(newUser); err != nil {
		return fmt.Errorf("unable to save new user: %w", err)
	}
	code, err := invite.Accept(newUser.ID())
	if err != nil {
		return fmt.Errorf("unable to accept invite: %w", err)
	}
	if !useEmail {
		if err := invite.Confirm(code); err != nil {
			return fmt.Errorf("unable to confirm invite: %w", err)
		}
	} else {
		fmt.Printf("send email to %s with code %s\n", email, code)
	}
	if err := s.userRepo.SaveInvite(invite); err != nil {
		return fmt.Errorf("unable to save invite: %w", err)
	}
	
	return nil
}

func (s *inviteService) ConfirmInvite(inviteID uuid.UUID, username, confirmationCode string) error {
	invite, err := s.userRepo.GetInviteByID(inviteID)
	if err != nil {
		return fmt.Errorf("unable to get invite by id: %w", err)
	}
	if invite.Status() != domain.InviteStatusAccepted {
		return fmt.Errorf("cant confirm invite with status %s", invite.Status())
	}
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return fmt.Errorf("unable to get user by username: username=%s: %w", username, err)
	}
	if user.ID().String() != inviteID.String() {
		return fmt.Errorf("this invite does not match with this user: username=%s, %w", username, err)
	}
	if err := invite.Confirm(confirmationCode); err != nil {
		return fmt.Errorf("unable to confirm invite: %w", err)
	}
	if err := s.userRepo.SaveInvite(invite); err != nil {
		return fmt.Errorf("unable to save invite: %w", err)
	}
	return nil
}
