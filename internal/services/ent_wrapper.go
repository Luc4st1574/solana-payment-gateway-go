package services

import (
	"bitsnake-server/internal/ent"
	"context"
)

// DatabaseClient is an interface for database operations.
type DatabaseClient interface {
	Get(ctx context.Context, id int) (*ent.Users, error)
	Client() *ent.Client // Exposes the raw *ent.Client
}

// EntDatabaseClient wraps the ent.Client to implement the DatabaseClient interface.
type EntDatabaseClient struct {
	client *ent.Client
}

// NewEntDatabaseClient creates a new EntDatabaseClient.
func NewEntDatabaseClient(client *ent.Client) *EntDatabaseClient {
	return &EntDatabaseClient{client: client}
}

// Get retrieves a user by ID.
func (e *EntDatabaseClient) Get(ctx context.Context, id int) (*ent.Users, error) {
	return e.client.Users.Get(ctx, id)
}

// Client exposes the raw *ent.Client.
func (e *EntDatabaseClient) Client() *ent.Client {
	return e.client
}
