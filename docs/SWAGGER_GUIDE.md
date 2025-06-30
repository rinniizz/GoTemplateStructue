# 📚 Swagger Documentation Guide

## 🎯 ภาพรวม

โปรเจกต์นี้ใช้ **Swagger/OpenAPI 2.0** ที่ generate อัตโนมัติจาก code comments โดยใช้ [swaggo/swag](https://github.com/swaggo/swag)

**ไม่ใช่** static files หรือ manual documentation แต่เป็น **auto-generated** จาก source code จริง

## 🔄 วิธีการ Generate Swagger Documentation

### Windows
```bash
# วิธีง่าย - ใช้ batch script
generate-swagger.bat

# วิธีใช้ go run (ไม่ต้องติดตั้ง swag CLI)
go run github.com/swaggo/swag/cmd/swag@latest init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false
```

### Linux/Mac
```bash
# ติดตั้ง swag CLI (ครั้งเดียว)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate documentation
swag init --generalInfo cmd/server/main.go --dir ./ --output docs --parseGoList=false

# หรือใช้ Makefile
make swagger
```

## 📂 ไฟล์ที่ถูกสร้าง

หลังจากรัน `swag init` จะได้ไฟล์:

```
docs/
├── docs.go         # Go package สำหรับ embed ใน application
├── swagger.json    # OpenAPI JSON schema
└── swagger.yaml    # OpenAPI YAML schema
```

## 🏷️ Swagger Comments Syntax

### ส่วน Main (ใน main.go)

```go
// @title Go Template API
// @version 1.0
// @description เทมเพลต API สำหรับ Golang พร้อม Best Practices
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
```

### ส่วน Handler Functions

```go
// Login godoc
// @Summary User login
// @Description Authenticate user and return tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body domain.LoginRequest true "User login credentials"
// @Success 200 {object} domain.APIResponse{data=domain.AuthResponse}
// @Failure 400 {object} domain.APIResponse
// @Failure 401 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
    // implementation...
}
```

### ส่วน Protected Endpoints

```go
// GetProfile godoc
// @Summary Get user profile
// @Description Get current user profile
// @Tags users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.APIResponse{data=domain.User}
// @Failure 401 {object} domain.APIResponse
// @Failure 500 {object} domain.APIResponse
// @Router /users/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
    // implementation...
}
```

## 🚀 Integration ใน Application

### 1. Import Generated Docs

```go
import (
    _ "go-template-structure/docs" // swagger docs
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)
```

### 2. Setup Route

```go
// Swagger documentation
router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
```

## 🌐 เข้าถึง Swagger UI

หลังจากรันเซิร์ฟเวอร์แล้ว:

- **Swagger UI:** http://localhost:8080/swagger/index.html
- **JSON Schema:** http://localhost:8080/swagger/doc.json
- **YAML Schema:** http://localhost:8080/swagger/swagger.yaml

## 🔧 การใช้งาน Authentication ใน Swagger UI

1. เปิด Swagger UI
2. กดปุ่ม **"Authorize"** ด้านบนขวา
3. ใส่ Bearer token: `Bearer your-jwt-token-here`
4. กด **"Authorize"**
5. ตอนนี้สามารถเรียก protected endpoints ได้

## 📋 Best Practices

### 1. Comments ที่ครบถ้วน
- ใส่ `@Summary` และ `@Description` ทุก endpoint
- ระบุ `@Tags` เพื่อจัดกลุ่ม API
- ใส่ `@Param` สำหรับ parameters ทั้งหมด
- ระบุ response ทุก status code

### 2. Data Models
- สร้าง struct ใน `domain` package
- ใช้ JSON tags สำหรับ field names
- ใส่ validation tags ด้วย

```go
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}
```

### 3. Consistent Response Format
- ใช้ `domain.APIResponse` เป็น wrapper ทุก response
- ใส่ generic type สำหรับ data: `{data=domain.User}`

## 🔍 Troubleshooting

### Swagger UI ไม่แสดง APIs
1. ตรวจสอบว่า docs ถูก generate แล้ว
2. ตรวจสอบ import `_ "go-template-structure/docs"`
3. ตรวจสอบ swagger comments syntax

### Generate ล้มเหลว
1. ตรวจสอบ Go syntax ใน files
2. ตรวจสอบ swagger comments format
3. ใช้ `--parseDependency --parseInternal` flags

### Empty Documentation
1. ตรวจสอบว่ามี swagger comments ใน handlers
2. ตรวจสอบ router paths ตรงกับ `@Router` comments
3. ลอง rebuild และ restart server

## 📱 ตัวอย่างการใช้งาน

### 1. Register User
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123",
    "full_name": "Test User"
  }'
```

### 2. Login
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 3. Use Protected Endpoint
```bash
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer YOUR_JWT_TOKEN"
```

## 🔄 Auto-Regeneration Workflow

เมื่อแก้ไข API:

1. แก้ไข handler functions
2. อัปเดต swagger comments
3. รัน `generate-swagger.bat` (Windows) หรือ `make swagger` (Linux/Mac)
4. Restart server
5. Refresh Swagger UI

**หมายเหตุ:** ในโหมด development สามารถใช้ hot reload tools เช่น Air เพื่อ auto-restart และ auto-regenerate docs
