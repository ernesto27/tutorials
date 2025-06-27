# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview
This is a simple Go HTTP API server that provides a "Hello World" JSON endpoint. The project consists of a single main.go file with minimal dependencies.

## Common Commands

### Development
- `go run main.go` - Run the server locally (listens on port 8080)
- `go build` - Build the executable
- `go mod tidy` - Clean up module dependencies

### Testing
- `go test ./...` - Run tests (no tests currently exist)
- `curl http://localhost:8080/hello` - Test the API endpoint when server is running

## Architecture
- **main.go**: Contains the entire application - HTTP handler, response struct, and server setup
- **go.mod**: Module definition with Go 1.23 requirement
- **Single endpoint**: `/hello` returns JSON response with "Hello World" message
- **Port**: Server runs on port 8080
- **No external dependencies**: Uses only Go standard library (net/http, encoding/json)

## Key Components
- `Response` struct: JSON response format with Message field
- `helloHandler`: HTTP handler function for the /hello endpoint
- Server setup uses Go's built-in HTTP server with default mux