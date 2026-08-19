package routes

import (
	"mini-order-api/config"
	"mini-order-api/handlers"
	"mini-order-api/repositories"
	"mini-order-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware mengizinkan akses request dari aplikasi frontend web manapun
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// SetupRouter mengonfigurasi rute, middleware Gin, dan Dependency Injection
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Pasang CORS Middleware
	r.Use(CORSMiddleware())

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
