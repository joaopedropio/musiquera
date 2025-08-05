package domain

import (
	"time"

	"github.com/google/uuid"
)

type Segment interface {
	TrackID() uuid.UUID
	Name() string
	Position() int64
}

func NewSegment(trackID uuid.UUID, name string, position int64) Segment {
	return &segment{
		trackID,
		name,
		position,
	}
}

type segment struct {
	trackID  uuid.UUID
	name     string
	position int64
}

func (s *segment) TrackID() uuid.UUID {
	return s.trackID
}

func (s *segment) Name() string {
	return s.name
}

func (s *segment) Position() int64 {
	return s.position
}

type Track interface {
	ID() uuid.UUID
	Name() string
	Lyrics() string
	File() string
	Duration() time.Duration
	Segments() []Segment
	CreatedAt() time.Time
}

type track struct {
	id        uuid.UUID
	name      string
	lyrics    string
	file      string
	duration  time.Duration
	segments  []Segment
	createdAt time.Time
}

func (s *track) ID() uuid.UUID {
	return s.id
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

func (s *track) CreatedAt() time.Time {
	return s.createdAt
}

func NewTrack(id uuid.UUID, name, lyrics, file string, duration time.Duration, segments []Segment, createdAt time.Time) Track {
	return &track{
		id,
		name,
		lyrics,
		file,
		duration,
		segments,
		createdAt,
	}
}
