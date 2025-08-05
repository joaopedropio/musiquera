package infra

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	"github.com/joaopedropio/musiquera/app/domain/entity"
	domainrepo "github.com/joaopedropio/musiquera/app/domain/repo"
	"github.com/joaopedropio/musiquera/app/utils"
)

func NewRepo(db *sqlx.DB) domainrepo.Repo {
	return &repo{
		db: db,
	}
}

type repo struct {
	db       *sqlx.DB
}

func CreateFullRelease(fullReleaseDB FullReleaseDB) domain.FullRelease {
	var tracks []domain.Track
	for _, trackDB := range fullReleaseDB.TracksField {
		var segments []domain.Segment
		for _, segmentDB := range trackDB.SegmentsField {
			segments = append(segments, domain.NewSegment(segmentDB.TrackIDField, segmentDB.NameField, segmentDB.PositionField))
		}
		tracks = append(tracks, domain.NewTrack(trackDB.IDField, trackDB.NameField, trackDB.LyricsField, trackDB.FileField, trackDB.DurationField, segments, trackDB.CreatedAtField))
	}
	return domain.NewFullRelease(
		fullReleaseDB.IDField,
		fullReleaseDB.NameField,
		fullReleaseDB.ReleaseTypeField,
		fullReleaseDB.CoverField,
		fullReleaseDB.ReleaseDateField.Date(),
		fullReleaseDB.Artist,
		tracks,
		fullReleaseDB.CreatedAtField,
	)
}

func CreateFullReleaseDB(fullRelease domain.FullRelease) *FullReleaseDB {
	var tracksDB []*TrackDB
	for _, track := range fullRelease.Tracks() {
		var segmentsDB []*SegmentDB
		for _, segment := range track.Segments() {
			segmentsDB = append(segmentsDB, &SegmentDB{
				segment.TrackID(),
				segment.Name(),
				segment.Position(),
			})
		}
		tracksDB = append(tracksDB, &TrackDB{
			track.ID(),
			fullRelease.ID(),
			track.Name(),
			track.Lyrics(),
			track.File(),
			track.Duration(),
			segmentsDB,
			track.CreatedAt(),
		})
	}

	return &FullReleaseDB{
		fullRelease.ID(),
		fullRelease.Name(),
		fullRelease.Type(),
		fullRelease.Cover(),
		utils.NewDateDB(fullRelease.ReleaseDate()),
		fullRelease.Artist().ID(),
		CreateArtistDB(fullRelease.Artist()),
		tracksDB,
		fullRelease.CreatedAt(),
	}
}

type SegmentDB struct {
	TrackIDField uuid.UUID `db:"track_id"`
	NameField string `db:"name"`
	PositionField int64 `db:"position"`
}

type TrackDB struct {
	IDField uuid.UUID `db:"id"`
	ReleaseIDField uuid.UUID `db:"release_id"`
	NameField string `db:"name"`
	LyricsField string `db:"lyrics"`
	FileField string `db:"file"`
	DurationField time.Duration `db:"duration"`
	SegmentsField []*SegmentDB
	CreatedAtField time.Time `db:"created_at"`
}

type FullReleaseDB struct {
	IDField        uuid.UUID `db:"id"`
	NameField      string `db:"name"`
	ReleaseTypeField domain.ReleaseType `db:"type"`
	CoverField     string `db:"cover"`
	ReleaseDateField   *utils.DateDB `db:"release_date"`
	ArtistIDField    uuid.UUID `db:"artist_id"`
	Artist *ArtistDB
	TracksField    []*TrackDB 
	CreatedAtField time.Time `db:"created_at"`
}


type ArtistDB struct {
	IDField                    uuid.UUID `db:"id"`
	NameField                  string    `db:"name"`
	ProfileCoverPhotoPathField string    `db:"cover"`
	CreatedAtField             time.Time `db:"created_at"`
}

func (a *ArtistDB) ID() uuid.UUID {
	return a.IDField
}

func (a *ArtistDB) Name() string {
	return a.NameField
}

func (a *ArtistDB) ProfileCoverPhotoPath() string {
	return a.ProfileCoverPhotoPathField
}

func (a *ArtistDB) CreatedAt() time.Time {
	return a.CreatedAtField
}

