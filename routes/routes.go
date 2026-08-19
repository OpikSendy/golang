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
		orders := v1.Group("/orders")
		{
			orders.POST("", handlers.CreateOrder)       // POST /api/v1/orders
			orders.GET("", handlers.GetAllOrders)       // GET  /api/v1/orders
			orders.GET("/:id", handlers.GetOrderByID)   // GET  /api/v1/orders/:id
		}
	}

	return r
}
