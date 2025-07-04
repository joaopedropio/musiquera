package app

import (
	"fmt"
	"time"

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
	songs := []domain.Song{
		domain.NewSong("Battery", "", "/media/master_of_puppets/Battery__Remastered___uzlOcupu5UE_/manifest.mpd", time.Minute * 5),
		domain.NewSong("Damage Inc.", "", "/media/master_of_puppets/Damage__Inc___Remastered___Abe3AZhcGQs_/manifest.mpd", time.Minute * 5),
		domain.NewSong("Disposable Heroes", "", "/media/master_of_puppets/Disposable_Heroes__Remastered___p3Y8VSVyYN8_/manifest.mpd", time.Minute * 5),
		domain.NewSong("Leper Messiah", "", "/media/master_of_puppets/Leper_Messiah__Remastered___dJp5r4HdRn4_/manifest.mpd", time.Minute * 5),
		domain.NewSong("Master Of Puppets", "", "/media/master_of_puppets/Master_Of_Puppets__Remastered___u6LahTuw02c_/manifest.mpd", time.Minute * 5),
		domain.NewSong("Orion", "", "/media/master_of_puppets/Orion__Remastered___z7bUJPj_8v0_/manifest.mpd", time.Minute * 5),
		domain.NewSong("The Thing That Should Never Be", "", "/media/master_of_puppets/The_Thing_That_Should_Not_Be__Remastered___gm9c_QpuMms_/manifest.mpd", time.Minute * 5),
		domain.NewSong("Welcome Home (Sanitarium)", "", "/media/master_of_puppets/Welcome_Home__Sanitarium___Remastered___G_868UwoJvM_/manifest.mpd", time.Minute * 5),
	}

	id, err := a.repo.AddAlbum(
		"Master of Puppets",
		domain.NewDate(1986, 3, 3),
		domain.NewArtist("Metallica"),
		songs)
	if err != nil {
		return fmt.Errorf("unable to add album: %w", err)
	}
	fmt.Println("albumID " + id.String())
	return nil
}
