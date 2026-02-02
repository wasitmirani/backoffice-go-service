package routes

import (
	"github.com/gin-gonic/gin"
	"backendapp/internal/app/controllers/auth"
	
)

func SetupRoutes(router *gin.Engine) {
	apiRoutes := router.Group("/api")
	{
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/login", gin.WrapF(auth.Login))
			authRoutes.POST("/register", gin.WrapF(auth.Register))
		}
	}

	// userRoutes := router.Group("/user")
	// {
	// 	userRoutes.GET("/:id", user.GetUser)
	// 	userRoutes.PUT("/:id", user.UpdateUser)
	// }
}