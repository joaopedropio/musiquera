package domain

type Artist interface {
	Name() string
	ProfileCoverPhotoPath() string
}

type artist struct {
	name                  string
	profileCoverPhotoPath string
}

func (a *artist) Name() string {
	return a.name
}

func (a *artist) ProfileCoverPhotoPath() string {
	return a.profileCoverPhotoPath
}

func NewArtist(name string, profilePhoto string) Artist {
	return &artist{
		name,
		profilePhoto,
	}
}
