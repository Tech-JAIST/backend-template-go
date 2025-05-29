# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Primary Commands
- **Build**: `go build -o ./$(basename $PWD)` (after `go mod download`)
- **Development**: `docker compose watch` - Starts hot-reload environment with API (localhost:8080), DB, and Adminer (localhost:8081)
- **Lint**: `golangci-lint run --fix ./...`
- **Format**: `golangci-lint fmt ./...`

### Docker-based Commands (if Go not installed locally)
- **Lint**: `docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v2.1.6-alpine golangci-lint run --fix ./...`
- **Format**: `docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v2.1.6-alpine golangci-lint fmt ./...`

### Task Runner (xc)
All above commands can be run using `xc` task runner (e.g., `xc dev`, `xc lint`). Install with: `go install github.com/joerdav/xc/cmd/xc@latest`

## Architecture Overview

This is a Go backend template using Echo web framework and GORM ORM with MySQL database.

### Core Stack
- **Web Framework**: Echo v4
- **ORM**: GORM with MySQL driver
- **Database**: MySQL with migrations via go-gormigrate
- **Development**: Docker Compose with hot-reload

### Directory Structure
- `main.go` - Entry point with Echo server setup, middleware, DB connection, and route registration
- `internal/config/` - Environment variable configuration for app, database, and GORM
- `internal/handler/` - HTTP handlers and routing (Echo-based)
- `internal/repository/` - Database operations abstraction layer
- `internal/db/model/` - GORM model definitions
- `internal/db/migration/` - Database migration management
- `internal/domain/` - Core business data types

### Key Patterns
- **Dependency Injection**: main.go → repository → handler chain
- **Repository Pattern**: Abstract database operations in `repository/` layer
- **Migration Strategy**: Use `InitialTables` for new projects, `Migrations` for schema changes in existing projects
- **Configuration**: Environment-based config with sensible defaults in `config/`

### API Implementation Flow
1. Add route in `internal/handler/handler.go`
2. Implement handler in `internal/handler/{group}.go`
3. Add repository methods in `internal/repository/{model}.go` if database operations needed
4. Call repository from handler

### Environment Variables
- `APP_ADDR` (default: `:8080`)
- `DB_*` variables for MySQL connection (host, port, user, pass, name)
- `SERVICE_NAME` (default: `app`)