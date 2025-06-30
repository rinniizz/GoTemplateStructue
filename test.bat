@echo off
echo 🧪 Testing Build...

echo 📦 Checking dependencies...
go mod tidy

echo 🔨 Testing build...
go build -o temp_test.exe cmd/server/main.go

if exist "temp_test.exe" (
    echo ✅ Build successful!
    del temp_test.exe
    echo.
    echo 🚀 Ready to run:
    echo   dev.bat   - Development mode
    echo   build.bat - Create production binary
) else (
    echo ❌ Build failed!
    echo Check the error messages above.
    exit /b 1
)

pause
