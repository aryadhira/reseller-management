# Golang Fiber Template

A simple and clean template for starting new web projects with Go and Fiber.

## Prerequisites

- Go 1.18 or higher
- Git

## Getting Started

1. Clone this repository:
   ```bash
   git clone https://github.com/aryadhira/go-fiber-template.git
   cd go-fiber-template
   ```

2. Copy the environment variables:
   ```bash
   cp .env-example .env
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Run the application:
   ```bash
   go run cmd/app/main.go
   ```

The application will be available at `http://localhost:9099` (or the port specified in your .env file).

## Project Structure

- `cmd/app/main.go` - Application entry point
- `internal/config` - Application configuration
- `internal/database` - Database connections and configurations
- `internal/handlers` - HTTP request handlers
- `internal/interfaces` - Interfaces function
- `internal/models` - Data models
- `internal/routes` - API routes
- `internal/services` - Business logic
- `internal/utils` - Utility functions


