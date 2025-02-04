package services

import (
	"context"
	"log"
	"time"

	"bitsnake-server/internal/ent"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func InitDB(dsn string) (*ent.Client, error) {
	// Connect to the database directly
	log.Println("Connecting to the database...")
	client, err := ent.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Database connected successfully!")

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Run the migration
	log.Println("Running database migrations...")
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database connected and schema migrated successfully!")
	return client, nil
}
