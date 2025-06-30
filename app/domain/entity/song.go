package domain

import "time"

type Song interface {
	Name() string
	Lyrics() string
	File() string
	Duration() time.Duration
}

type song struct {
	name     string
	lyrics   string
	file     string
	duration time.Duration
}

func (s *song) Name() string {
	return s.name
}

func (s *song) Lyrics() string {
	return s.lyrics
}

func (s *song) File() string {
	return s.file
}

func (s *song) Duration() time.Duration {
	return s.duration
}

func NewSong(name, lyrics, file string, duration time.Duration) Song {
	return &song{
		name,
		lyrics,
		file,
		duration,
	}
}
