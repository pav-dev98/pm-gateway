package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/pav-dev98/pm-gateway/config"
	"github.com/pav-dev98/pm-gateway/internal/middleware"
	authpb "github.com/pav-dev98/pm-proto/auth"
)

func main() {
	cfg := config.Load()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	gwmux := runtime.NewServeMux()

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	if err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, gwmux, cfg.AuthService, dialOpts); err != nil {
		log.Fatalf("no se pudo registrar el Auth gateway: %v", err)
	}

	mux := http.NewServeMux()

	// tu API
	mux.Handle("/", gwmux)

	// sirve el json
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/api.swagger.json")
	})

	// sirve la UI
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger.json"),
	))

	handler := middleware.CORS(middleware.Logger(middleware.Auth(cfg)(mux)))

	log.Printf("Gateway corriendo en puerto %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, handler); err != nil {
		log.Fatalf("servidor HTTP terminó: %v", err)
	}
}
