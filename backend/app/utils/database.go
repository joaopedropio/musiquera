package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
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
		position INT NOT NULL,
		FOREIGN KEY (track_id) REFERENCES tracks(id)
	)
	`
}

func MustCreateTestSqliteDatabase() (string, *sqlx.DB) {
	dbName := uuid.New().String() + ".db"
	db, err := sqlx.Open("sqlite3", dbName+"?_foreign_keys=on")
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