func CreateArtist(a *ArtistDB) domain.Artist {
	return domain.NewArtist(a.IDField, a.NameField, a.ProfileCoverPhotoPathField, a.CreatedAtField)
}

func CreateArtistDB(a domain.Artist) *ArtistDB {
	return &ArtistDB{
		IDField:                    a.ID(),
		NameField:                  a.Name(),
		ProfileCoverPhotoPathField: a.ProfileCoverPhotoPath(),
		CreatedAtField:             a.CreatedAt(),
	}
}

func (r *repo) AddArtist(artist domain.Artist) error {
	artistDB := CreateArtistDB(artist)
	query := `
	INSERT INTO artists (id, name, cover, created_at)
	VALUES (:id, :name, :cover, :created_at);
	`
	_, err := r.db.NamedExec(query, artistDB)
	if err != nil {
		return fmt.Errorf("unable to insert artist: %w", err)
	}
	return nil
}

func (r *repo) GetArtists() ([]domain.Artist, error) {
	var artistsDB []*ArtistDB
	query := `
	SELECT * FROM artists;
	`
	err := r.db.Select(&artistsDB, query)
	if err != nil {
		return nil, fmt.Errorf("unable to select all artists: %w", err)
	}
	var artists []domain.Artist
	for _, artist := range artistsDB {
		artists = append(artists, CreateArtist(artist))
	}
	return artists, nil
}

func (r *repo) GetReleasesByArtist(artistName string) ([]domain.FullRelease, error) {
	var fullReleases []*FullReleaseDB

	query := `
	SELECT 
		r.id AS id,
		r.name AS name,
		r.type AS type,
		r.cover AS cover,
		r.release_date AS release_date,
		r.artist_id AS artist_id,
		r.created_at AS created_at,
		a.id AS "artist.id",
		a.name AS "artist.name",
		a.cover AS "artist.cover",
		a.created_at AS "artist.created_at"
	FROM releases r
	JOIN artists a ON r.artist_id = a.id
	WHERE a.name = ?
	`

	if err := r.db.Select(&fullReleases, query, artistName); err != nil {
		return nil, fmt.Errorf("querying releases by artist name: %w", err)
	}

	for _, release := range fullReleases {
		var tracks []*TrackDB

		trackQuery := `
		SELECT 
			id AS id,
			release_id AS release_id,
			name AS name,
			lyrics AS lyrics,
			file AS file,
			duration AS duration,
			created_at AS created_at
		FROM tracks
		WHERE release_id = ?
		`
		if err := r.db.Select(&tracks, trackQuery, release.IDField); err != nil {
			return nil, fmt.Errorf("querying tracks for release %s: %w", release.IDField, err)
		}

		for _, track := range tracks {
			var segments []*SegmentDB
			segmentQuery := `
			SELECT 
				track_id AS track_id,
				name AS name,
				position AS position
			FROM segments
			WHERE track_id = ?
			`
			if err := r.db.Select(&segments, segmentQuery, track.IDField); err != nil {
				return nil, fmt.Errorf("querying segments for track %s: %w", track.IDField, err)
			}
			track.SegmentsField = segments
		}

		release.TracksField = tracks
	}

	var result []domain.FullRelease
	for _, releaseDB := range fullReleases {
		result = append(result, CreateFullRelease(*releaseDB))
	}

	return result, nil
}

func (r *repo) AddFullRelease(fullRelease domain.FullRelease) error {
	var err error
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("unable to begin transaction: %w", err)
	}
	defer utils.CommitOrRollback(tx, &err)

	releaseDB := CreateFullReleaseDB(fullRelease)

	releaseQuery := `
	INSERT INTO releases (id, name, type, cover, release_date, artist_id, created_at)
	VALUES (:id, :name, :type, :cover, :release_date, :artist_id, :created_at);
	`
	_, err = tx.NamedExec(releaseQuery, releaseDB)
	if err != nil {
		return fmt.Errorf("unable to insert full release: %w", err)
	}

	trackQuery := `
	INSERT INTO tracks (id, release_id, name, lyrics, file, duration, created_at)
	VALUES (:id, :release_id, :name, :lyrics, :file, :duration, :created_at);
	`
	segmentQuery := `
	INSERT INTO segments (track_id, name, position)
	VALUES (:track_id, :name, :position);
	`
	for _, trackDB := range releaseDB.TracksField {
		_, err := tx.NamedExec(trackQuery, trackDB)
		if err!= nil {
			return fmt.Errorf("unable to insert track: %w", err)
		}

		for _, segmentDB := range trackDB.SegmentsField {
			_, err := tx.NamedExec(segmentQuery, segmentDB)
			if err != nil {
				return fmt.Errorf("unable to insert segment: %w", err)
			}
		}
	}
	return nil
}

