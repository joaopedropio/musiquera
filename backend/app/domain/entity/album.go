package domain

type Album interface {
	Name() string
	ReleaseDate() Date
}

type album struct {
	name        string
	releaseDate Date
}

func (a *album) Name() string {
	return a.name
}

func (a *album) ReleaseDate() Date {
	return a.releaseDate
}

func NewAlbum(name string, releaseDate Date) Album {
	return &album{
		name,
		releaseDate,
	}
}
