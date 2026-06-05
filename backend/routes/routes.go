package routes

import (
	"url-shortener/handlers"
	"url-shortener/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
	}

	api := r.Group("/api")
	api.Use(middleware.AuthRequired())
	{
		api.POST("/urls", handlers.CreateURL)
	}

	r.GET("/:shortCode", handlers.Redirect)
}
