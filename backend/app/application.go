package app

import (
	"fmt"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	domainrepo "github.com/joaopedropio/musiquera/app/domain/repo"
	infra "github.com/joaopedropio/musiquera/app/infra"
)

type Application interface {
	Repo() domainrepo.Repo
	Environment() Environment
}

type application struct {
	repo domainrepo.Repo
	env Environment
}

func (a *application) Environment() Environment {
	return a.env
}

func (a *application) Repo() domainrepo.Repo {
	return a.repo
}

func NewApplication() (Application, error) {
	repo := infra.NewRepo()
	env := GetEnvironmentVariables()
	a := &application{
		repo,
		env,
	}
	if err := a.feed(); err != nil {
		return nil, fmt.Errorf("unable to feed: %w", err)
	}
	return a, nil
}

func (a *application) feed() error {
	id, err := a.repo.AddAlbum(
		"Master of Puppets",
		domain.NewDate(1986, 3, 3),
		domain.NewArtist("Metallica"),
		nil)
	if err != nil {
		return fmt.Errorf("unable to add album: %w", err)
	}
	fmt.Println("albumID " + id.String())
	return nil
}
