package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SpotifyClientID     string
	SpotifyClientSecret string
	DBPath              string
}

var AppConfig Config

func Load() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default values")
	}

	AppConfig = Config{
		SpotifyClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		SpotifyClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		DBPath:              "db/idlyfy.sqlite",
	}

	return nil
}
