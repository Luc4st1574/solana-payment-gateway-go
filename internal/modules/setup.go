package modules

import (
	"bitsnake-server/internal/config"
	"bitsnake-server/internal/ent"
	"bitsnake-server/internal/solana/client"
	"context"
	"log"
	"net/http"
	"regexp"

	"github.com/gofiber/fiber/v2"
)

// DatabaseClient is an interface for database interactions.
type DatabaseClient interface {
	Get(ctx context.Context, id int) (*ent.Users, error)
}

// IsValidWalletAddress validates a Solana wallet address.
func IsValidWalletAddress(address string) bool {
	match, _ := regexp.MatchString("^[1-9A-HJ-NP-Za-km-z]{32,44}$", address)
	return match
}

// fiberResponseWriter adapts a Fiber context to the http.ResponseWriter interface.
type fiberResponseWriter struct {
	ctx *fiber.Ctx
}

// Header returns an http.Header. In this simple adapter we just return an empty header.
func (frw *fiberResponseWriter) Header() http.Header {
	// For a more complete implementation you may want to track headers
	return http.Header{}
}

// Write writes the data to the Fiber response.
func (frw *fiberResponseWriter) Write(data []byte) (int, error) {
	// Write data to the Fiber response.
	return frw.ctx.Write(data)
}

// WriteHeader sets the Fiber response status code.
func (frw *fiberResponseWriter) WriteHeader(statusCode int) {
	frw.ctx.Status(statusCode)
}

// SetupRoutes initializes the routes for the application.
func SetupRoutes(app *fiber.App, cfg *config.Config, dbClient DatabaseClient, solanaClient *client.Client, db *ent.Client) {
	// Verify Payment Route
	app.Post("/verify-payment", func(c *fiber.Ctx) error {
		// Define the request payload
		type VerifyRequest struct {
			UserID         int     `json:"user_id"`
			ExpectedAmount float64 `json:"expected_amount"`
		}
		req := new(VerifyRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Fetch user's wallet address
		user, err := dbClient.Get(context.Background(), req.UserID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}

		// Validate the wallet address
		if !IsValidWalletAddress(user.WalletAddress) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid wallet address"})
		}

		// Create our adapter so that we can pass an http.ResponseWriter
		frw := &fiberResponseWriter{ctx: c}

		// Call VerifyPayment (the updated implementation writes cookies and the response)
		isPaid, err := VerifyPayment(context.Background(), db, solanaClient, frw, user.WalletAddress, req.ExpectedAmount)
		if err != nil {
			log.Printf("Payment verification error: %v", err)
			// Note: Since VerifyPayment may already have written some headers,
			// you might not want to write additional data. Adjust as needed.
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		// Optionally log the payment status.
		log.Printf("Payment verification complete: %v", isPaid)

		// Return nil because VerifyPayment already wrote the response (cookie and JSON payload)
		return nil
	})

	// Send Payment Route
	app.Post("/send-payment", func(c *fiber.Ctx) error {
		type SendRequest struct {
			UserID int     `json:"user_id"`
			Amount float64 `json:"amount"`
		}
		req := new(SendRequest)
		if err := c.BodyParser(req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}

		// Fetch user's wallet address
		user, err := dbClient.Get(context.Background(), req.UserID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}

		// Validate the wallet address
		if !IsValidWalletAddress(user.WalletAddress) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid wallet address"})
		}

		// Convert amount to lamports (1 SOL = 1e9 lamports)
		amountInLamports := uint64(req.Amount * 1e9)

		// Send payment using the updated implementation
		signature, err := SendPayment(cfg.SolanaSenderPrivKey, user.WalletAddress, amountInLamports)
		if err != nil {
			log.Printf("Send payment error: %v", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.JSON(fiber.Map{"transaction_signature": signature})
	})
}
