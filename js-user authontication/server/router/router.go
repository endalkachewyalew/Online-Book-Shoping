package router

import (
	"github.com/Endalk/Online-Book-Shoping/server/controller"
	"github.com/Endalk/Online-Book-Shoping/server/middlewares"

	"github.com/gin-gonic/gin"
)

// SetupRouter setup routing here
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Middlewares
	router.Use(middlewares.ErrorHandler)
	router.Use(middlewares.CORSMiddleware())

	// routes
	router.GET("/ping", controller.Pong)
	router.POST("/register", controller.Create)
	router.POST("/login", controller.Login)
	router.GET("/session", controller.Session)
	return router
}
