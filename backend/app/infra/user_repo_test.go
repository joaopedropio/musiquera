package infra_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	"github.com/joaopedropio/musiquera/app/infra"
	"github.com/joaopedropio/musiquera/app/utils"
)

func TestUserRepo_SaveUser(t *testing.T) {
	// Arrange
	dbName, db := utils.MustCreateTestSqliteDatabase()
	defer utils.MustDestroySqliteDatabase(dbName, db)

	repo := infra.NewUserRepo(db)
	id := uuid.New()
	username := "username"
	name := "User Name"
	password := "password"
	email := "example@mail.com"
	createdAt := time.Now()
	user := domain.NewUser(id, email, username, name, password, createdAt)

	// Act
	err := repo.AddUser(user)
	assert.Nil(t, err)
	dbUser, err := repo.GetUserByUsername(user.Username())
	assert.Nil(t, err)

	// Assert
	assert.NotNil(t, dbUser)
	assert.Equal(t, id, dbUser.ID())
	assert.Equal(t, password, dbUser.Password())
	assert.Equal(t, username, dbUser.Username())
	assert.Equal(t, name, dbUser.Name())
	assert.Equal(t, email, dbUser.Email())
	assert.True(t, utils.IsTimeEqual(createdAt, dbUser.CreatedAt()))
}

func TestUserRepo_SaveInvite(t *testing.T) {
	dbName, db := utils.MustCreateTestSqliteDatabase()
	defer utils.MustDestroySqliteDatabase(dbName, db)

	repo := infra.NewUserRepo(db)
	inviteID, err := repo.CreateInvite()
	assert.Nil(t, err)

	invite, err := repo.GetInviteByID(inviteID)
	assert.NotNil(t, invite)
	assert.Nil(t, err)
	assert.Equal(t, inviteID, invite.ID())
	assert.Equal(t, uuid.Nil, invite.UserID())
	assert.Equal(t, domain.InviteStatusPending, invite.Status())

	userID := uuid.New()
	username := "username"
	name := "User Name"
	password := "password"
	email := "example@mail.com"
	createdAt := time.Now()
	user := domain.NewUser(userID, email, username, name, password, createdAt)

	err = repo.AddUser(user)
	assert.Nil(t, err)

	err = invite.SetUserID(userID)
	assert.Nil(t, err)

	err = repo.SaveInvite(invite)
	assert.Nil(t, err)

	invite, err = repo.GetInviteByID(inviteID)
	assert.Nil(t, err)
	assert.NotNil(t, invite)
	assert.Equal(t, inviteID, invite.ID())
	assert.Equal(t, userID, invite.UserID())
	assert.Equal(t, domain.InviteStatusAccepted, invite.Status())
}
