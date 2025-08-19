# go-simple

**Company:** fern-sdk-sandbox  
**Language:** go  
**SDK Path:** ../../sdks/go

This is a sandbox project for testing the locally generated SDK for fern-sdk-sandbox.

## ⚠️ Important Note

**The initial generated code is not guaranteed to run without modifications.** You may need to:
- Update import statements to match the actual SDK structure
- Correct client class names and initialization parameters
- Adjust method calls and parameter names based on the SDK's API
- Set up proper environment variables and authentication

## Prerequisites

- Go 1.21 or later
- Make (optional, for using Makefile commands)

## Getting Started

1. **Download dependencies**:
   ```bash
   go mod download
   ```

2. **Run your sandbox**:
   ```bash
   go run main.go
   ```
   
   Or using Make:
   ```bash
   make run
   ```

3. **Build the binary** (optional):
   ```bash
   make build
   ./bin/sandbox
   ```

## Development Commands

- `make help` - Show available commands
- `make build` - Build the sandbox binary
- `make run` - Run the sandbox
- `make clean` - Clean build artifacts
- `make tidy` - Tidy go modules
- `make test` - Run tests
- `make fmt` - Format code
- `make vet` - Run go vet
- `make lint` - Run linter (requires golangci-lint)
- `make check` - Run all code quality checks

## Environment Variables

Create a `.env` file in this directory with your configuration:

```
API_KEY=your_api_key_here
```
