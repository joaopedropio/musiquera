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
		status TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`
}

func MustCreateTestSqliteDatabase() (string, *sqlx.DB) {
	dbName := uuid.New().String() + ".db"
	db, err := sqlx.Open("sqlite3", dbName+"?_foreign_keys=on")
	if err != nil {
		panic(fmt.Errorf("unable to create sqlite db test connection: name %s, %w", dbName, err))
	}
	db.MustExec(DatabaseSchema())
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
