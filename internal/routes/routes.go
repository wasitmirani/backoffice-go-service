package routes

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all application routes
// This function can be used to organize routes by domain/feature
func SetupRoutes(router *gin.Engine) {
	// Health check routes are handled in app.go
	
	// API v1 routes
	api := router.Group("/api/v1")
	{
		// Auth routes
		setupAuthRoutes(api)
		
		// User routes
		setupUserRoutes(api)
		
		// Add more route groups here
	}
}

// setupAuthRoutes sets up authentication routes
func setupAuthRoutes(api *gin.RouterGroup) {
	_ = api.Group("/auth")
	// {
	// 	auth.POST("/login", authHandler.Login)
	// 	auth.POST("/register", authHandler.Register)
	// 	auth.POST("/refresh", authHandler.RefreshToken)
	// 	auth.POST("/logout", authHandler.Logout)
	// }
}

// setupUserRoutes sets up user management routes
func setupUserRoutes(api *gin.RouterGroup) {
	_ = api.Group("/users")
	// {
	// 	users.GET("", userHandler.List)
	// 	users.GET("/:id", userHandler.Get)
	// 	users.POST("", userHandler.Create)
	// 	users.PUT("/:id", userHandler.Update)
	// 	users.DELETE("/:id", userHandler.Delete)
	// }
}
