@echo off
echo 🚀 Starting Go Template Structure in Mock Mode...
echo.

REM Check if Go is installed
go version >nul 2>&1
if errorlevel 1 (
    echo ❌ Go is not installed or not in PATH
    echo Please install Go 1.21+ from https://golang.org/dl/
    pause
    exit /b 1
)

REM Set environment variables
set MOCK_MODE=true
set SERVER_PORT=8080
set LOG_LEVEL=info
set JWT_SECRET=your-super-secret-jwt-key

echo 🔧 Running in MOCK MODE - no database required
echo 📍 Server will start on http://localhost:8080
echo 📚 Swagger UI: http://localhost:8080/swagger/index.html
echo 🏥 Health Check: http://localhost:8080/health
echo.
echo ⏳ Starting server... (Press Ctrl+C to stop)
echo.

REM Run the application
go run cmd/server/main.go

if errorlevel 1 (
    echo.
    echo ❌ Failed to start server!
    echo 💡 Try running: test-build.bat
    pause
)

pause
