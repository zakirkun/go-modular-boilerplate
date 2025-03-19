package app

import (
	"go-modular-boilerplate/internal/pkg/bus"
	"go-modular-boilerplate/internal/pkg/logger"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Module represents an application module
type Module interface {
	// Name returns the name of the module
	Name() string

	// Initialize initializes the module
	Initialize(db *gorm.DB, logger *logger.Logger, event *bus.EventBus) error

	// RegisterRoutes registers the module's routes
	RegisterRoutes(e *echo.Echo, group string)

	// Migrations returns the module's database migrations
	Migrations() error

	// Logger returns the module's logger
	Logger() *logger.Logger
}
