package routes

import (
	"mini-order-api/config"
	"mini-order-api/handlers"
	"mini-order-api/repositories"
	"mini-order-api/services"

	"github.com/gin-gonic/gin"
)

// SetupRouter mengonfigurasi rute, middleware Gin, dan Dependency Injection
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Inisialisasi Dependency Injection: DB -> Repository -> Service -> Handler
	orderRepo := repositories.NewOrderRepository(config.DB)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Global Health Check
	r.GET("/health", handlers.HealthCheck)

	// API Grouping (v1)
	v1 := r.Group("/api/v1")
	{
		// Order Endpoints
		orders := v1.Group("/orders")
		{
			orders.POST("", orderHandler.CreateOrder)       // POST /api/v1/orders
			orders.GET("", orderHandler.GetAllOrders)       // GET  /api/v1/orders
			orders.GET("/:id", orderHandler.GetOrderByID)   // GET  /api/v1/orders/:id
		}

		// Webhook Endpoints
		webhooks := v1.Group("/webhooks")
		{
			webhooks.POST("/payment", orderHandler.PaymentWebhook) // POST /api/v1/webhooks/payment
		}
	}

	return r
}
