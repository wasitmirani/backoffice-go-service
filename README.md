# Backoffice Go Service

A production-ready Golang starter kit built with Gin framework, featuring clean architecture, multi-database support, daily file logging, and comprehensive CI/CD pipeline.

## 🚀 Features

- **Gin Framework** - High-performance HTTP web framework for Go
- **Clean Architecture** - Controller → Service → Database pattern
- **Multi-Database Support** - PostgreSQL, MySQL with easy extensibility
- **Daily File Logging** - Automatic log rotation with daily file generation
- **JWT Authentication** - Secure token-based authentication
- **Multi-App Support** - Designed for multiple applications
- **Docker Support** - Containerization-ready with optimized Dockerfile
- **CI/CD Pipeline** - GitHub Actions for automated testing and deployment
- **Environment Configuration** - Flexible .env-based configuration
- **Health Checks** - Built-in health and readiness endpoints
- **Graceful Shutdown** - Proper application lifecycle management

## 📁 Project Structure

```
.
├── cmd/                          # Application entry points
│   └── main.go                  # Main application entry
├── config/                       # Configuration management
│   └── config.go                # Configuration loader
├── internal/                     # Internal application code
│   ├── app/                     # Application layer
│   │   ├── app.go              # Application initialization
│   │   ├── controllers/        # HTTP controllers
│   │   │   ├── auth/          # Authentication controllers
│   │   │   └── user/          # User management controllers
│   │   ├── models/            # Domain models
│   │   └── middleware/        # HTTP middleware
│   ├── services/                # Business logic layer
│   │   ├── auth_service.go
│   │   └── user_service.go
│   ├── pkg/                     # Shared packages
│   │   ├── database/           # Database abstraction layer
│   │   ├── logger/             # Logging (stdout/file/stack)
│   │   ├── errors/             # Error handling
│   │   ├── utils/              # Utilities (JWT, hash)
│   │   └── validator/          # Validation
│   ├── routes/                  # Route definitions
│   ├── database/                # Database migrations
│   └── infrastructure/          # External integrations
├── deployments/                  # Deployment configs
│   ├── docker/
│   └── kubernetes/
├── scripts/                      # Utility scripts
├── tests/                        # Test files
├── .github/                      # GitHub Actions workflows
├── Dockerfile                    # Docker image definition
├── Makefile                      # Build automation
└── go.mod                        # Go module definition
```

## 🛠️ Getting Started

### Prerequisites

- Go 1.24 or higher
- PostgreSQL or MySQL (or any supported database)
- Git
- Docker (optional, for containerized deployment)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourorg/backoffice-go-service.git
   cd backoffice-go-service
   ```

2. **Set up environment variables**
   ```bash
   cp ENV_EXAMPLE.md .env
   # Edit .env with your configuration
   ```

3. **Install dependencies**
   ```bash
   go mod download
   go mod tidy
   ```

4. **Configure database**
   - Update database credentials in `.env`
   - Run migrations (if available)

5. **Run the application**
   ```bash
   # Development mode
   make dev
   
   # Or directly
   go run cmd/main.go
   ```

## 🔧 Configuration

### Environment Variables

Key configuration options (see `ENV_EXAMPLE.md` for complete list):

```bash
# Application
APP_NAME="Backoffice Service"
APP_ENV=local
APP_DEBUG=true

# Server
SERVER_HOST=0.0.0.0
SERVER_PORT=8080
GIN_MODE=debug

# Database
DB_DRIVER=postgresql
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=backoffice
DB_USE_GORM=true

# JWT
JWT_SECRET=your-secret-key-min-32-chars
JWT_EXPIRATION=24h

