package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/pav-dev98/pm-gateway/config"
    "github.com/pav-dev98/pm-gateway/internal/handlers"
    "github.com/pav-dev98/pm-gateway/internal/middleware"
)

func main() {
    cfg := config.Load()
    r := gin.Default()

    // Middleware global
    r.Use(middleware.Logger())
    r.Use(middleware.CORS())

    // Rutas públicas
    auth := r.Group("/auth")
    {
        auth.POST("/login", handlers.Login(cfg))
        auth.POST("/register", handlers.Register(cfg))
    }

    // Rutas protegidas
    api := r.Group("/api")
    api.Use(middleware.Auth(cfg)) // valida JWT aquí
    {
        api.GET("/users", handlers.GetUsers(cfg))
        api.GET("/users/:id", handlers.GetUser(cfg))
    }

    log.Printf("Gateway corriendo en puerto %s", cfg.Port)
    r.Run(":" + cfg.Port)
}