func (r *repo) GetRelease(id uuid.UUID) (domain.Release, error) {
	fullRelease, err := r.GetFullRelease(id)
	if err != nil {
		return nil, fmt.Errorf("unable to get full release: %w", err)
	}
	return domain.NewRelease(
		fullRelease.Name(),
		fullRelease.Type(),
		fullRelease.ReleaseDate(),
		fullRelease.Cover(),
		fullRelease.Artist(),
	), nil
}

func (r *repo) GetFullRelease(id uuid.UUID) (domain.FullRelease, error) {
	var release FullReleaseDB

	query := `
	SELECT 
		r.id AS id,
		r.name AS name,
		r.type AS type,
		r.cover AS cover,
		r.release_date AS release_date,
		r.artist_id AS artist_id,
		r.created_at AS created_at,
		a.id AS "artist.id",
		a.name AS "artist.name",
		a.cover AS "artist.cover",
		a.created_at AS "artist.created_at"
	FROM releases r
	JOIN artists a ON r.artist_id = a.id
	WHERE r.id = ?
	`

	if err := r.db.Get(&release, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("release not found: %w", err)
		}
		return nil, fmt.Errorf("querying full release: %w", err)
	}

	var tracks []*TrackDB
	trackQuery := `
	SELECT 
		id AS id,
		release_id AS release_id,
		name AS name,
		lyrics AS lyrics,
		file AS file,
		duration AS duration,
		created_at AS created_at
	FROM tracks
	WHERE release_id = ?
	`

	if err := r.db.Select(&tracks, trackQuery, release.IDField); err != nil {
		return nil, fmt.Errorf("querying tracks for release %s: %w", release.IDField, err)
	}

	for _, track := range tracks {
		var segments []*SegmentDB
		segmentQuery := `
		SELECT 
			track_id AS track_id,
			name AS name,
			position AS position
		FROM segments
		WHERE track_id = ?
		`
		if err := r.db.Select(&segments, segmentQuery, track.IDField); err != nil {
			return nil, fmt.Errorf("querying segments for track %s: %w", track.IDField, err)
		}
		track.SegmentsField = segments
	}

	release.TracksField = tracks

	return CreateFullRelease(release), nil
}

func (r *repo) GetMostRecentRelease() (domain.FullRelease, error) {
	var release FullReleaseDB

	query := `
	SELECT 
		r.id AS id,
		r.name AS name,
		r.type AS type,
		r.cover AS cover,
		r.release_date AS release_date,
		r.artist_id AS artist_id,
		r.created_at AS created_at,
		a.id AS "artist.id",
		a.name AS "artist.name",
		a.cover AS "artist.cover",
		a.created_at AS "artist.created_at"
	FROM releases r
	JOIN artists a ON r.artist_id = a.id
	ORDER BY r.created_at DESC
	LIMIT 1
	`

	if err := r.db.Get(&release, query); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("no releases found: %w", err)
		}
		return nil, fmt.Errorf("querying most recent release: %w", err)
	}

	var tracks []*TrackDB
	trackQuery := `
	SELECT 
		id AS id,
		release_id AS release_id,
		name AS name,
		lyrics AS lyrics,
		file AS file,
		duration AS duration,
		created_at AS created_at
	FROM tracks
	WHERE release_id = ?
	`

	if err := r.db.Select(&tracks, trackQuery, release.IDField); err != nil {
		return nil, fmt.Errorf("querying tracks: %w", err)
	}

	for _, track := range tracks {
		var segments []*SegmentDB
		segmentQuery := `
		SELECT 
			track_id AS track_id,
			name AS name,
			position AS position
		FROM segments
		WHERE track_id = ?
		`

		if err := r.db.Select(&segments, segmentQuery, track.IDField); err != nil {
			return nil, fmt.Errorf("querying segments: %w", err)
		}
		track.SegmentsField = segments
	}

	release.TracksField = tracks

	return CreateFullRelease(release), nil
}

