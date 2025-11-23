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
	@echo "🧑‍💻 Starting development mode..."
	@echo "📝 Generating Swagger docs..."
	@if command -v swag >/dev/null 2>&1; then \
		swag init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false; \
		echo "✅ Swagger docs generated!"; \
	else \
		echo "⚠️ swag not found, installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		swag init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false; \
		echo "✅ Swagger docs generated!"; \
	fi
	@if command -v air >/dev/null 2>&1; then \
		echo "🔥 Using Air for hot reload..."; \
		if [ ! -f .air.toml ]; then \
			echo "📝 Creating .air.toml..."; \
			echo 'root = "."' > .air.toml; \
			echo 'testdata_dir = "testdata"' >> .air.toml; \
			echo 'tmp_dir = "tmp"' >> .air.toml; \
			echo '' >> .air.toml; \
			echo '[build]' >> .air.toml; \
			echo '  args_bin = []' >> .air.toml; \
			echo '  bin = "./tmp/main"' >> .air.toml; \
			echo '  cmd = "go build -o ./tmp/main ./cmd/server/main.go"' >> .air.toml; \
			echo '  delay = 1000' >> .air.toml; \
			echo '  exclude_dir = ["assets", "tmp", "vendor", "testdata", "docs", ".git"]' >> .air.toml; \
			echo '  exclude_file = []' >> .air.toml; \
			echo '  exclude_regex = ["_test.go"]' >> .air.toml; \
			echo '  exclude_unchanged = false' >> .air.toml; \
			echo '  follow_symlink = false' >> .air.toml; \
			echo '  full_bin = ""' >> .air.toml; \
			echo '  include_dir = []' >> .air.toml; \
			echo '  include_ext = ["go", "tpl", "tmpl", "html"]' >> .air.toml; \
			echo '  include_file = []' >> .air.toml; \
			echo '  kill_delay = "0s"' >> .air.toml; \
			echo '  log = "build-errors.log"' >> .air.toml; \
			echo '  poll = false' >> .air.toml; \
			echo '  poll_interval = 0' >> .air.toml; \
			echo '  rerun = false' >> .air.toml; \
			echo '  rerun_delay = 500' >> .air.toml; \
			echo '  send_interrupt = false' >> .air.toml; \
			echo '  stop_on_root = false' >> .air.toml; \
		fi; \
		echo "🚀 Starting server with Hot Reload..."; \
		echo "📍 Server: http://localhost:8080"; \
		echo "📚 Swagger UI: http://localhost:8080/swagger/index.html"; \
		echo "🏥 Health Check: http://localhost:8080/health"; \
		echo "🔄 Auto-restart enabled - changes will trigger rebuild"; \
		echo "💡 Press Ctrl+C to stop"; \
		GIN_MODE=debug LOG_LEVEL=debug air; \
	elif command -v nodemon >/dev/null 2>&1; then \
		echo "🔄 Using nodemon for file watching..."; \
		echo "🚀 Starting server with File Watching..."; \
		echo "📍 Server: http://localhost:8080"; \
		echo "📚 Swagger UI: http://localhost:8080/swagger/index.html"; \
		echo "🏥 Health Check: http://localhost:8080/health"; \
		GIN_MODE=debug LOG_LEVEL=debug nodemon --exec "go run cmd/server/main.go" --ext go --watch cmd --watch internal --watch pkg --delay 2; \
	else \
		echo "⚠️  No hot reload tool found. Installing Air..."; \
		go install github.com/cosmtrek/air@latest; \
		echo "🔄 Please run 'make dev' again to use hot reload"; \
		echo "💡 Or install nodemon: npm install -g nodemon"; \
		echo "🔧 Starting in simple development mode..."; \
		echo "📍 Server: http://localhost:8080"; \
		echo "📚 Swagger UI: http://localhost:8080/swagger/index.html"; \
		echo "🏥 Health Check: http://localhost:8080/health"; \
		GIN_MODE=debug LOG_LEVEL=debug go run $(MAIN_PATH)/main.go; \
	fi

# Build the application
build:
	@echo "🔨 Building the application..."
	@echo "📝 Generating Swagger docs..."
	@if command -v swag >/dev/null 2>&1; then \
		swag init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false; \
		echo "✅ Swagger docs generated!"; \
	else \
		echo "⚠️ swag not found, installing..."; \
		go install github.com/swaggo/swag/cmd/swag@latest; \
		swag init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false; \
		echo "✅ Swagger docs generated!"; \
	fi
	@echo "🔨 Building binary..."
	go build -o bin/$(BINARY_NAME) $(MAIN_PATH)
	@echo "✅ Build completed: bin/$(BINARY_NAME)"

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
	swag init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false

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
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/gotemplate?sslmode=disable" up

migrate-down:
	@echo "⬇️ Reverting database migrations..."
	migrate -path migrations -database "postgresql://postgres:password@localhost:5432/gotemplate?sslmode=disable" down

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
