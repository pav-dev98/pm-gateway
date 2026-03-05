package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/pav-dev98/pm-gateway/config"
)

func GetUsers(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Aquí luego llamaremos al Users Service via gRPC
		c.JSON(http.StatusOK, gin.H{
			"message": "get users - proximamente gRPC",
		})
	}
}

func GetUser(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "get user by id - proximamente gRPC",
			"id":      id,
		})
	}
}