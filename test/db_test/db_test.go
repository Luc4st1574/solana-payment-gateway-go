package test

import (
	"context"
	"testing"

	"bitsnake-server/internal/ent"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestMockDatabaseClient is a mock implementation of the database client.
type TestMockDatabaseClient struct {
	mock.Mock
}

func (m *TestMockDatabaseClient) Open(driver, source string) (*ent.Client, error) {
	args := m.Called(driver, source)
	return args.Get(0).(*ent.Client), args.Error(1)
}

func (m *TestMockDatabaseClient) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *TestMockDatabaseClient) CreateSchema(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestInitDB(t *testing.T) {
	// Mock variables
	dsn := "mock-dsn"
	mockClient := new(TestMockDatabaseClient)
	mockEntClient := &ent.Client{}
	ctx := context.Background()

	// Setup expectations
	mockClient.On("Open", "postgres", dsn).Return(mockEntClient, nil)
	mockClient.On("CreateSchema", ctx).Return(nil)
	mockClient.On("Close").Return(nil)

	// Test database initialization
	t.Log("Starting InitDB test...")
	client, err := mockClient.Open("postgres", dsn)
	assert.NoError(t, err, "Expected no error during database initialization")
	assert.NotNil(t, client, "Expected a valid database client")
	t.Log("Database initialization succeeded.")

	// Test schema creation
	t.Log("Testing schema creation...")
	err = mockClient.CreateSchema(ctx)
	assert.NoError(t, err, "Expected no error during schema creation")
	t.Log("Schema creation succeeded.")

	// Test database cleanup
	t.Log("Testing database cleanup...")
	err = mockClient.Close()
	assert.NoError(t, err, "Expected no error during database closure")
	t.Log("Database cleanup succeeded.")

	// Assert expectations
	mockClient.AssertExpectations(t)
}
