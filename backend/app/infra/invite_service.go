package infra

import (
	"fmt"

	"github.com/google/uuid"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
)

const InviteRouteURLFormat = "https://%s/api/invite/%s"

type InviteService interface {
	CreateInvite() (domain.Invite, string, error)
	AcceptInvite(inviteID uuid.UUID, name, username, password, email string, useEmail bool) error
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

func (s *inviteService) CreateInvite() (domain.Invite, string, error) {
	invite := domain.CreateInvite()
	if err := s.userRepo.SaveInvite(invite); err != nil {
		return nil, "", fmt.Errorf("unable to save invite: %w", err)
	}
	return invite, fmt.Sprintf(InviteRouteURLFormat, s.appURL, invite.ID()), nil
}

func (s *inviteService) AcceptInvite(inviteID uuid.UUID, name, username, password, email string, useEmail bool) error {
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
	return nil
}
