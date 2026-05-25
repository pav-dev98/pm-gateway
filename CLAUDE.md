# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this service does

`pm-gateway` is an API Gateway built with Go and [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway). It exposes a REST/JSON surface and translates each request into a gRPC call against the corresponding backend microservice. JWT validation is enforced centrally so downstream services don't need to repeat it.

## Commands

```bash
go run ./cmd/main.go                  # run the gateway
go build -o pm-gateway ./cmd/main.go  # build binary
go test ./...                         # run tests (none exist yet)
```

Environment is loaded from `.env` via `godotenv`. Copy the existing `.env` and adjust values before running.

## Architecture

```
HTTP/JSON client (Next.js frontend)
    ↓
net/http server (cmd/main.go)
    ├─ Middleware chain: CORS → Logger → JWT auth (skips /v1/auth/*)
    └─ runtime.ServeMux (grpc-gateway)
        ↓ HTTP↔gRPC translation generated in pm-proto
        ├─ Auth Service  @ AUTH_SERVICE_ADDR (default :50051)
        │     POST /v1/auth/login
        │     POST /v1/auth/register
        └─ User Service  @ USER_SERVICE_ADDR (default :50052)  ← not registered yet
```

**Packages:**
- `cmd/main.go` — builds the grpc-gateway mux, registers service handlers, wraps with middleware, starts the HTTP server.
- `config/config.go` — reads env vars; provides `Config` struct with `Port`, `JWTSecret`, `AuthService`, `UserService`.
- `internal/middleware/cors.go` — CORS headers, handles preflight OPTIONS.
- `internal/middleware/logger.go` — request logger (method, path, duration).
- `internal/middleware/auth.go` — JWT validation. Paths under `/v1/auth/` are public; everything else requires a valid `Bearer` token signed with `JWT_SECRET`.

## Proto / gRPC

Proto definitions and generated code live in the external module `github.com/pav-dev98/pm-proto` (repo: `pav-dev98/pm-proto`). That module also publishes the grpc-gateway reverse-proxy code (`*.pb.gw.go`). HTTP routes are declared with `google.api.http` annotations in the `.proto` files — to change a route, edit the proto in `pm-proto`, regenerate, publish a new version, and bump the module here with `go get github.com/pav-dev98/pm-proto@latest && go mod tidy`.

To register a new service in the gateway, import its package and call its `RegisterXxxHandlerFromEndpoint` against the shared `runtime.ServeMux` in `cmd/main.go`.

## Environment variables

| Variable | Default | Notes |
|---|---|---|
| `PORT` | `8080` | HTTP listen port |
| `JWT_SECRET` | `secret` | Change before any deployment |
| `AUTH_SERVICE_ADDR` | `localhost:50051` | gRPC address for Auth Service |
| `USER_SERVICE_ADDR` | `localhost:50052` | gRPC address for User Service |

## Current status

- Auth service is wired through grpc-gateway: `POST /v1/auth/login` and `POST /v1/auth/register` proxy directly to the Auth gRPC service.
- User service handler is not registered yet — proto definitions in `pm-proto/users` are still empty.
- gRPC dial uses insecure credentials (no TLS) — intentional for local dev.
- No tests exist yet.
