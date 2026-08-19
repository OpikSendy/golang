package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck mengembalikan status server
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "Mini Order Management API is running healthy! 🚀",
	})
}
