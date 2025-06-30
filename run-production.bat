@echo off
echo 🚀 Starting Go Template Structure with Database...
echo.

REM Set environment variables for production mode
set MOCK_MODE=false
set SERVER_PORT=8080
set LOG_LEVEL=info
set JWT_SECRET=your-super-secret-jwt-key

echo 🔌 Running in PRODUCTION MODE - requires database
echo 📍 Server will start on http://localhost:8080
echo 📚 Swagger UI: http://localhost:8080/swagger/index.html
echo 🏥 Health Check: http://localhost:8080/health
echo.
echo ⚠️  Make sure PostgreSQL and Redis are running!
echo.

REM Run the application
go run cmd/server/main.go

pause
