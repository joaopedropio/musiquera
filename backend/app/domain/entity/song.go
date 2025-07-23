package domain

import "time"

type Track interface {
	Name() string
	Lyrics() string
	File() string
	Duration() time.Duration
}

type track struct {
	name     string
	lyrics   string
	file     string
	duration time.Duration
}

func (s *track) Name() string {
	return s.name
}

func (s *track) Lyrics() string {
	return s.lyrics
}

func (s *track) File() string {
	return s.file
}

func (s *track) Duration() time.Duration {
	return s.duration
}

func NewTrack(name, lyrics, file string, duration time.Duration) Track {
	return &track{
		name,
		lyrics,
		file,
		duration,
	}
}
