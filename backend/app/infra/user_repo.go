package infra

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	entity "github.com/joaopedropio/musiquera/app/domain/entity"

	_ "github.com/mattn/go-sqlite3"
)

type UserDB struct {
	IDField uuid.UUID `db:"id"`
	UsernameField string `db:"username"`
	NameField     string `db:"name"`
	PasswordField string `db:"password"`
	CreatedAtField time.Time `db:"createdAt"`
}

func (u *UserDB) ID() uuid.UUID {
	return u.IDField
}

func (u *UserDB) Username() string {
	return u.UsernameField
}

func (u *UserDB) Name() string {
	return u.NameField
}

func (u *UserDB) Password() string {
	return u.PasswordField
}

func (u *UserDB) CreatedAt() time.Time {
	return u.CreatedAtField
}

type UserRepo interface {
	GetUserByUsername(username string) (entity.User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) GetUserByUsername(username string) (entity.User, error) {
	var users []*UserDB
	if err := r.db.Select(&users, "SELECT * FROM users WHERE username = ?", username); err != nil {
		return nil, fmt.Errorf("unable to select user by username: %w", err)
	}
	if len(users) > 1 {
		return nil, fmt.Errorf("should only have 1 user with username %s but have %d", username, len(users))
	}

	if len(users) == 0 {
		return nil, fmt.Errorf("user %s not found", username)
	}
	return users[0], nil

	//savedPassword := os.Getenv("SAVED_PASSWORD")
	//user := entity.NewUser(uuid.New(), "pio", "Joao Pedro", savedPassword, time.Now())
	//return user, nil
}
