package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	CLIENT_ID     string
	CLIENT_SECRET string

	STATE string
	PORT  string

	KEYCLOAK_ISSUER string
)

func LoadConfig(envPath string) {
	log.Println("Config package initialized")

	log.Println("Loading environment from: ", envPath)
	if err := godotenv.Load(envPath); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	CLIENT_ID = os.Getenv("CLIENT_ID")
	CLIENT_SECRET = os.Getenv("CLIENT_SECRET")

	STATE = os.Getenv("STATE")
	PORT = os.Getenv("PORT")

	KEYCLOAK_ISSUER = os.Getenv("KEYCLOAK_ISSUER")
}