# Logging
LOG_CHANNEL=file          # stdout, file, or stack
LOG_LEVEL=debug
LOG_FILE_PATH=./storage/logs
LOG_FILE_NAME=app
LOG_DAILY_ROTATE=true     # Daily files or single file
LOG_MAX_SIZE=10           # MB
LOG_MAX_BACKUPS=5
LOG_MAX_AGE=28            # days
```

### Logging Configuration

The application supports three logging modes:

- **stdout**: Logs to console (default for development)
- **file**: Logs to daily rotated files (`app-2024-01-15.log`)
- **stack**: Logs to both stdout and file

Set `LOG_DAILY_ROTATE=true` for daily files or `false` for single file with size-based rotation.

## 📦 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - Login user
- `POST /api/v1/auth/logout` - Logout user
- `POST /api/v1/auth/refresh` - Refresh JWT token

### Users
- `GET /api/v1/users` - List users (with pagination)
- `GET /api/v1/users/:id` - Get user by ID
- `POST /api/v1/users` - Create user
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user

### Health
- `GET /health` - Health check
- `GET /ready` - Readiness check

## 🏗️ Architecture

### Controller → Service → Database

```
HTTP Request
    ↓
Controller (HTTP handling, validation)
    ↓
Service (Business logic)
    ↓
Database Driver (Data access)
```

### Database Drivers

The application supports multiple database drivers through a clean abstraction:

- **PostgreSQL** - Full support with GORM and raw SQL
- **MySQL** - Full support with GORM and raw SQL
- **Extensible** - Easy to add MongoDB, SQLite, etc.

### Multi-Database Support

Configure multiple databases in `.env`:

```bash
# Primary database
DB_DRIVER=postgresql
DB_HOST=localhost
...

# Secondary database (optional)
DB_SECONDARY_DRIVER=mysql
DB_SECONDARY_HOST=analytics-db
...
```

## 🐳 Docker

### Build Docker Image

```bash
make docker-build
# Or
docker build -t backoffice-service:latest .
```

### Run Container

```bash
make docker-run
# Or
docker run -p 8080:8080 --env-file .env backoffice-service:latest
```

### Docker Compose

```bash
cd deployments/docker
docker-compose up -d
```

## 🧪 Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage

# Run verbose tests
make test-verbose
```

## 🔍 Code Quality

```bash
# Run linters
make lint

# Fix linting issues
make lint-fix

# Run security scan
make security

# Format code
make fmt
```

## 🚀 CI/CD

The project includes comprehensive GitHub Actions workflows:

- **CI**: Automated testing, linting, building, and security scanning
- **CD**: Automated Docker builds and deployments
- **Release**: Automated release creation with binaries

See [CI_CD.md](CI_CD.md) for detailed documentation.

## 📝 Makefile Commands

```bash
make help          # Show all available commands
make build         # Build the application
make run           # Run the application
make test          # Run tests
make lint          # Run linters
make docker-build  # Build Docker image
make docker-run    # Run Docker container
make dev           # Run in development mode
make install-tools # Install development tools
```

## 🔐 Security

- JWT-based authentication
- Password hashing with bcrypt
- Security scanning with gosec
- Non-root Docker user
- Input validation

## 📚 Documentation

- [ARCHITECTURE.md](ARCHITECTURE.md) - Architecture overview
- [CI_CD.md](CI_CD.md) - CI/CD documentation
- [LOGGING_SETUP.md](LOGGING_SETUP.md) - Logging configuration
- [ENV_EXAMPLE.md](ENV_EXAMPLE.md) - Environment variables reference
- [STRUCTURE_OPTIMIZATION.md](STRUCTURE_OPTIMIZATION.md) - Structure guide

## 🛣️ Roadmap

- [ ] Add MongoDB driver implementation
- [ ] Add SQLite driver implementation
- [ ] Add Redis caching layer
- [ ] Add rate limiting middleware
- [ ] Add API documentation (Swagger)
- [ ] Add database migrations system
- [ ] Add monitoring and metrics

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Commit Message Format

Follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <subject>

<body>

<footer>
```

Types: `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`, `revert`

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 👤 Author

Your Name - [GitHub Profile](https://github.com/yourusername)

## 🙏 Acknowledgments

- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [Viper](https://github.com/spf13/viper)
- Go community for excellent libraries and tools

---

⭐ Star this repo if you find it useful!
