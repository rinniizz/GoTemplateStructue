@echo off
echo 🔥 Development Mode (Hot Reload)

REM Check if Air is installed
air -v >nul 2>&1
if %errorlevel% neq 0 (
    echo 📦 Installing Air...
    go install github.com/cosmtrek/air@latest
    
    if %errorlevel% neq 0 (
        echo ❌ Air installation failed
        goto MANUAL_MODE
    )
)

REM Check if .air.toml exists, if not create it
if not exist ".air.toml" (
    echo 📝 Creating .air.toml configuration...
    call :CREATE_AIR_CONFIG
)

echo 📝 Generating Swagger docs...
go run github.com/swaggo/swag/cmd/swag@latest init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false

if %errorlevel% equ 0 (
    echo ✅ Swagger docs generated!
) else (
    echo ⚠️ Swagger generation failed, continuing anyway...
)

echo 🚀 Starting server with Hot Reload...
echo 📍 Server: http://localhost:8080
echo 📚 Swagger UI: http://localhost:8080/swagger/index.html
echo 🏥 Health Check: http://localhost:8080/health
echo.
echo 🔄 Auto-restart enabled - changes will trigger rebuild
echo 💡 Press Ctrl+C to stop
echo.

REM Set development environment
set GIN_MODE=debug
set SERVER_PORT=8080
set LOG_LEVEL=debug

air -c .air.toml
goto END

:MANUAL_MODE
echo.
echo 🔧 Manual Development Mode
echo 💡 You'll need to restart manually when files change
echo.
echo 📍 Server: http://localhost:8080
echo 📚 Swagger UI: http://localhost:8080/swagger/index.html
echo.

set GIN_MODE=debug
set SERVER_PORT=8080
set LOG_LEVEL=debug

go run cmd/server/main.go
goto END

:CREATE_AIR_CONFIG
echo # Hot reload configuration for Air > .air.toml
echo root = "." >> .air.toml
echo testdata_dir = "testdata" >> .air.toml
echo tmp_dir = "tmp" >> .air.toml
echo. >> .air.toml
echo [build] >> .air.toml
echo   args_bin = [] >> .air.toml
echo   bin = "./tmp/main.exe" >> .air.toml
echo   cmd = "go build -o ./tmp/main.exe ./cmd/server/main.go" >> .air.toml
echo   delay = 1000 >> .air.toml
echo   exclude_dir = ["assets", "tmp", "vendor", "testdata", "docs", ".git"] >> .air.toml
echo   exclude_file = [] >> .air.toml
echo   exclude_regex = ["_test.go"] >> .air.toml
echo   exclude_unchanged = false >> .air.toml
echo   follow_symlink = false >> .air.toml
echo   full_bin = "" >> .air.toml
echo   include_dir = [] >> .air.toml
echo   include_ext = ["go", "tpl", "tmpl", "html"] >> .air.toml
echo   include_file = [] >> .air.toml
echo   kill_delay = "0s" >> .air.toml
echo   log = "build-errors.log" >> .air.toml
echo   poll = false >> .air.toml
echo   poll_interval = 0 >> .air.toml
echo   rerun = false >> .air.toml
echo   rerun_delay = 500 >> .air.toml
echo   send_interrupt = false >> .air.toml
echo   stop_on_root = false >> .air.toml
echo. >> .air.toml
echo [color] >> .air.toml
echo   app = "" >> .air.toml
echo   build = "yellow" >> .air.toml
echo   main = "magenta" >> .air.toml
echo   runner = "green" >> .air.toml
echo   watcher = "cyan" >> .air.toml
echo. >> .air.toml
echo [log] >> .air.toml
echo   main_only = false >> .air.toml
echo   time = false >> .air.toml
echo. >> .air.toml
echo [misc] >> .air.toml
echo   clean_on_exit = false >> .air.toml
echo. >> .air.toml
echo [screen] >> .air.toml
echo   clear_on_rebuild = false >> .air.toml
echo   keep_scroll = true >> .air.toml
goto :eof

:END
pause
