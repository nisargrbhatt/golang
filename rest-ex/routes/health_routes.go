package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nisargrbhatt/rest-ex/controllers"
)

func HealthRoute(router *gin.Engine) {
	router.GET("/health", controllers.CreateUser())
}
