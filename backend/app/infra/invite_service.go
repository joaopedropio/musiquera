package infra

import (
	"fmt"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
)

const InviteRouteURLFormat = "https://%s/api/invite/%s"

type InviteService interface {
	CreateInvite() (domain.Invite, string, error)
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
