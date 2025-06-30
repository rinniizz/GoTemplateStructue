.PHONY: run build test clean lint swagger

# Variables
BINARY_NAME=gotemplate
MAIN_PATH=./cmd/server

# Run the application
run:
	@echo "🚀 Starting the application..."
	go run $(MAIN_PATH)/main.go

# Run with hot reload (requires Air)
dev:
	@echo "🔥 Starting development server with hot reload..."
	air

# Run in mock mode (no database required)
mock:
	@echo "🔧 Starting in MOCK MODE - no database required..."
	MOCK_MODE=true go run $(MAIN_PATH)/main.go

# Run in mock mode for Windows
mock-win:
	@echo "🔧 Starting in MOCK MODE - no database required..."
	set MOCK_MODE=true && go run $(MAIN_PATH)/main.go

# Build the application
build:
	@echo "🔨 Building the application..."
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)

# Run tests
test:
	@echo "🧪 Running tests..."
	go test -v ./...

# Run tests with coverage
test-coverage:
	@echo "📊 Running tests with coverage..."
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Lint code
lint:
	@echo "🔍 Linting code..."
	golangci-lint run

# Generate swagger documentation
swagger:
	@echo "📚 Generating Swagger documentation..."
	swag init -g cmd/server/main.go -o api/swagger

# Clean build artifacts
clean:
	@echo "🧹 Cleaning up..."
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install development dependencies
dev-deps:
	@echo "📦 Installing development dependencies..."
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/cosmtrek/air@latest

# Download dependencies
deps:
	@echo "📥 Downloading dependencies..."
	go mod download
	go mod tidy

# Run database migrations (requires migrate tool)
migrate-up:
	@echo "⬆️ Running database migrations..."
	migrate -path scripts/migrations -database "postgresql://postgres:password@localhost:5432/gotemplate?sslmode=disable" up

migrate-down:
	@echo "⬇️ Reverting database migrations..."
	migrate -path scripts/migrations -database "postgresql://postgres:password@localhost:5432/gotemplate?sslmode=disable" down

# Docker commands
docker-build:
	@echo "🐳 Building Docker image..."
	docker build -t $(BINARY_NAME) .

docker-run:
	@echo "🐳 Running with Docker Compose..."
	docker-compose up -d

docker-stop:
	@echo "🛑 Stopping Docker containers..."
	docker-compose down

# Help
help:
	@echo "Available commands:"
	@echo "  run          - Run the application"
	@echo "  mock         - Run in mock mode (no database)"
	@echo "  dev          - Run with hot reload (Air)"
	@echo "  build        - Build the application"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  lint         - Lint the code"
	@echo "  swagger      - Generate Swagger docs"
	@echo "  clean        - Clean build artifacts"
	@echo "  dev-deps     - Install development dependencies"
	@echo "  deps         - Download and tidy dependencies"
	@echo "  docker-build - Build Docker image"
	@echo "  docker-run   - Run with Docker Compose"
	@echo "  docker-stop  - Stop Docker containers"
	@echo "  help         - Show this help message"
