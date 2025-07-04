package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Environment struct {
	WebStaticFilesDir string
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
	return Environment{
		WebStaticFilesDir: staticFilesPath,
	}
}
