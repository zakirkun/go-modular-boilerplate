package user

import (
	"go-modular-boilerplate/internal/pkg/bus"
	"go-modular-boilerplate/internal/pkg/logger"
	"go-modular-boilerplate/modules/users/domain/entity"
	"go-modular-boilerplate/modules/users/domain/repository"
	"go-modular-boilerplate/modules/users/domain/service"
	"go-modular-boilerplate/modules/users/handler"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Module implements the application Module interface for the user module
type Module struct {
	db          *gorm.DB
	logger      *logger.Logger
	userService *service.UserService
	userHandler *handler.UserHandler
	event       *bus.EventBus
}

// Name returns the name of the module
func (m *Module) Name() string {
	return "user"
}

// Initialize initializes the module
func (m *Module) Initialize(db *gorm.DB, log *logger.Logger, event *bus.EventBus) error {
	m.db = db
	m.logger = log
	m.event = event

	m.logger.Info("Initializing user module")

	// Initialize repositories
	userRepo := repository.NewUserRepositoryImpl()
	m.logger.Debug("User repository initialized")

	// Initialize services
	m.userService = service.NewUserService(userRepo)
	m.logger.Debug("User service initialized")

	// Initialize handlers
	m.userHandler = handler.NewUserHandler(m.logger, m.event, m.userService)
	m.logger.Debug("User handler initialized")

	// register event listeners
	m.logger.Info("Registering user module event listeners")
	m.event.SubscribeFunc("user.created", m.userHandler.Handle)

	m.logger.Info("User module initialized successfully")
	return nil
}

// RegisterRoutes registers the module's routes
func (m *Module) RegisterRoutes(e *echo.Echo, basePath string) {
	m.logger.Info("Registering user routes at %s/users", basePath)
	m.userHandler.RegisterRoutes(e, basePath)
	m.logger.Debug("User routes registered successfully")
}

// Migrations returns the module's migrations
func (m *Module) Migrations() error {
	m.logger.Info("Registering user module migrations")
	return m.db.AutoMigrate(&entity.User{})
}

// Logger returns the module's logger
func (m *Module) Logger() *logger.Logger {
	return m.logger
}

// NewModule creates a new user module
func NewModule() *Module {
	return &Module{}
}
