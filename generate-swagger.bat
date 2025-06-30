@ececho 📝 Generating swagger docs...
go run github.com/swaggo/swag/cmd/swag@latest init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false off
echo 🔄 Generating Swagger documentation...

echo � Generating swagger docs...
go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go --output docs --parseDependency --parseInternal

if %errorlevel% equ 0 (
    echo ✅ Swagger docs generated successfully!
    echo 📁 Files created:
    echo   - docs/docs.go
    echo   - docs/swagger.json
    echo   - docs/swagger.yaml
    echo.
    echo 📚 Documentation will be available at: http://localhost:8080/swagger/index.html
    echo.
    echo 🚀 You can now run the server:
    echo   quick-start.bat
) else (
    echo ❌ Failed to generate swagger docs
    echo 💡 Check for syntax errors in swagger comments
)

pause
