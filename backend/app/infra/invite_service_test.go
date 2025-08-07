package infra_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joaopedropio/musiquera/app/database"
	"github.com/joaopedropio/musiquera/app/infra"
)

func TestInviteService_ShouldCreateInvite(t *testing.T) {
	dbName, db := database.MustCreateTestSqliteDatabase()
	defer database.MustDestroySqliteDatabase(dbName, db)

	appURL := "localhost:8080"
	userRepo := infra.NewUserRepo(db)
	service := infra.NewInviteService(appURL, userRepo)
	invite, inviteLink, err := service.CreateInvite()
	assert.NoError(t, err)
	assert.NotNil(t, invite)
	assert.Equal(t, fmt.Sprintf("https://localhost:8080/invite/%s", invite.ID().String()), inviteLink)
}
