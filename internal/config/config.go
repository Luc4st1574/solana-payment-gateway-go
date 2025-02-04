package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort            string
	DBHost                string
	DBPort                string
	DBUser                string
	DBPassword            string
	DBName                string
	SolanaNetwork         string
	SolanaRecieverWallet  string
	SolanaRecieverPrivKey string
	SolanaSenderWallet    string
	SolanaSenderPrivKey   string
}

// LoadConfig loads configuration from the specified .env file path or the default path
func LoadConfig(envPath ...string) (*Config, error) {
	envFile := ".env"
	if len(envPath) > 0 {
		envFile = envPath[0]
	}

	log.Printf("Loading .env file from: %s", envFile)

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("No .env file found at %s, falling back to environment variables", envFile)
	}

	return &Config{
		ServerPort:            getEnv("SERVER_PORT", "8080"),
		DBHost:                getEnv("DB_HOST", "localhost"),
		DBPort:                getEnv("DB_PORT", "5432"),
		DBUser:                mustGetEnv("DB_USER"),
		DBPassword:            mustGetEnv("DB_PASSWORD"),
		DBName:                mustGetEnv("DB_NAME"),
		SolanaNetwork:         mustGetEnv("SOLANA_RPC_URL"),
		SolanaRecieverWallet:  mustGetEnv("SOLANA_RECIEVER_WALLET"),
		SolanaRecieverPrivKey: mustGetEnv("SOLANA_RECIEVER_PRIVATE_KEY"),
		SolanaSenderWallet:    mustGetEnv("SOLANA_SENDER_WALLET"),
		SolanaSenderPrivKey:   mustGetEnv("SOLANA_SENDER_PRIVATE_KEY"),
	}, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func mustGetEnv(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		log.Fatalf("Environment variable %s is required but not set. Ensure it's in the .env file or exported in the environment.", key)
	}
	return value
}
