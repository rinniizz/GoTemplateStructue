# 🧑‍💻 Development Guide

## 🎯 ภาพรวม

คู่มือนี้จะแนะนำวิธีการพัฒนาโปรเจกต์ Go Template Structure ในสภาพแวดล้อม Development พร้อม Auto Swagger Generation

## � Quick Start Development

### Windows
```bash
dev.bat
```

### Linux/Mac
```bash
make dev
```

**Features ที่ได้:**
- �🔥 Hot reload (Air)
- 📚 Auto-generate Swagger docs
-  Debug logging
- �️ ใช้ PostgreSQL/Redis จริง
- �📍 Server: http://localhost:8080
- 📖 Swagger UI: http://localhost:8080/swagger/index.html

## 📚 Auto Swagger Generation

### การทำงาน

โปรเจกต์นี้จะ **generate Swagger docs อัตโนมัติ** ทุกครั้งที่:
- รัน `dev.bat` (Windows)
- รัน `make dev` (Linux/Mac)
- รัน `build.bat` (Windows)
- รัน `make build` (Linux/Mac)

### ไฟล์ที่ Generate

```
docs/
├── docs.go         # Go embed file
├── swagger.json    # Swagger JSON spec
└── swagger.yaml    # Swagger YAML spec
```

### เพิ่ม API Documentation

เพิ่ม Swagger comments ใน handler functions:

```go
// GetUsers godoc
// @Summary      Get all users
// @Description  Get a list of all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {array}   User
// @Failure      500  {object}  ErrorResponse
// @Router       /api/v1/users [get]
func (h *Handler) GetUsers(c *gin.Context) {
    // implementation
}
```

## 🔥 Hot Reload Options

### Air (แนะนำ) - ออกแบบมาสำหรับ Go โดยเฉพาะ

**ข้อดี:**
- ⚡ รวดเร็วและเสถียร
- 🎯 ออกแบบมาสำหรับ Go โดยเฉพาะ
- 🔧 Configure ได้ละเอียด
- 📦 ไม่ต้องพึ่งภาษาอื่น (Node.js)
- 🛠️ ติดตั้งอัตโนมัติ
- 📚 Generate Swagger docs อัตโนมัติ

**Windows:**
```bash
dev.bat
```

**Linux/Mac:**
```bash
make dev
```

### Manual Restart (ทางเลือก)

สำหรับกรณีที่ไม่ต้องการ hot reload:

**Windows:**
```bash
go run cmd/server/main.go
```

**Linux/Mac:**
```bash
make run
```

## 🛠️ Development Environment

### Environment Variables

เมื่อรันใน Development Mode จะมีการตั้งค่าดังนี้:

```bash
GIN_MODE=debug          # Debug mode สำหรับ Gin
SERVER_PORT=8080        # Port ของเซิร์ฟเวอร์
LOG_LEVEL=debug         # Log level แบบละเอียด
```

### Features ใน Development Mode

1. **📝 Debug Logging**
   - แสดง request/response details
   - Error stack traces ละเอียด
   - Performance metrics

2. **🔄 CORS Enabled**
   - เรียก API จาก frontend development server ได้
   - รองรับ localhost หลาย port

3. **📚 Auto Swagger Generation**
   - Generate swagger docs ใหม่ทุกครั้งที่เริ่มต้น
   - อัปเดตทันทีเมื่อแก้ไข API comments

## 🚀 Development Workflow

### 1. เริ่มต้นพัฒนา
```bash
# Windows
run-dev.bat

# Linux/Mac  
make dev
```

### 2. แก้ไข Code
- แก้ไขไฟล์ .go ใดๆ
- เซิร์ฟเวอร์จะ restart อัตโนมัติ (ถ้าใช้ hot reload)

### 3. ทดสอบ API
```bash
# ใช้ Swagger UI
http://localhost:8080/swagger/index.html

# หรือใช้ curl script
test-api.bat  # Windows
```

### 4. แก้ไข Swagger Comments
```go
// UpdateUser godoc
// @Summary Update user information
// @Description Update user data by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body domain.UpdateUserRequest true "User data"
// @Success 200 {object} domain.APIResponse{data=domain.User}
// @Failure 400 {object} domain.APIResponse
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
    // implementation...
}
```

### 5. Generate Swagger Docs
```bash
# Manual generate
generate-swagger.bat  # Windows
make swagger         # Linux/Mac

# หรือ restart development server (จะ generate อัตโนมัติ)
```

## 🔍 Debugging

### 1. Debug ด้วย VS Code

เพิ่มในไฟล์ `.vscode/launch.json`:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug Go Server",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/cmd/server",
            "env": {
                "GIN_MODE": "debug",
                "LOG_LEVEL": "debug"
            },
            "cwd": "${workspaceFolder}",
            "args": []
        }
    ]
}
```

### 2. Debug ด้วย Delve
```bash
# ติดตั้ง delve
go install github.com/go-delve/delve/cmd/dlv@latest

# รัน debugger
dlv debug cmd/server/main.go
```

### 3. Logging ใน Development
```go
import "go-template-structure/pkg/logger"

// ใช้ logger ใน code
logger.Debug("Debug message", map[string]interface{}{
    "user_id": userID,
    "action": "update_profile",
})

logger.Error("Error occurred", map[string]interface{}{
    "error": err.Error(),
    "context": "user_service",
})
```

## 📝 Code Style และ Best Practices

### 1. Code Formatting
```bash
# Format code
go fmt ./...

# Import optimization
go mod tidy
```

### 2. Linting
```bash
# Windows (ถ้ามี golangci-lint)
golangci-lint run

# Linux/Mac
make lint
```

### 3. Testing
```bash
# รัน tests
go test ./...

# Test coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 🔧 Configuration Tips

### 1. .env.example
```bash
# Development environment
SERVER_HOST=localhost
SERVER_PORT=8080
GIN_MODE=debug
LOG_LEVEL=debug
JWT_SECRET=dev-secret-key

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=gotemplate_dev

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=
```

### 2. Git Hooks (Optional)
```bash
# สร้าง pre-commit hook
echo '#!/bin/sh
go fmt ./...
go test ./...
golangci-lint run' > .git/hooks/pre-commit

chmod +x .git/hooks/pre-commit
```

## 🚨 Troubleshooting

### 1. Air ไม่ทำงาน
```bash
# ลบ .air.toml และให้สร้างใหม่
del .air.toml
run-dev.bat
```

### 2. Port 8080 ถูกใช้แล้ว
```bash
# Windows
netstat -ano | findstr :8080
taskkill /PID [PID] /F

# หรือเปลี่ยน port
set SERVER_PORT=8081
```

### 3. Swagger ไม่อัปเดต
```bash
# Force regenerate
generate-swagger.bat
# แล้ว restart server
```

### 4. Go modules errors
```bash
go clean -modcache
go mod download
go mod tidy
```

## 💡 Pro Tips

1. **ใช้ VS Code Extensions:**
   - Go extension pack
   - REST Client สำหรับทดสอบ API
   - GitLens สำหรับ Git integration

2. **API Testing:**
   - ใช้ Swagger UI สำหรับ interactive testing
   - ใช้ Postman หรือ Insomnia สำหรับ automation
   - เขียน integration tests

3. **Performance Monitoring:**
   - ใช้ `pprof` สำหรับ profiling
   - Monitor memory usage
   - ดู response times ใน logs

4. **Security Testing:**
   - ทดสอบ JWT token expiration
   - ทดสอบ CORS policies
   - ทดสอบ input validation
