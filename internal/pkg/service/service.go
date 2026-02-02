package service

import (
	"context"
)

// Service is the base interface for all services
type Service interface {
	// Health checks if the service is healthy
	Health(ctx context.Context) error
}

// UserService defines the interface for user business logic
type UserService interface {
	Service
	
	// CreateUser creates a new user
	CreateUser(ctx context.Context, req interface{}) (interface{}, error)
	
	// GetUser retrieves a user by ID
	GetUser(ctx context.Context, id string) (interface{}, error)
	
	// GetUserByEmail retrieves a user by email
	GetUserByEmail(ctx context.Context, email string) (interface{}, error)
	
	// UpdateUser updates an existing user
	UpdateUser(ctx context.Context, id string, req interface{}) (interface{}, error)
	
	// DeleteUser deletes a user by ID
	DeleteUser(ctx context.Context, id string) error
	
	// ListUsers retrieves a list of users
	ListUsers(ctx context.Context, limit, offset int) ([]interface{}, error)
}

// AuthService defines the interface for authentication business logic
type AuthService interface {
	Service
	
	// Login authenticates a user and returns a token
	Login(ctx context.Context, email, password string) (interface{}, error)
	
	// Register registers a new user
	Register(ctx context.Context, req interface{}) (interface{}, error)
	
	// ValidateToken validates a JWT token
	ValidateToken(ctx context.Context, token string) (interface{}, error)
	
	// RefreshToken refreshes an access token
	RefreshToken(ctx context.Context, refreshToken string) (interface{}, error)
	
	// Logout logs out a user
	Logout(ctx context.Context, token string) error
}

