package infra

import (
	"os"
	"time"

	"github.com/google/uuid"

	entity "github.com/joaopedropio/musiquera/app/domain/entity"
)

type UserRepo interface {
	GetUserByUsername(username string) (entity.User, error)
}

type userRepo struct {
}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (r *userRepo) GetUserByUsername(username string) (entity.User, error) {
	savedPassword := os.Getenv("SAVED_PASSWORD")
	user := entity.NewUser(uuid.New(), "pio", "Joao Pedro", savedPassword, time.Now())
	return user, nil
}
