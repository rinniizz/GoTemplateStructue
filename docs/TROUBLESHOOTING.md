# 🛠️ Troubleshooting Guide

## ❌ ปัญหาที่พบบ่อย

### 1. Build Error: "undefined: gin.CustomLogger"

**สาเหตุ:** Version ของ Gin ไม่รองรับ

**วิธีแก้:**
```bash
go mod tidy
go clean -modcache
go mod download
```

### 2. Port Already in Use

**Error:** `bind: address already in use`

**วิธีแก้:**
```bash
# Windows - หา process ที่ใช้ port 8080
netstat -ano | findstr :8080

# ปิด process
taskkill /PID [PID_NUMBER] /F

# หรือเปลี่ยน port ในไฟล์ .env
SERVER_PORT=8081
```

### 3. Database Connection Failed

**Error:** `Failed to connect to database`

**วิธีแก้:**
1. ใช้ Mock Mode แทน:
   ```bash
   set MOCK_MODE=true
   go run cmd/server/main.go
   ```

2. หรือติดตั้ง PostgreSQL และแก้ไขไฟล์ `.env`

### 4. Redis Connection Failed

**Error:** `Failed to connect to Redis`

**วิธีแก้:**
1. ใช้ Mock Mode (จะใช้ Memory Cache แทน)
2. หรือติดตั้ง Redis และแก้ไขไฟล์ `.env`

### 5. Go Module Error

**Error:** `module not found`

**วิธีแก้:**
```bash
go mod init go-template-structure
go mod tidy
go mod download
```

### 6. Permission Denied (Linux/Mac)

**Error:** `permission denied`

**วิธีแก้:**
```bash
chmod +x scripts/*.sh
sudo ./scripts/build.sh
```

### 7. Swagger UI Not Working

**Error:** `404 Not Found` เมื่อเข้า `/swagger/index.html`

**วิธีแก้:**
1. ตรวจสอบว่า swagger docs ถูก generate แล้ว:
   ```bash
   # Windows
   generate-swagger.bat
   
   # Linux/Mac
   swag init -g cmd/server/main.go --output docs
   ```

2. ตรวจสอบว่ามีไฟล์ `docs/docs.go`:
   ```bash
   ls docs/docs.go  # Linux/Mac
   dir docs\docs.go  # Windows
   ```

3. ตรวจสอบว่า import docs ใน main.go:
   ```go
   _ "go-template-structure/docs"
   ```

### 8. Swagger Generation Failed

**Error:** `swag: command not found` หรือ `'swag' is not recognized`

**วิธีแก้:**
1. ใช้ go run แทน swag CLI:
   ```bash
   go run github.com/swaggo/swag/cmd/swag@latest init -g cmd/server/main.go --output docs
   ```

2. หรือติดตั้ง swag CLI:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```

### 9. Empty Swagger Documentation

**Error:** Swagger UI แสดง แต่ไม่มี API endpoints

**วิธีแก้:**
1. ตรวจสอบ swagger comments ใน handlers:
   ```go
   // @Summary User login
   // @Description Authenticate user and return tokens
   // @Tags auth
   // @Router /auth/login [post]
   ```

2. ใช้ flags ครบถ้วนในการ generate:
   ```bash
   swag init -g cmd/server/main.go --output docs --parseDependency --parseInternal
   ```

## 🔍 การตรวจสอบ

### ตรวจสอบ Go Version
```bash
go version
# ต้องเป็น Go 1.21 หรือใหม่กว่า
```

### ตรวจสอบ Dependencies
```bash
go mod verify
go mod graph
```

### ตรวจสอบ Build
```bash
# Windows
test-build.bat

# Linux/Mac
go build -o temp cmd/server/main.go
```

### ตรวจสอบ Health Check
```bash
curl http://localhost:8080/health
```

## 📝 Log Files

### ตำแหน่งไฟล์ Log
- **Application Log:** Console output
- **Build Log:** `build-errors.log` (ถ้าใช้ Air)
- **Docker Log:** `docker-compose logs`

### ดู Logs
```bash
# Docker
docker-compose logs -f app

# Local
# Logs จะแสดงใน console
```

## 🆘 ขอความช่วยเหลือ

หากยังแก้ไขไม่ได้:

1. ✅ ตรวจสอบ Go version (ต้อง 1.21+)
2. ✅ รัน `go mod tidy`
3. ✅ ลองใช้ Mock Mode
4. ✅ ตรวจสอบไฟล์ `.env`
5. ✅ ดู error message ให้ละเอียด

### สร้าง Issue Report
เมื่อสร้าง issue ให้ใส่ข้อมูลนี้:
- OS และ version
- Go version
- Error message ครบถ้วน
- ขั้นตอนที่ทำมา
- ไฟล์ configuration (.env)
