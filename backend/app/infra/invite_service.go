package infra

import (
	"fmt"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
)

const InviteRouteURLFormat = "https://%s/api/invite/%s"

type InviteService interface {
	CreateInvite() (string, error)
}

type inviteService struct {
	userRepo UserRepo
	appURL   string
}

func (s *inviteService) CreateInvite() (string, error) {
	invite := domain.CreateInvite()
	if err := s.userRepo.SaveInvite(invite); err != nil {
		return "", fmt.Errorf("unable to save invite: %w", err)
	}
	return fmt.Sprintf(InviteRouteURLFormat, s.appURL, invite.ID()), nil
}
