package domain_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/joaopedropio/musiquera/app/database"
	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	"github.com/joaopedropio/musiquera/app/infra"
)

func TestInvite_WhenCreatingInvite_UserShouldBeEmptyAndStatusPending(t *testing.T) {
	dbName, db := database.MustCreateTestSqliteDatabase()
	defer database.MustDestroySqliteDatabase(dbName, db)

	repo := infra.NewUserRepo(db)
	invite := domain.CreateInvite()
	assert.Nil(t, invite.UserID())
	assert.Equal(t, domain.InviteStatusPending, invite.Status())
	assert.NotNil(t, invite)
	err := repo.SaveInvite(invite)
	assert.NoError(t, err)

	inviteDB, err := repo.GetInviteByID(invite.ID())
	assert.NoError(t, err)
	assert.NotNil(t, inviteDB)
	assert.Nil(t, inviteDB.UserID())
	assert.Equal(t, domain.InviteStatusPending, inviteDB.Status())
}

func TestInvite_WhenUserConfirmInvite_WhenInviteIsConfirmed(t *testing.T) {
	dbName, db := database.MustCreateTestSqliteDatabase()
	defer database.MustDestroySqliteDatabase(dbName, db)

	repo := infra.NewUserRepo(db)
	invite := domain.CreateInvite()
	user := createUser(t, repo)
	code, err := invite.Accept(user.ID())
	assert.NoError(t, err)
	err = invite.Confirm(code)
	assert.NoError(t, err)
	err = repo.SaveInvite(invite)
	assert.NoError(t, err)

	inviteDB, err := repo.GetInviteByID(invite.ID())
	assert.NoError(t, err)
	assert.Equal(t, invite.ID(), inviteDB.ID())
	assert.Equal(t, invite.UserID(), inviteDB.UserID())
}

func TestInvite_WhenUserAcceptInvite_ShouldSetUserIDAndConfirmationCode(t *testing.T) {
	dbName, db := database.MustCreateTestSqliteDatabase()
	defer database.MustDestroySqliteDatabase(dbName, db)

	repo := infra.NewUserRepo(db)
	invite := domain.CreateInvite()
	user := createUser(t, repo)
	code, err := invite.Accept(user.ID())
	assert.NoError(t, err)
	assert.Len(t, code, 6)
	assert.Equal(t, invite.ConfirmationCode(), code)
	assert.Equal(t, user.ID().String(),  invite.UserID().String())
	err = repo.SaveInvite(invite)
	assert.NoError(t, err)

	inviteDB, err := repo.GetInviteByID(invite.ID())
	assert.NoError(t, err)
	assert.Equal(t, invite.ID(), inviteDB.ID())
	assert.Equal(t, invite.UserID(), inviteDB.UserID())
}

func createUser(t *testing.T, repo infra.UserRepo) domain.User {
	hashedPassword, err := infra.NewPasswordService("").HashPassword("12345")
	assert.NoError(t, err)
	user := domain.NewUser(uuid.New(), "example@mail.com", "username", "User name", hashedPassword, time.Now())
	err = repo.AddUser(user)
	assert.NoError(t, err)
	return user
}
