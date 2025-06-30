@ececho 📝 Step 1: Generating swagger docs from code...
go run github.com/swaggo/swag/cmd/swag@latest init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false off
echo 🚀 Go Template with Auto-Generated Swagger...

echo � Step 1: Generating swagger docs from code...
go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go --output docs --parseDependency --parseInternal

if %errorlevel% equ 0 (
    echo ✅ Swagger docs generated successfully!
    echo 📁 Generated files:
    echo   - docs/docs.go
    echo   - docs/swagger.json
    echo   - docs/swagger.yaml
) else (
    echo ⚠️  Swagger generation failed, but will continue...
)

echo 🔧 Step 2: Installing dependencies...
go mod tidy

echo 🧪 Step 3: Testing build...
go build -o temp_test.exe cmd/server/main.go

if exist "temp_test.exe" (
    echo ✅ Build successful!
    del temp_test.exe
    
    echo.
    echo 🚀 Step 4: Starting server in Mock Mode...
    set MOCK_MODE=true
    set GIN_MODE=debug
    set SERVER_PORT=8080
    set LOG_LEVEL=info
    
    echo.
    echo 🌟 Access URLs:
    echo   📍 Server: http://localhost:8080
    echo   📚 Swagger UI: http://localhost:8080/swagger/index.html
    echo   🏥 Health Check: http://localhost:8080/health
    echo.
    echo 💡 Tips:
    echo   - API endpoints are auto-generated from code comments
    echo   - Use Mock Mode - no database required
    echo   - Authentication endpoints work with demo data
    echo   - Press Ctrl+C to stop the server
    echo.
    echo 📚 Swagger UI: http://localhost:8080/swagger/index.html
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
