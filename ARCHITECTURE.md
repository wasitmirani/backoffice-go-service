# Architecture Documentation

## Overview

This is an optimized Go starter kit designed for building scalable applications with support for multiple database drivers and multiple applications. The architecture follows clean architecture principles with clear separation of concerns.

## Project Structure

```
.
├── cmd/                          # Application entry points
│   └── main.go                  # Main application entry
│
├── config/                       # Configuration management
│   └── config.go                # Configuration loader with multi-DB support
│
├── internal/                     # Internal application code
│   ├── app/                     # Application layer
│   │   ├── app.go              # Main application struct and initialization
│   │   ├── controllers/        # HTTP controllers/handlers
│   │   │   ├── auth/
│   │   │   └── user/
│   │   ├── models/             # Domain models
│   │   └── middleware/         # HTTP middleware
│   │
│   ├── pkg/                     # Shared packages
│   │   ├── database/           # Database abstraction layer
│   │   │   ├── driver.go       # Driver interface
│   │   │   ├── factory.go      # Driver factory and manager
│   │   │   ├── postgres.go     # PostgreSQL driver implementation
│   │   │   ├── mysql.go        # MySQL driver implementation
│   │   │   └── errors.go       # Database-specific errors
│   │   │
│   │   ├── repository/         # Repository interfaces
│   │   │   └── repository.go   # Base repository interfaces
│   │   │
│   │   ├── service/            # Service interfaces
│   │   │   └── service.go      # Service interfaces
│   │   │
│   │   ├── logger/             # Logging package
│   │   │   └── logger.go       # Logger interface and implementation
│   │   │
│   │   ├── errors/             # Error handling
│   │   │   └── app_error.go    # Application error types
│   │   │
│   │   ├── utils/              # Utility functions
│   │   │   ├── hash.go
│   │   │   └── jwt.go
│   │   │
│   │   └── validator/          # Validation utilities
│   │       └── validator.go
│   │
│   ├── services/                # Business logic layer
│   │   ├── auth_service.go
│   │   └── user/
│   │       └── user_service.go
│   │
│   ├── routes/                  # Route definitions
│   │   └── routes.go
│   │
│   ├── database/                # Database-specific code
│   │   └── migrations/          # Database migrations
│   │
│   └── infrastructure/          # External service integrations
│       ├── email/
│       ├── messaging/
│       ├── redis/
│       └── storage/
│
├── deployments/                 # Deployment configurations
│   ├── docker/
│   └── kubernetes/
│
├── scripts/                     # Utility scripts
│   ├── deploy.sh
│   ├── migrate.sh
│   └── seed.sh
│
├── tests/                       # Test files
│
├── go.mod                       # Go module definition
├── go.sum                       # Go dependencies
└── README.md                    # Project documentation
```

## Key Features

### 1. Multi-Database Support

The architecture supports multiple database drivers through a clean abstraction layer:

- **PostgreSQL** - Full support with GORM and raw SQL
- **MySQL** - Full support with GORM and raw SQL
- **MongoDB** - Interface ready (implementation pending)
- **SQLite** - Interface ready (implementation pending)

#### Usage Example:

```go
// In config/config.go, you can configure multiple databases:
cfg.Database.Primary = DatabaseConnectionConfig{
    Driver: "postgresql",
    Host: "localhost",
    // ... other config
}

cfg.Database.Databases["analytics"] = DatabaseConnectionConfig{
    Driver: "mysql",
    Host: "analytics-db",
    // ... other config
}
```

### 2. Database Driver Abstraction

All database operations go through the `Driver` interface, making it easy to:
- Switch between databases
- Support multiple databases simultaneously
- Mock databases for testing
- Add new database drivers

### 3. Repository Pattern

The repository pattern provides a clean abstraction for data access:
- Interface-based design
- Easy to mock for testing
- Database-agnostic business logic

### 4. Service Layer

Business logic is separated into services:
- Clean interfaces
- Testable
- Reusable across different applications

### 5. Multi-App Support

The structure supports running multiple applications:
- Each app can have its own entry point in `cmd/`
- Shared internal packages
- Independent configurations

## Architecture Layers

### 1. Presentation Layer (`controllers/`, `routes/`)
- HTTP handlers
- Request/response handling
- Input validation

### 2. Application Layer (`app/`)
- Application initialization
- Dependency injection
- Route setup
- Middleware configuration

### 3. Business Logic Layer (`services/`)
- Business rules
- Use cases
- Domain logic

### 4. Data Access Layer (`repository/`, `database/`)
- Data persistence
- Database operations
- Query optimization

### 5. Infrastructure Layer (`infrastructure/`)
- External services
- Third-party integrations
- Caching, messaging, storage

## Configuration

Configuration is managed through environment variables and config files:

```go
// Environment variables:
DB_DRIVER=postgresql
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=backoffice
DB_USE_GORM=true

// Multiple databases:
DB_ANALYTICS_DRIVER=mysql
DB_ANALYTICS_HOST=analytics-db
// ...
```

## Dependency Injection

Dependencies are initialized in `app/app.go`:
1. Configuration loading
2. Database connections
3. Repositories
4. Services
5. Controllers/Handlers
6. Routes

## Error Handling

Centralized error handling through `internal/pkg/errors/app_error.go`:
- Structured errors
- HTTP status code mapping
- Error wrapping

## Logging

Structured logging through `internal/pkg/logger/logger.go`:
- Interface-based design
- Easy to swap implementations
- Structured fields support

## Best Practices

1. **Interfaces First**: Define interfaces before implementations
2. **Dependency Injection**: Pass dependencies through constructors
3. **Error Handling**: Always handle errors explicitly
4. **Context Propagation**: Use context.Context for cancellation and timeouts
5. **Testing**: Write tests for all layers
6. **Documentation**: Document public APIs

## Adding a New Database Driver

1. Create a new file in `internal/pkg/database/` (e.g., `mongodb.go`)
2. Implement the `Driver` interface
3. Add the driver type to `DriverType` enum
4. Update `factory.go` to support the new driver
5. Add configuration support in `config/config.go`

## Adding a New Application

1. Create a new entry point in `cmd/` (e.g., `cmd/admin/main.go`)
2. Create application-specific config if needed
3. Initialize the app with `app.New()`
4. Configure routes and middleware
5. Start the server

## Testing Strategy

- Unit tests for services and repositories
- Integration tests for database operations
- E2E tests for API endpoints
- Mock interfaces for external dependencies

