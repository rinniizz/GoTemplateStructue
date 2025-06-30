@echo off
echo 🧪 Testing Go Template API with Swagger...

REM Check if server is running
echo Checking if server is running...
curl -s http://localhost:8080/health >nul 2>&1
if %errorlevel% neq 0 (
    echo ❌ Server is not running!
    echo 💡 Start the server first:
    echo   run-mock.bat
    echo   or
    echo   run-with-swagger.bat
    pause
    exit /b 1
)

echo ✅ Server is running!
echo.

echo 🔐 Testing Authentication Endpoints...
echo.

echo 📝 1. Testing User Registration...
curl -X POST http://localhost:8080/api/v1/auth/register ^
  -H "Content-Type: application/json" ^
  -d "{\"username\":\"testuser\",\"email\":\"test@example.com\",\"password\":\"password123\",\"full_name\":\"Test User\"}"
echo.
echo.

echo 🔑 2. Testing User Login...
for /f %%i in ('curl -s -X POST http://localhost:8080/api/v1/auth/login -H "Content-Type: application/json" -d "{\"email\":\"admin@example.com\",\"password\":\"password123\"}" ^| findstr "access_token"') do set TOKEN_LINE=%%i
echo Login response:
curl -X POST http://localhost:8080/api/v1/auth/login ^
  -H "Content-Type: application/json" ^
  -d "{\"email\":\"admin@example.com\",\"password\":\"password123\"}"
echo.
echo.

echo 👤 3. Testing Protected Endpoint (Get Profile)...
echo Note: Using demo token for mock mode
set DEMO_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJlbWFpbCI6ImFkbWluQGV4YW1wbGUuY29tIiwiZXhwIjo5OTk5OTk5OTk5fQ.demo_token_for_mock_mode
curl -X GET http://localhost:8080/api/v1/users/profile ^
  -H "Authorization: Bearer %DEMO_TOKEN%"
echo.
echo.

echo 📋 4. Testing Get All Users...
curl -X GET http://localhost:8080/api/v1/users ^
  -H "Authorization: Bearer %DEMO_TOKEN%"
echo.
echo.

echo 🏥 5. Testing Health Check...
curl -X GET http://localhost:8080/health
echo.
echo.

echo ✅ All tests completed!
echo.
echo 📚 You can also test these APIs using Swagger UI:
echo   http://localhost:8080/swagger/index.html
echo.
echo 💡 Tips:
echo   - All endpoints are auto-generated in Swagger
echo   - Use the "Authorize" button in Swagger UI to set Bearer token
echo   - Mock mode provides demo data for testing
echo.

pause
