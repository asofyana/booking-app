# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go web application for room booking management built with the Gin framework. It uses SQLite for data storage and follows a layered architecture.

## Key Commands

- **Run**: `go run main.go`
- **Build**: `go build -o booking-app main.go`
- **Database**: `sqlite3 booking.db` (use `scripts/schema.sql` to set up schema)

## Architecture

**Layered Architecture:**
```
main.go → internal/app/app.go → (handlers → services → repositories → database)
```

- **Handlers** (`internal/handler/`): HTTP request/response logic. Routes are defined in `routes.go`.
- **Services** (`internal/services/`): Business logic. Each service wraps a repository interface.
- **Repositories** (`internal/repository/`): Data access layer implementing interfaces.
- **Entities** (`internal/entity/`): Data models (User, Room, Booking, Lookup).
- **Utils** (`internal/utils/`): Shared utilities (password hashing, translation, config).

**Dependency Injection Pattern:**
Services are constructed with repository dependencies in `app.go`, promoting testability.

## Key Files

| File | Purpose |
|------|---------|
| `internal/app/app.go` | App initialization, DI setup, router configuration |
| `internal/handler/routes.go` | Route registration with auth middleware |
| `internal/services/auth_service.go` | Session-based auth (`Auth`, `AdminAuth` middleware) |
| `internal/utils/security.go` | bcrypt password hashing/verification |
| `internal/utils/config.go` | YAML config loader (singleton) |
| `config.yaml` | Config: port, dbFile, language, defaultPassword hash |

## Authentication

- Session-based using `gorilla/sessions` with cookies
- Sessions expire after 15 minutes (MaxAge = 900)
- `Auth` middleware: requires any logged-in user
- `AdminAuth` middleware: requires admin role
- Users stored in `USERS` table with bcrypt-hashed passwords

## Database Schema

- `USERS`: User accounts with role/status
- `ROOM`: Room inventory
- `BOOKING`: Booking records with status tracking
- `BOOKING_ROOM`: Many-to-many mapping (booking ↔ rooms)
- `LOOKUP`: Key-value pairs for activity codes, user roles/status

## i18n

Uses `github.com/nicksnyder/go-i18n/v2`. Translations loaded from `locales/en-US.json` and `locales/id-ID.json`. Configured language via `config.yaml` (default: `id-ID`).

## Frontend

- Templates in `templates/` using Go's `html/template`
- Static assets in `static/` (Bootstrap, jQuery, datatables)
- Custom JS in `static/js/custom.js`
