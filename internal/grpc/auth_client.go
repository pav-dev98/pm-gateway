package grpc

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/pav-dev98/pm-proto/auth"
)

// AuthClient encapsula la conexión gRPC con el Auth service
type AuthClient struct {
	client pb.AuthServiceClient
}

// NewAuthClient crea la conexión al Auth service
func NewAuthClient(addr string) *AuthClient {
	conn, err := grpc.NewClient(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()), // sin TLS por ahora
	)
	if err != nil {
		log.Fatalf("no se pudo conectar al Auth service: %v", err)
	}

	return &AuthClient{
		client: pb.NewAuthServiceClient(conn),
	}
}

// Login llama al Auth service via gRPC
func (a *AuthClient) Login(email, password string) (*pb.LoginResponse, error) {
	// Timeout de 5 segundos para la llamada
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := a.client.Login(ctx, &pb.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}