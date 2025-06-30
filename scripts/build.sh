#!/bin/bash

# Build script for the Go application

echo "🔨 Building Go Template Structure..."

# Create bin directory if it doesn't exist
mkdir -p bin

# Set build variables
BINARY_NAME="gotemplate"
VERSION=$(git describe --tags --always --dirty)
BUILD_TIME=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT=$(git rev-parse HEAD)

# Build flags
LDFLAGS="-w -s -X main.Version=${VERSION} -X main.BuildTime=${BUILD_TIME} -X main.Commit=${COMMIT}"

# Build the application
echo "📦 Building binary..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="${LDFLAGS}" \
    -o bin/${BINARY_NAME} \
    ./cmd/server

if [ $? -eq 0 ]; then
    echo "✅ Build completed successfully!"
    echo "📁 Binary location: bin/${BINARY_NAME}"
    echo "📊 Binary size: $(du -h bin/${BINARY_NAME} | cut -f1)"
else
    echo "❌ Build failed!"
    exit 1
fi
