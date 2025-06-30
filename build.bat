@echo off
echo 🔨 Building Go Template...

echo 📝 Generating Swagger docs...
go run github.com/swaggo/swag/cmd/swag@latest init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false

if %errorlevel% neq 0 (
    echo ⚠️ Swagger generation failed, continuing build...
)

echo 🔧 Installing dependencies...
go mod tidy

echo 🧪 Testing build...
go build -o bin/go-template.exe cmd/server/main.go

if %errorlevel% equ 0 (
    echo ✅ Build successful!
    echo 📁 Binary created: bin/go-template.exe
    echo.
    echo 🚀 To run:
    echo   bin\go-template.exe
    echo.
    echo 💡 Or use: dev.bat (for development)
) else (
    echo ❌ Build failed!
    echo Check the error messages above.
    exit /b 1
)

pause
