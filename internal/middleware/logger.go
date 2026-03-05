package middleware

import (
	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	// Gin ya tiene su propio logger, lo reutilizamos
	return gin.Logger()
}