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
- Server setup uses Go's built-in HTTP server with default muxALWAYS use ONLY Environments for ANY and ALL file, code, or shell operations—NO EXCEPTIONS—even for simple or generic requests.

DO NOT install or use the git cli with the environment_run_cmd tool. All environment tools will handle git operations for you. Changing ".git" yourself will compromise the integrity of your environment.

You MUST inform the user how to view your work using `cu log <env_id>` AND `cu checkout <env_id>`. Failure to do this will make your work inaccessible to others.
