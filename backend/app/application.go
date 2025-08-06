package app

import (
	"fmt"
	"log"

	"github.com/joaopedropio/musiquera/app/database"
	domainrepo "github.com/joaopedropio/musiquera/app/domain/repo"
	infra "github.com/joaopedropio/musiquera/app/infra"

	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type Application interface {
	DBConnection() *sqlx.DB
	Close() error
	LoginService() infra.LoginService
	Repo() domainrepo.Repo
	Environment() Environment
	UserRepo() infra.UserRepo
	InviteService() infra.InviteService
}

type application struct {
	db            *sqlx.DB
	repo          domainrepo.Repo
	env           Environment
	userRepo      infra.UserRepo
	loginService  infra.LoginService
	inviteService infra.InviteService
}

func (a *application) Environment() Environment {
	return a.env
}

func (a *application) Repo() domainrepo.Repo {
	return a.repo
}

func (a *application) UserRepo() infra.UserRepo {
	return a.userRepo
}

func (a *application) LoginService() infra.LoginService {
	return a.loginService
}

func (a *application) InviteService() infra.InviteService {
	return a.inviteService
}

func (a *application) DBConnection() *sqlx.DB {
	return a.db
}

func (a *application) Close() error {
	fmt.Println("closing db connection")
	if err := a.db.Close(); err != nil {
		log.Fatalf("unable to close db connection: %s", err)
	}
	return nil
}

func NewApplication() (Application, error) {
	env := GetEnvironmentVariables()
	db, err := sqlx.Open("sqlite", env.DatabaseDir+"/musiquera.db?_foreign_keys=on")
	if err != nil {
		panic(fmt.Errorf("unable to start db connection: %w", err))
	}
	repo := infra.NewRepo(db)
	userRepo := infra.NewUserRepo(db)
	loginService := infra.NewLoginService(env.JWTSecret, userRepo)
	inviteService := infra.NewInviteService(env.AppURL, userRepo)
	a := &application{
		db,
		repo,
		env,
		userRepo,
		loginService,
		inviteService,
	}
	a.schema(db)
	return a, nil
}

func (a *application) schema(db *sqlx.DB) {
	db.MustExec(database.DatabaseSchema())
}
