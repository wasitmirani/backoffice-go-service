package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// PostgresDriver implements the Driver interface for PostgreSQL
type PostgresDriver struct {
	config *PostgresConfig
	db     *sql.DB
	gormDB *gorm.DB
}

// PostgresConfig holds PostgreSQL configuration
type PostgresConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	UseGorm         bool
}

// NewPostgresDriver creates a new PostgreSQL driver instance
func NewPostgresDriver(cfg *PostgresConfig) *PostgresDriver {
	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}
	if cfg.MaxOpenConns == 0 {
		cfg.MaxOpenConns = 25
	}
	if cfg.MaxIdleConns == 0 {
		cfg.MaxIdleConns = 5
	}
	if cfg.ConnMaxLifetime == 0 {
		cfg.ConnMaxLifetime = 5 * time.Minute
	}
	if cfg.ConnMaxIdleTime == 0 {
		cfg.ConnMaxIdleTime = 10 * time.Minute
	}

	return &PostgresDriver{
		config: cfg,
	}
}

// Connect establishes a connection to PostgreSQL
func (d *PostgresDriver) Connect(ctx context.Context) error {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		d.config.Host,
		d.config.Port,
		d.config.User,
		d.config.Password,
		d.config.DBName,
		d.config.SSLMode,
	)

	var err error
	d.db, err = sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open postgres connection: %w", err)
	}

	// Set connection pool settings
	d.db.SetMaxOpenConns(d.config.MaxOpenConns)
	d.db.SetMaxIdleConns(d.config.MaxIdleConns)
	d.db.SetConnMaxLifetime(d.config.ConnMaxLifetime)
	d.db.SetConnMaxIdleTime(d.config.ConnMaxIdleTime)

	// Test connection
	if err := d.db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping postgres: %w", err)
	}

	// Initialize GORM if requested
	if d.config.UseGorm {
		d.gormDB, err = gorm.Open(postgres.New(postgres.Config{
			Conn: d.db,
		}), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to initialize gorm: %w", err)
		}
	}

	return nil
}

// Close closes the database connection
func (d *PostgresDriver) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

// Ping checks if the database connection is alive
func (d *PostgresDriver) Ping(ctx context.Context) error {
	if d.db == nil {
		return fmt.Errorf("database connection is not established")
	}
	return d.db.PingContext(ctx)
}

// GetDB returns the underlying database connection
func (d *PostgresDriver) GetDB() interface{} {
	if d.config.UseGorm && d.gormDB != nil {
		return d.gormDB
	}
	return d.db
}

// GetSQLDB returns *sql.DB
func (d *PostgresDriver) GetSQLDB() *sql.DB {
	return d.db
}

// GetGormDB returns *gorm.DB if GORM is enabled
func (d *PostgresDriver) GetGormDB() interface{} {
	return d.gormDB
}

// Type returns the driver type
func (d *PostgresDriver) Type() DriverType {
	return DriverPostgreSQL
}

// Health checks the health of the database connection
func (d *PostgresDriver) Health(ctx context.Context) error {
	return d.Ping(ctx)
}

