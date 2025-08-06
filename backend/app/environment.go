package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Environment struct {
	WebStaticFilesDir string
	JWTSecret         string
	DatabaseDir       string
	HTTPPort          string
}

func GetEnvironmentVariables() Environment {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("unable to load go dot env: %s\n", err.Error())
	}
	staticFilesPath := os.Getenv("STATIC_FILES")
	if staticFilesPath == "" {
		panic("STATIC_FILES env var not set")
	}
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		panic("JWT SECRET env var not set")
	}
	dbDir := os.Getenv("DATABASE_DIR")
	if dbDir == "" || !dirExists(dbDir) {
		panic("DATABASE_DIR should be a valid directory")
	}
	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	return Environment{
		WebStaticFilesDir: staticFilesPath,
		JWTSecret:         jwtSecret,
		DatabaseDir:       dbDir,
		HTTPPort:          ":" + httpPort,
	}
}

func dirExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}
