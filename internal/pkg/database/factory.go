package database

import (
	"context"
	"fmt"
)

// Factory creates database drivers based on configuration
type Factory struct{}

// NewFactory creates a new database factory
func NewFactory() *Factory {
	return &Factory{}
}

// CreateDriver creates a database driver based on the driver type
func (f *Factory) CreateDriver(driverType DriverType, config interface{}) (Driver, error) {
	switch driverType {
	case DriverPostgreSQL:
		cfg, ok := config.(*PostgresConfig)
		if !ok {
			return nil, fmt.Errorf("invalid postgres config type")
		}
		return NewPostgresDriver(cfg), nil

	case DriverMySQL:
		cfg, ok := config.(*MySQLConfig)
		if !ok {
			return nil, fmt.Errorf("invalid mysql config type")
		}
		return NewMySQLDriver(cfg), nil

	case DriverMongoDB:
		// MongoDB implementation would go here
		return nil, fmt.Errorf("mongodb driver not yet implemented")

	case DriverSQLite:
		// SQLite implementation would go here
		return nil, fmt.Errorf("sqlite driver not yet implemented")

	default:
		return nil, fmt.Errorf("unsupported driver type: %s", driverType)
	}
}

// Manager manages multiple database connections
type Manager struct {
	drivers map[string]Driver
	factory *Factory
}

// NewManager creates a new database manager
func NewManager() *Manager {
	return &Manager{
		drivers: make(map[string]Driver),
		factory: NewFactory(),
	}
}

// AddDriver adds a database driver with a name
func (m *Manager) AddDriver(name string, driver Driver) error {
	if _, exists := m.drivers[name]; exists {
		return fmt.Errorf("driver with name %s already exists", name)
	}
	m.drivers[name] = driver
	return nil
}

// GetDriver retrieves a driver by name
func (m *Manager) GetDriver(name string) (Driver, error) {
	driver, exists := m.drivers[name]
	if !exists {
		return nil, fmt.Errorf("driver with name %s not found", name)
	}
	return driver, nil
}

// ConnectAll connects all registered drivers
func (m *Manager) ConnectAll(ctx context.Context) error {
	for name, driver := range m.drivers {
		if err := driver.Connect(ctx); err != nil {
			return fmt.Errorf("failed to connect driver %s: %w", name, err)
		}
	}
	return nil
}

// CloseAll closes all registered drivers
func (m *Manager) CloseAll() error {
	var errs []error
	for name, driver := range m.drivers {
		if err := driver.Close(); err != nil {
			errs = append(errs, fmt.Errorf("failed to close driver %s: %w", name, err))
		}
	}
	if len(errs) > 0 {
		return fmt.Errorf("errors closing drivers: %v", errs)
	}
	return nil
}

// Health checks the health of all drivers
func (m *Manager) Health(ctx context.Context) map[string]error {
	results := make(map[string]error)
	for name, driver := range m.drivers {
		results[name] = driver.Health(ctx)
	}
	return results
}

