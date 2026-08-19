package routes

import (
	"mini-order-api/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRouter mengonfigurasi rute dan middleware Gin
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Global Health Check
	r.GET("/health", handlers.HealthCheck)

	// API Grouping (v1)
	v1 := r.Group("/api/v1")
	{
		v1.GET("/orders", handlers.GetOrdersExample)
	}

	return r
}
