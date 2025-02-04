package main

import (
	"log"
	"os"
	"time"

	"bitsnake-server/internal/config"
	"bitsnake-server/internal/modules"
	"bitsnake-server/internal/services"
	"bitsnake-server/internal/solana/client"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	// === 1. Load Configuration ===
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	log.Printf("Loaded Config: SERVER_PORT=%s, DB_HOST=%s, DB_PORT=%s, DB_USER=%s, DB_NAME=%s, SOLANA_NETWORK=%s",
		cfg.ServerPort, cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.SolanaNetwork)

	// === 2. Initialize Database ===
	log.Println("Connecting to the database...")
	// Build the DSN (Data Source Name) from the config values.
	dsn := "postgres://" + cfg.DBUser + ":" + cfg.DBPassword + "@" + cfg.DBHost + ":" + cfg.DBPort + "/" + cfg.DBName + "?sslmode=disable"
	// Initialize the raw Ent client.
	rawDB, err := services.InitDB(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// Ensure the database is closed on exit.
	defer func() {
		if err := rawDB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		}
	}()
	// Wrap the raw database client with our DatabaseClient interface implementation.
	dbClient := services.NewEntDatabaseClient(rawDB)

	// === 3. Initialize Solana Client ===
	log.Println("Initializing Solana client...")
	solClient := client.New()
	if solClient == nil {
		log.Fatalf("Failed to initialize Solana client")
	}

	// === 4. Initialize Fiber App and Middleware ===
	app := fiber.New()

	// Set up request logging to a file.
	logFile, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()
	app.Use(logger.New(logger.Config{
		Output: logFile,
	}))

	// Apply a simple rate limiter middleware.
	app.Use(limiter.New(limiter.Config{
		Max:        100,
		Expiration: 60 * time.Second,
	}))

	// === 5. Setup Application Routes ===
	// The SetupRoutes function registers endpoints for payment verification,
	// sending payments, and game session creation.
	modules.SetupRoutes(app, cfg, dbClient, solClient, rawDB)

	// === 6. Start the Server ===
	log.Printf("Starting server on port %s...", cfg.ServerPort)
	if err := app.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
