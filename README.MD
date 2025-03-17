# Modular Go Application

A scalable, modular Golang application built with Echo framework, GORM ORM, and Domain-Driven Design principles. This project features a dynamic module registration system that allows adding new functionality without modifying the core application.

## Features

- **Modular Architecture**: Each feature is contained in its own module with clear boundaries
- **Dynamic Module Binding**: Modules are registered at runtime and loaded automatically
- **Domain-Driven Design**: Clean separation of domain, application, and infrastructure layers
- **RESTful API**: Built with Echo framework for high performance
- **Database Support**: MySQL integration with GORM
- **Docker Support**: Ready for containerized deployment
- **Comprehensive Logging**: Module-aware logging system

## Modules

The application currently includes the following modules for examples:

1. **User Module**: 
   - User management functionality
   - CRUD operations for user accounts


## Getting Started

### Prerequisites

- Go 1.20 or higher
- MySQL 8.0 or higher
- Docker and Docker Compose (for containerized setup)

### Running with Docker

The easiest way to run the application is using Docker Compose:

1. Clone the repository:
   ```bash
   git clone https://github.com/zakirkun/go-modular-boilerplate.git
   cd go-modular-boilerplate
   ```

2. Make the helper scripts executable:
   ```bash
   chmod +x run.sh cleanup.sh
   ```

3. Start the application:
   ```bash
   ./run.sh
   ```

4. The API will be available at http://localhost:8080

5. To stop the application:
   ```bash
   docker-compose down
   ```

### Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/zakirkun/go-modular-boilerplate.git
   cd go-modular-boilerplate
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Set up your environment variables (copy from .env.example):
   ```bash
    config.toml
   # Edit config.toml file with your local configuration
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

## API Endpoints

### User Module

- `GET /api/users`: Get all users
- `GET /api/users/:id`: Get a user by ID
- `POST /api/users`: Create a new user
- `PUT /api/users/:id`: Update a user
- `DELETE /api/users/:id`: Delete a user

## Configuration


### Logging
- `LOG_LEVEL`: Logging level (DEBUG, INFO, WARN, ERROR, OFF) (default: "INFO")

## Adding a New Module

To create a new module:

1. Create a new directory under `modules`
2. Implement the module interface defined in `internal/app/module.go`
3. Register the module in `main.go`

Example of minimal module implementation:

```go
package mymodule

import (
	"github.com/labstack/echo/v4"
	"go-modular-boilerplate/your-project/pkg/logger"
	"gorm.io/gorm"
)

type Module struct {
	db     *gorm.DB
	logger *logger.Logger
}

func (m *Module) Name() string {
	return "mymodule"
}

func (m *Module) Initialize(db *gorm.DB, log *logger.Logger) error {
	m.db = db
	m.logger = log
	m.logger.Info("My module initialized")
	return nil
}

func (m *Module) RegisterRoutes(e *echo.Echo, basePath string) {
	// Register your routes here
}

func (m *Module) Migrations() []interface{} {
	return []interface{}{
		// Your entity structs for migration
	}
}

func (m *Module) Logger() *logger.Logger {
	return m.logger
}

func NewModule() *Module {
	return &Module{}
}
```

## Docker Support

The application includes:

- `Dockerfile`: Multi-stage build for the Go application
- `docker-compose.yml`: Configuration for the app and MySQL
- `init.sql`: Database initialization script
- Helper scripts:
  - `run.sh`: Start the application with Docker Compose
  - `cleanup.sh`: Clean up Docker resources

## Logging

The application uses a custom logging system that:

- Supports multiple log levels (DEBUG, INFO, WARN, ERROR, OFF)
- Includes timestamps and module names in log entries
- Creates module-specific loggers
- Can be configured globally via environment variables

Example log output:
```
2025-03-17 14:30:05.123 [app] INFO: Registered module: user
2025-03-17 14:30:05.125 [user] INFO: Initializing user module
2025-03-17 14:30:05.130 [user] INFO: User module initialized successfully
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.