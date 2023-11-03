package main

import (
	"awesomeProject/finbox_project/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes
	r.GET("/api/live/:symbol", controllers.LiveTickerSimulation)

	r.Run(":8080")
}
