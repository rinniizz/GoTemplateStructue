#!/bin/bash

# Deployment script

set -e

echo "🚀 Starting deployment..."

# Build the application
echo "🔨 Building application..."
make build

# Run tests
echo "🧪 Running tests..."
make test

# Build Docker image
echo "🐳 Building Docker image..."
make docker-build

# Deploy with Docker Compose
echo "📦 Deploying with Docker Compose..."
docker-compose down
docker-compose up -d

# Wait for services to be ready
echo "⏳ Waiting for services to be ready..."
sleep 30

# Health check
echo "🏥 Performing health check..."
if curl -f http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ Deployment successful! Application is running."
    echo "📚 Swagger UI: http://localhost:8080/swagger/index.html"
else
    echo "❌ Health check failed!"
    docker-compose logs
    exit 1
fi
