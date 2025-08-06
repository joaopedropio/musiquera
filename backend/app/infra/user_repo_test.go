package infra_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/joaopedropio/musiquera/app/database"
	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	"github.com/joaopedropio/musiquera/app/infra"
	"github.com/joaopedropio/musiquera/app/utils"
)

func TestUserRepo_SaveUser(t *testing.T) {
	// Arrange
	dbName, db := database.MustCreateTestSqliteDatabase()
	defer database.MustDestroySqliteDatabase(dbName, db)

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
	dbName, db := database.MustCreateTestSqliteDatabase()
	defer database.MustDestroySqliteDatabase(dbName, db)

	repo := infra.NewUserRepo(db)
	invite := domain.CreateInvite()
 	err := repo.SaveInvite(invite)
	assert.Nil(t, err)

	inviteDB, err := repo.GetInviteByID(invite.ID())
	assert.NotNil(t, inviteDB)
	assert.Nil(t, err)
	assert.Equal(t, invite.ID(), inviteDB.ID())
	assert.Nil(t,inviteDB.UserID())
	assert.Equal(t, domain.InviteStatusPending, inviteDB.Status())

	userID := uuid.New()
	username := "username"
	name := "User Name"
	password := "password"
	email := "example@mail.com"
	createdAt := time.Now()
	user := domain.NewUser(userID, email, username, name, password, createdAt)

	err = repo.AddUser(user)
	assert.Nil(t, err)

	code, err := inviteDB.Accept(userID)
	assert.Nil(t, err)
	fmt.Printf("code: %s\n", code)

	err = repo.SaveInvite(inviteDB)
	assert.Nil(t, err)

	inviteDB, err = repo.GetInviteByID(invite.ID())
	assert.Nil(t, err)
	assert.NotNil(t, inviteDB)
	assert.Equal(t, invite.ID(), inviteDB.ID())
	assert.Equal(t, userID, *inviteDB.UserID())
	assert.Equal(t, domain.InviteStatusAccepted, inviteDB.Status())
}
