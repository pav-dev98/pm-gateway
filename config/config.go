package config

import (
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    Port        string
    JWTSecret   string
    AuthService string  // dirección gRPC ej: localhost:50051
    UserService string  // dirección gRPC ej: localhost:50052
}

func Load() *Config { // devuelve un puntero a un Config
    godotenv.Load() // carga el .env

    return &Config{
        Port:        getEnv("PORT", "8080"),
        JWTSecret:   getEnv("JWT_SECRET", "secret"),
        AuthService: getEnv("AUTH_SERVICE_ADDR", "localhost:50051"),
        UserService: getEnv("USER_SERVICE_ADDR", "localhost:50052"),
    }
}

func getEnv(key, fallback string) string {
    if val := os.Getenv(key); val != "" {
        return val
    }
    return fallback
}