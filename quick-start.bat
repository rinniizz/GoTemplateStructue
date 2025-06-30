@echo off
echo 🔄 Quick Test & Run...

echo 📋 Testing build...
go build -o temp_test.exe cmd/server/main.go

if exist "temp_test.exe" (
    echo ✅ Build successful!
    del temp_test.exe
    
    echo.
    echo 🚀 Starting in Mock Mode...
    set MOCK_MODE=true
    set SERVER_PORT=8080
    set LOG_LEVEL=info
    
    echo 📍 Server starting on http://localhost:8080
    echo 📚 API Docs: http://localhost:8080/swagger/index.html
    echo 🏥 Health Check: http://localhost:8080/health
    echo.
    echo ⏳ Starting... (Press Ctrl+C to stop)
    echo.
    
    go run cmd/server/main.go
) else (
    echo ❌ Build failed!
    echo Check the error messages above.
)

pause
