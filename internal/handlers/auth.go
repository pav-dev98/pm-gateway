package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	grpcclients "github.com/pav-dev98/pmgateway/internal/grpc"
	"github.com/pav-dev98/pmgateway/config"
)

type LoginBody struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(cfg *config.Config) gin.HandlerFunc {
	// Creamos el cliente UNA sola vez, no en cada request
	authClient := grpcclients.NewAuthClient(cfg.AuthService)

	return func(c *gin.Context) {
		var body LoginBody

		// Validar que el body tenga email y password
		if err := c.ShouldBindJSON(&body); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Llamar al Auth service via gRPC
		res, err := authClient.Login(body.Email, body.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciales inválidas"})
			return
		}

		// Devolver los tokens al frontend
		c.JSON(http.StatusOK, gin.H{
			"access_token":  res.AccessToken,
			"refresh_token": res.RefreshToken,
			"token_type":    res.TokenType,
		})
	}
}

func Register(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "register endpoint - proximamente gRPC",
		})
	}
}