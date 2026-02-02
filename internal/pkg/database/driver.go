package database

import (
	"context"
	"database/sql"
)

// DriverType represents the type of database driver
type DriverType string

const (
	DriverPostgreSQL DriverType = "postgresql"
	DriverMySQL      DriverType = "mysql"
	DriverMongoDB    DriverType = "mongodb"
	DriverSQLite     DriverType = "sqlite"
)

// Driver interface for database operations
type Driver interface {
	// Connect establishes a connection to the database
	Connect(ctx context.Context) error
	
	// Close closes the database connection
	Close() error
	
	// Ping checks if the database connection is alive
	Ping(ctx context.Context) error
	
	// GetDB returns the underlying database connection
	GetDB() interface{}
	
	// GetSQLDB returns *sql.DB for SQL databases (nil for NoSQL)
	GetSQLDB() *sql.DB
	
	// GetGormDB returns *gorm.DB for GORM operations (nil if not using GORM)
	GetGormDB() interface{}
	
	// Type returns the driver type
	Type() DriverType
	
	// Health checks the health of the database connection
	Health(ctx context.Context) error
}

// Transaction interface for database transactions
type Transaction interface {
	Begin(ctx context.Context) (interface{}, error)
	Commit(ctx context.Context, tx interface{}) error
	Rollback(ctx context.Context, tx interface{}) error
}

