package infra

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	_ "modernc.org/sqlite"

	"github.com/joaopedropio/musiquera/app/database"
	entity "github.com/joaopedropio/musiquera/app/domain/entity"
)

type UserDB struct {
	IDField        uuid.UUID `db:"id"`
	EmailField     string    `db:"email"`
	UsernameField  string    `db:"username"`
	NameField      string    `db:"name"`
	PasswordField  string    `db:"password"`
	CreatedAtField time.Time `db:"created_at"`
}

func (u *UserDB) ID() uuid.UUID {
	return u.IDField
}

func (u *UserDB) Email() string {
	return u.EmailField
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

type InviteDB struct {
	IDField        uuid.UUID           `db:"id"`
	UserIDField    database.NullUUID           `db:"user_id"`
	StatusField    entity.InviteStatus `db:"status"`
	ConfirmationCodeField string `db:"code"`
	CreatedAtField time.Time           `db:"created_at"`
}

func (i *InviteDB) ID() uuid.UUID {
	return i.IDField
}

func (i *InviteDB) UserID() *uuid.UUID {
	return i.UserIDField.Ptr()
}

func (i *InviteDB) Status() entity.InviteStatus {
	return i.StatusField
}

func (i *InviteDB) ConfirmationCode() string {
	return i.ConfirmationCodeField
}

func (i *InviteDB) CreatedAt() time.Time {
	return i.CreatedAtField
}

func CreateInviteDB(invite entity.Invite) *InviteDB {
	return &InviteDB{
		IDField:        invite.ID(),
		UserIDField:    database.NewNullUUID(invite.UserID()),
		StatusField:    invite.Status(),
		ConfirmationCodeField: invite.ConfirmationCode(),
		CreatedAtField: invite.CreatedAt(),
	}
}

func CreateInviteFromInviteDB(inviteDB *InviteDB) entity.Invite {
	return entity.NewInvite(
		inviteDB.IDField,
		inviteDB.UserIDField.Ptr(),
		inviteDB.StatusField,
		inviteDB.ConfirmationCodeField,
		inviteDB.CreatedAtField)

}

type UserRepo interface {
	GetUserByUsername(username string) (entity.User, error)
	AddUser(user entity.User) error
	SaveInvite(invite entity.Invite) error
	GetInviteByID(id uuid.UUID) (entity.Invite, error)
}

func CreateUserDB(user entity.User) *UserDB {
	return &UserDB{
		IDField:        user.ID(),
		EmailField:     user.Email(),
		UsernameField:  user.Username(),
		NameField:      user.Name(),
		PasswordField:  user.Password(),
		CreatedAtField: user.CreatedAt(),
	}
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) AddUser(user entity.User) error {
	query := `
	INSERT INTO users (id, name, email, username, password, created_at)
	VALUES (:id, :name, :email, :username, :password, :created_at);
	`
	dbUser := CreateUserDB(user)
	_, err := r.db.NamedExec(query, dbUser)
	if err != nil {
		return fmt.Errorf("unable to insert user %s: %w", user.Username(), err)
	}
	return nil
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
}

func (r *userRepo) GetInviteByID(id uuid.UUID) (entity.Invite, error) {
	query := `
	SELECT id, user_id, code, status, created_at FROM invites
	WHERE id = ?;
	`
	var dbInvite InviteDB
	err := r.db.Get(&dbInvite, query, id)
	if err != nil {
		return nil, fmt.Errorf("unable to get invite from db: %w", err)
	}
	return CreateInviteFromInviteDB(&dbInvite), nil
}

func (r *userRepo) SaveInvite(invite entity.Invite) error {
	exists, err := r.inviteExists(invite)
	if err != nil {
		return fmt.Errorf("unable to check if invite exists: %w", err)
	}
	if exists {
		return r.updateInvite(invite)
	}

	return r.createInvite(invite)
}

func (r *userRepo) inviteExists(invite entity.Invite) (bool, error) {
	query := `
	SELECT EXISTS(SELECT 1 from invites WHERE id = ? LIMIT 1);
	`
	var exists bool
	err := r.db.Get(&exists, query, invite.ID().String())
	if err != nil {
		return false, fmt.Errorf("unable to check if invite exists: %w", err)
	}
	return exists, nil
}

func (r *userRepo) createInvite(invite entity.Invite) error {
	query := `
	INSERT INTO invites (id, user_id, code, status, created_at)
	VALUES (:id, :user_id, :code, :status, :created_at)
	`
	inviteDB := CreateInviteDB(invite)
	_, err := r.db.NamedExec(query, inviteDB)
	if err != nil {
		return fmt.Errorf("unable to insert new invite: %w", err)
	}

	return nil
}

func (r *userRepo) updateInvite(invite entity.Invite) error {
	query := `
	UPDATE invites SET status = :status, user_id = :user_id, code = :code
	WHERE id = :id;
	`

	dbInvite := CreateInviteDB(invite)
	_, err := r.db.NamedExec(query, dbInvite)
	if err != nil {
		return fmt.Errorf("unable to save invite: %w", err)
	}
	return err
}
