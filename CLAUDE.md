# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this service does

`pm-gateway` is an API Gateway built with Go/Gin. It exposes REST endpoints and routes traffic to backend microservices via gRPC. It handles JWT validation centrally so downstream services don't need to.

## Commands

```bash
go run ./cmd/main.go          # run the gateway
go build -o pm-gateway ./cmd/main.go  # build binary
go test ./...                 # run tests (none exist yet)
```

Environment is loaded from `.env` via `godotenv`. Copy the existing `.env` and adjust values before running.

## Architecture

```
HTTP client
    ↓
Gin router (cmd/main.go)
    ├─ Middleware: CORS → Logger → JWT auth (on protected routes)
    ├─ Public:    POST /auth/login, POST /auth/register
    └─ Protected: GET  /api/users, GET /api/users/:id
        ↓
gRPC clients (internal/grpc/)
    ├─ Auth Service  @ AUTH_SERVICE_ADDR (default :50051)
    └─ User Service  @ USER_SERVICE_ADDR (default :50052)  ← stubbed
```

**Packages:**
- `cmd/main.go` — router setup, middleware wiring, route registration
- `config/config.go` — reads env vars; provides `Config` struct with `Port`, `JWTSecret`, `AuthServiceAddr`, `UserServiceAddr`
- `internal/handlers/auth.go` — Login (implemented), Register (stub)
- `internal/handlers/users.go` — GetUsers / GetUser (stubs pending User Service gRPC)
- `internal/middleware/auth.go` — JWT validation; injects claims into Gin context
- `internal/grpc/auth_client.go` — thin gRPC wrapper around `AuthServiceClient`

## Proto / gRPC

Proto definitions live in the external module `github.com/pav-dev98/pm-proto` (repo: `pav-dev98/pm-proto`). Generated types are imported directly — do not add `.proto` files to this repo. To update proto, bump the module version in `go.mod` and run `go mod tidy`.

## Environment variables

| Variable | Default | Notes |
|---|---|---|
| `PORT` | `8080` | HTTP listen port |
| `JWT_SECRET` | `secret` | Change before any deployment |
| `AUTH_SERVICE_ADDR` | `localhost:50051` | gRPC address for Auth Service |
| `USER_SERVICE_ADDR` | `localhost:50052` | gRPC address for User Service |

## Current status

- Auth login flow is fully implemented end-to-end (HTTP → gRPC → Auth Service → JWT tokens returned).
- Register, GetUsers, and GetUser are stubs returning placeholder strings — User Service gRPC integration is not yet done.
- gRPC connections use insecure credentials (no TLS) — intentional for local dev.
- No tests exist yet.
