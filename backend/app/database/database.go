package database

import (
	"database/sql/driver"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	domain "github.com/joaopedropio/musiquera/app/domain/entity"
	_ "modernc.org/sqlite"
)

func DatabaseSchema() string {
	return `
	CREATE TABLE IF NOT EXISTS users (
		id TEXT PRIMARY KEY CHECK (length(id) = 36),
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS invites (
		id TEXT PRIMARY KEY CHECK (length(id) = 36),
		user_id TEXT,
		code TEXT,
		status TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS artists (
		id TEXT PRIMARY KEY CHECK (length(id) = 36),
		name TEXT NOT NULL,
		cover TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);

	CREATE TABLE IF NOT EXISTS releases (
		id TEXT PRIMARY KEY CHECK (length(id) = 36),
		name TEXT NOT NULL,
		cover TEXT NOT NULL,
		type TEXT NOT NULL,
		release_date DATETIME NOT NULL,
		artist_id TEXT NOT NULL CHECK(length(id) = 36),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (artist_id) REFERENCES artists(id)
	);

	CREATE TABLE IF NOT EXISTS tracks (
		id TEXT PRIMARY KEY CHECK (length(id) = 36),
		name TEXT NOT NULL,
		lyrics TEXT,
		file TEXT NOT NULL,
		duration INT NOT NULL,
		release_id TEXT NOT NULL CHECK (length(id) = 36),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (release_id) REFERENCES releases(id)
	);

	CREATE TABLE IF NOT EXISTS segments (
		track_id TEXT NOT NULL CHECK (length(track_id) = 36),
		name TEXT NOT NULL,
		position INT NOT NULL,
		FOREIGN KEY (track_id) REFERENCES tracks(id)
	)
	`
}

func MustCreateTestSqliteDatabase() (string, *sqlx.DB) {
	dbName := uuid.New().String() + ".db"
	db, err := sqlx.Open("sqlite", dbName+"?_foreign_keys=on")
	if err != nil {
		panic(fmt.Errorf("unable to create sqlite db test connection: name %s, %w", dbName, err))
	}
	_, err = db.Exec(DatabaseSchema())
	if err != nil {
		MustDestroySqliteDatabase(dbName, db)
		panic(fmt.Errorf("unable to run migrations: %w", err))
	}
	return dbName, db
}

func MustDestroySqliteDatabase(dbName string, db *sqlx.DB) {
	if err := db.Close(); err != nil {
		log.Fatalf("unable to destroy database %s: %v", dbName, err)
	}
	if err := os.Remove(dbName); err != nil {
		log.Fatalf("unable to remove database %s: %v", dbName, err)
	}
}

func CommitOrRollback(tx *sqlx.Tx, errPtr *error) {
	if *errPtr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			*errPtr = fmt.Errorf("rollback failed: %v (original error: %w)", rollbackErr, *errPtr)
		}
		return
	}
	if commitErr := tx.Commit(); commitErr != nil {
		*errPtr = fmt.Errorf("commit failed: %w", commitErr)
	}
}

func NewDateDB(d domain.Date) *DateDB {
	return &DateDB{
		value: d,
	}
}

type DateDB struct {
	value domain.Date
}

func (d *DateDB) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		d.value = domain.NewDate(v.Year(), int(v.Month()), v.Day())
		return nil
	case string:
		t, err := time.Parse("2006-01-02", v)
		if err != nil {
			return fmt.Errorf("unable to parse date string: %w", err)
		}
		d.value = domain.NewDate(t.Year(), int(t.Month()), t.Day())
		return nil
	default:
		return fmt.Errorf("unsupported Scan type for Date: %T", value)
	}
}

func (d DateDB) Value() (driver.Value, error) {
	return d.value.String(), nil
}

func (d DateDB) String() string {
	return d.value.String()
}

func (d DateDB) Date() domain.Date {
	return d.value
}

type NullUUID struct {
	UUID  uuid.UUID
	Valid bool
}

func NewNullUUID(id *uuid.UUID) NullUUID {
	if id == nil {
		return NullUUID{
			uuid.Nil,
			false,
		}
	}

	return NullUUID{
		*id,
		true,
	}
}

// Scan implements the Scanner interface (for reading from DB)
func (n *NullUUID) Scan(value interface{}) error {
	if value == nil {
		n.UUID, n.Valid = uuid.UUID{}, false
		return nil
	}

	switch v := value.(type) {
	case string:
		id, err := uuid.Parse(v)
		if err != nil {
			return fmt.Errorf("invalid uuid string: %w", err)
		}
		n.UUID = id
		n.Valid = true
		return nil
	case []byte:
		id, err := uuid.ParseBytes(v)
		if err != nil {
			return fmt.Errorf("invalid uuid bytes: %w", err)
		}
		n.UUID = id
		n.Valid = true
		return nil
	default:
		return fmt.Errorf("cannot scan UUID from %T", value)
	}
}

// Value implements the driver.Valuer interface (for writing to DB)
func (n NullUUID) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.UUID.String(), nil
}

// Ptr returns a *uuid.UUID or nil
func (n NullUUID) Ptr() *uuid.UUID {
	if !n.Valid {
		return nil
	}
	return &n.UUID
}
