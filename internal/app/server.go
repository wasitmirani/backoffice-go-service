package app

import (
    "context"
    "net/http"
    "time"
    
    "github.com/gin-gonic/gin"
    
    "project/internal/api/v1"
    "project/internal/config"
    "project/internal/pkg/logger"
    "project/internal/services"
)

type Application struct {
    config *config.Config
    logger logger.Logger
    server *http.Server
    router *gin.Engine
    
    // Services
    userService    *services.UserService
    authService    *services.AuthService
    productService *services.ProductService
}

func New(cfg *config.Config, log logger.Logger) (*Application, error) {
    // Set Gin mode
    gin.SetMode(cfg.Server.Mode)
    
    // Create router
    router := gin.New()
    router.Use(gin.Recovery())
    
    app := &Application{
        config: cfg,
        logger: log,
        router: router,
    }
    
    // Initialize dependencies
    if err := app.initDependencies(); err != nil {
        return nil, err
    }
    
    // Setup routes
    app.setupRoutes()
    
    // Create HTTP server
    app.server = &http.Server{
        Addr:         ":" + cfg.Server.Port,
        Handler:      router,
        ReadTimeout:  cfg.Server.ReadTimeout,
        WriteTimeout: cfg.Server.WriteTimeout,
        IdleTimeout:  cfg.Server.IdleTimeout,
    }
    
    return app, nil
}

func (app *Application) initDependencies() error {
    // Initialize database connection
    // Initialize repositories
    // Initialize services
    // Return any error
    return nil
}

func (app *Application) setupRoutes() {
    // Health check
    app.router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "healthy"})
    })
    
    // API v1 routes
    v1.SetupRoutes(app.router.Group("/api/v1"), app.userService, app.authService)
}

func (app *Application) Start() error {
    app.logger.Info("Starting server", "port", app.config.Server.Port)
    return app.server.ListenAndServe()
}

func (app *Application) Shutdown(ctx context.Context) error {
    app.logger.Info("Shutting down server...")
    return app.server.Shutdown(ctx)
}