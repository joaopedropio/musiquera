package domain

import (
	"time"
)

type Segment interface {
	Name() string
	Position() int64
}

func NewSegment(name string, position int64) Segment {
	return &segment{
		name,
		position,
	}
}

type segment struct {
	name string
	position int64
}

func (s *segment) Name () string {
	return s.name
}

func (s *segment) Position() int64 {
	return s.position
}

type Track interface {
	Name() string
	Lyrics() string
	File() string
	Duration() time.Duration
	Segments() []Segment
}

type track struct {
	name     string
	lyrics   string
	file     string
	duration time.Duration
	segments []Segment
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

func (s *track) Segments() []Segment {
	return s.segments
}

func NewTrack(name, lyrics, file string, duration time.Duration, segments []Segment) Track {
	return &track{
		name,
		lyrics,
		file,
		duration,
		segments,
	}
}
