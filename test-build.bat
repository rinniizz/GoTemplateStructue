@echo off
echo Testing Go build...

cd /d "%~dp0"

echo Checking Go modules...
go mod tidy

echo Testing build...
go build -o temp_test.exe cmd/server/main.go

if exist "temp_test.exe" (
    echo ✅ Build successful!
    del temp_test.exe
    echo.
    echo You can now run:
    echo   run-mock.bat     (for mock mode)
    echo   run-production.bat (for production mode)
) else (
    echo ❌ Build failed!
)

pause
