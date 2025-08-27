package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nisargrbhatt/rest-ex/configs"
	"github.com/nisargrbhatt/rest-ex/routes"
)

func main() {
	r := gin.Default()
	configs.ConnectDB()

	routes.HealthRoute(r)

	r.Run("localhost:8080")

}
