package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDriver implements the Driver interface for MySQL
type MySQLDriver struct {
	config *MySQLConfig
	db     *sql.DB
	gormDB *gorm.DB
}

// MySQLConfig holds MySQL configuration
type MySQLConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	Charset         string
	ParseTime       bool
	Loc             string
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
	UseGorm         bool
}

// NewMySQLDriver creates a new MySQL driver instance
func NewMySQLDriver(cfg *MySQLConfig) *MySQLDriver {
	if cfg.Charset == "" {
		cfg.Charset = "utf8mb4"
	}
	if cfg.ParseTime {
		if cfg.Loc == "" {
			cfg.Loc = "Local"
		}
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

	return &MySQLDriver{
		config: cfg,
	}
}

// Connect establishes a connection to MySQL
func (d *MySQLDriver) Connect(ctx context.Context) error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s",
		d.config.User,
		d.config.Password,
		d.config.Host,
		d.config.Port,
		d.config.DBName,
		d.config.Charset,
	)

	if d.config.ParseTime {
		dsn += "&parseTime=True"
		if d.config.Loc != "" {
			dsn += "&loc=" + d.config.Loc
		}
	}

	var err error
	d.db, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open mysql connection: %w", err)
	}

	// Set connection pool settings
	d.db.SetMaxOpenConns(d.config.MaxOpenConns)
	d.db.SetMaxIdleConns(d.config.MaxIdleConns)
	d.db.SetConnMaxLifetime(d.config.ConnMaxLifetime)
	d.db.SetConnMaxIdleTime(d.config.ConnMaxIdleTime)

	// Test connection
	if err := d.db.PingContext(ctx); err != nil {
		return fmt.Errorf("failed to ping mysql: %w", err)
	}

	// Initialize GORM if requested
	if d.config.UseGorm {
		d.gormDB, err = gorm.Open(mysql.New(mysql.Config{
			Conn: d.db,
		}), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("failed to initialize gorm: %w", err)
		}
	}

	return nil
}

// Close closes the database connection
func (d *MySQLDriver) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

// Ping checks if the database connection is alive
func (d *MySQLDriver) Ping(ctx context.Context) error {
	if d.db == nil {
		return fmt.Errorf("database connection is not established")
	}
	return d.db.PingContext(ctx)
}

// GetDB returns the underlying database connection
func (d *MySQLDriver) GetDB() interface{} {
	if d.config.UseGorm && d.gormDB != nil {
		return d.gormDB
	}
	return d.db
}

// GetSQLDB returns *sql.DB
func (d *MySQLDriver) GetSQLDB() *sql.DB {
	return d.db
}

// GetGormDB returns *gorm.DB if GORM is enabled
func (d *MySQLDriver) GetGormDB() interface{} {
	return d.gormDB
}

// Type returns the driver type
func (d *MySQLDriver) Type() DriverType {
	return DriverMySQL
}

// Health checks the health of the database connection
func (d *MySQLDriver) Health(ctx context.Context) error {
	return d.Ping(ctx)
}

