package domain

type Artist interface {
	Name() string
}

type artist struct {
	name string
}

func (a *artist) Name() string {
	return a.name
}

func NewArtist(name string) Artist {
	return &artist{
		name,
	}
}
