package app

import (
	"fmt"
	"go-modular-boilerplate/internal/pkg/bus"
	"go-modular-boilerplate/internal/pkg/config"
	"go-modular-boilerplate/internal/pkg/database"
	"go-modular-boilerplate/internal/pkg/logger"
	"go-modular-boilerplate/internal/pkg/server"
	_validator "go-modular-boilerplate/internal/pkg/validator"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

// App represents the application
type App struct {
	db      *gorm.DB
	server  *server.ServerContext
	modules []Module
	r       *echo.Echo
	logger  *logger.Logger
}

// NewApp creates a new application
func NewApp(cfg *logger.Config) (*App, error) {
	appLogger, err := logger.NewLogger(*cfg, config.GetString("server.app_name"))
	if err != nil {
		return nil, err
	}
	return &App{
		modules: make([]Module, 0),
		logger:  appLogger,
	}, nil
}

func (a *App) SetRouter() *echo.Echo {
	return echo.New()
}

// RegisterModule registers a module with the application
func (a *App) RegisterModule(module Module) {
	a.modules = append(a.modules, module)
	a.logger.Info("Registered module: %s", module.Name())
}

// Initialize initializes the application
func (a *App) Initialize() error {
	a.logger.Info("Initializing application...")

	// Initialize database
	var err *error
	a.db, err = a.SetDatabase().OpenDB()
	if err != nil {
		a.logger.Error("Failed to initialize database: %v", err)
		return *err
	}

	// Set database instance for all modules
	database.DB = a.db

	// Run migrations for all modules
	for _, module := range a.modules {
		migrations := module.Migrations()
		if len(migrations) > 0 {
			a.logger.Info("Running migrations for module: %s", module.Name())
			if err := a.db.AutoMigrate(migrations...); err != nil {
				a.logger.Error("Failed to run migrations for module %s: %v", module.Name(), err)
				return err
			}
			a.logger.Info("Migrations completed for module: %s", module.Name())
		}
	}

	// initialize router
	a.r = a.SetRouter()
	a.r.Use(middleware.Logger())
	a.r.Use(middleware.Recover())
	a.r.Use(middleware.CORS())

	// validate request
	a.r.Validator = _validator.NewCustomValidator()

	// Initialize modules
	for _, module := range a.modules {
		a.logger.Info("Initializing module: %s", module.Name())

		// Create module-specific logger
		moduleLogger := a.logger.WithPrefix(module.Name())
		if err := module.Initialize(a.db, moduleLogger); err != nil {
			a.logger.Error("Failed to initialize module %s: %v", module.Name(), err)
			return err
		}

		a.logger.Info("Module initialized: %s", module.Name())
	}

	// event bus initialization
	event := bus.NewEventBus()
	for _, module := range a.modules {
		module.RegisterEventDrivers(event)
	}

	// Initialize HTTP server
	a.server = a.SetServer()

	// api version
	version := fmt.Sprintf("/api/v%s", config.GetString("server.api_version"))

	// Register routes for all modules
	for _, module := range a.modules {
		a.logger.Info("Registering routes for module: %s", module.Name())
		module.RegisterRoutes(a.r, version)
		a.logger.Info("Routes registered for module: %s", module.Name())
	}

	// append handler to server
	a.server.Handler = a.r

	a.logger.Info("Application initialization completed")

	return nil
}

// Start starts the application
func (a *App) Start() {
	a.logger.Info("Starting server on %s", a.server.Host)
	a.server.Run()
}

// setup database model
func (a *App) SetDatabase() *database.DBModel {
	return &database.DBModel{
		ServerMode:   config.GetString("server.mode"),
		Driver:       config.GetString("database.db_driver"),
		Host:         config.GetString("database.db_host"),
		Port:         config.GetString("database.db_port"),
		Name:         config.GetString("database.db_name"),
		Username:     config.GetString("database.db_username"),
		Password:     config.GetString("database.db_password"),
		MaxIdleConn:  config.GetInt("pool.conn_idle"),
		MaxOpenConn:  config.GetInt("pool.conn_max"),
		ConnLifeTime: config.GetInt("pool.conn_lifetime"),
	}
}

// Setup Web Server
func (a *App) SetServer() *server.ServerContext {
	return &server.ServerContext{
		Host:         ":" + config.GetString("server.port"),
		ReadTimeout:  time.Duration(config.GetInt("server.http_timeout")),
		WriteTimeout: time.Duration(config.GetInt("server.http_timeout")),
	}
}
