Reading
Reading
Golang Gin Boilerplate - Personal Starter Kit for Web Services
A clean, structured, and ready-to-use Golang Gin Web Service Boilerplate, perfect for quickly starting personal or small-scale projects. This starter kit provides a solid foundation with essential features like authentication, configuration management, database integration, and organized project structure.

🚀 Features
Gin Framework - High-performance HTTP web framework for Go

Structured Architecture - Clean separation of concerns with organized directories

Authentication System - Ready-to-use login API with JWT support

Environment Configuration - Easy configuration management using .env files

Database Integration - PostgreSQL/MySQL setup with migrations

Go Modules - Modern dependency management

Testing Setup - Pre-configured testing structure

Docker Support - Containerization-ready configuration

📁 Project Structure
text
golang-gin-boilerplate/
├── cmd/              # Application entry points
│   └── main.go      # Main application entry
├── config/           # Configuration management
│   └── config.go    # Configuration loader
├── db/               # Database migrations and setup
│   └── migrations/   # Database migration files
├── internal/         # Internal application code
│   ├── auth/         # Authentication logic
│   ├── handlers/     # HTTP request handlers
│   ├── middleware/   # Gin middleware
│   ├── models/       # Data models/structs
│   └── services/     # Business logic layer
├── tests/            # Test files
├── .env              # Environment variables
├── .gitignore        # Git ignore rules
├── go.mod            # Go module definition
├── go.sum            # Go dependencies checksum
└── README.md         # This file
🛠️ Getting Started
Prerequisites
Go 1.16 or higher

PostgreSQL/MySQL (or any preferred database)

Git

Installation
Clone the repository

bash
git clone https://github.com/wasitmirani/golang-gin-boilerplate.git
cd golang-gin-boilerplate
Set up environment variables

bash
cp .env.example .env
# Edit .env with your configuration
Install dependencies

bash
go mod tidy
Set up database

Configure your database connection in .env

Run database migrations from the db/migrations/ directory

Run the application

bash
go run cmd/main.go
Quick Start for Personal Projects
For personal web services, simply:

Update the .env file with your database credentials

Modify authentication settings in internal/auth/ if needed

Add your business logic in internal/services/

Create new API endpoints in internal/handlers/

🏗️ Architecture Overview
This boilerplate follows a layered architecture:

Handlers - Handle HTTP requests/responses

Services - Implement business logic

Repositories - Manage data access

Models - Define data structures

Middleware - Cross-cutting concerns (auth, logging, etc.)

🔧 Configuration
Edit the .env file to configure:

Database connection (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

Server settings (PORT, HOST)

JWT secrets for authentication

Application environment (development/production)

📦 Available Endpoints
Authentication
POST /api/auth/login - User authentication with JWT token generation

🧪 Testing
Run tests with:

bash
go test ./tests/...
🐳 Docker Support
Build and run with Docker:

bash
docker build -t gin-boilerplate .
docker run -p 8080:8080 --env-file .env gin-boilerplate
📝 Customization for Your Needs
Adding New Features
New API Endpoints: Create handlers in internal/handlers/

New Business Logic: Add services in internal/services/

New Data Models: Define in internal/models/

New Database Tables: Create migration files in db/migrations/

Personal Use Tips
Replace placeholder JWT secrets with secure random strings

Adjust authentication logic in internal/auth/ for your requirements

Extend the configuration in config/config.go for additional settings

Add your custom middleware in internal/middleware/

🤝 Contributing
While this is primarily for personal use, improvements are welcome:

Fork the repository

Create a feature branch

Commit your changes

Push to the branch

Open a Pull Request

📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

👤 Author
Wasit Mirani - [GitHub Profile](https://github.com/wasitmirani)

🙏 Acknowledgments
Gin Web Framework

Go community for excellent libraries and tools

