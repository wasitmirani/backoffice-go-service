// Package database provides database driver implementations with dynamic imports
// This file ensures database drivers are registered at compile time
package database

// Import database drivers to register them with database/sql
// These blank imports ensure the drivers are available when needed
import (
	// PostgreSQL driver
	_ "github.com/lib/pq"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	// Future drivers can be added here:
	// _ "github.com/mattn/go-sqlite3" // SQLite
	// _ "go.mongodb.org/mongo-driver/mongo" // MongoDB (NoSQL, different approach)
)

// init ensures all database drivers are registered
// This function runs automatically when the package is imported
func init() {
	// Drivers are automatically registered via their blank imports above
	// No additional initialization needed
}
