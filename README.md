# Go Template Structure

เทมเพลตโครงสร้าง Golang แบบ Clean Architecture ⚡ Quick Start

### 🔥 Development Mode (แนะนำ)

**รันครั้งเดียว ได้หมดเลย:**
```bash
# Windows
dev.bat

# Linux/Mac
make dev
```

**คุณสมบัติที่ได้:**
- 🔥 Hot reload อัตโนมัติเมื่อแก้ไขไฟล์ (Air)
- � Auto-generate Swagger docs ทุกครั้งที่รัน
- 🐛 Debug logging enabled
- 📍 Server: http://localhost:8080
- 📖 Swagger UI: http://localhost:8080/swagger/index.html
- 🏥 Health Check: http://localhost:8080/health

### 🧪 ทดสอบ Build

```bash
# Windows
test.bat

# Linux/Mac
make test
```

### 🚀 Production Build

```bash
# Windows (สร้าง bin/gotemplate.exe พร้อม swagger docs)
build.bat

# Linux/Mac (สร้าง bin/gotemplate พร้อม swagger docs)
make build
```

### 🔧 Commands Overview

#### Windows Batch Scripts
- `dev.bat` - 🔥 **Development Mode พร้อม Air Hot Reload + Auto Swagger (แนะนำ)**
- `build.bat` - 🔨 **สร้าง Production Binary พร้อม Auto Swagger**
- `test.bat` - 🧪 **ทดสอบ Build**

#### Makefile Commands (Linux/Mac)
- `make dev` - 🔥 Development Mode พร้อม Air Hot Reload + Auto Swagger
- `make build` - 🔨 สร้าง Production Binary พร้อม Auto Swagger
- `make test` - 🧪 รันเทสต์
- `make clean` - ลบไฟล์ build artifacts

## 🏗️ โครงสร้างโปรเจกต์

```
.
├── cmd/                    # Main applications
│   └── server/            # Server application
├── internal/              # Private application code
│   ├── config/           # Configuration
│   ├── domain/           # Business entities
│   ├── repository/       # Data layer
│   ├── service/          # Business logic
│   ├── handler/          # HTTP handlers
│   ├── middleware/       # HTTP middlewares
│   └── interfaces/       # Interface definitions
├── pkg/                   # Public libraries
│   ├── utils/            # Utility functions
│   ├── logger/           # Logging package
│   └── database/         # Database connections
├── docs/                  # Documentation
│   ├── docs.go           # Auto-generated Swagger
│   ├── swagger.json      # Swagger JSON schema
│   ├── swagger.yaml      # Swagger YAML schema
│   ├── SWAGGER_GUIDE.md  # Swagger documentation guide
│   ├── TROUBLESHOOTING.md # Troubleshooting guide
│   └── DEVELOPMENT.md    # Development guide
├── test/                  # Test files
├── bin/                   # Build output (ignored by git)
├── .env.example          # Environment variables example
├── Dockerfile            # Docker configuration
├── docker-compose.yml    # Docker compose
├── Makefile              # Build automation (Linux/Mac)
├── dev.bat               # Development mode (Windows)
├── build.bat             # Build binary (Windows)
├── test.bat              # Test build (Windows)
└── README.md             # This file
```

## 🚀 คุณสมบัติ

- ✅ Clean Architecture
- ✅ Dependency Injection
- ✅ Environment Configuration
- ✅ Database Integration (PostgreSQL)
- ✅ Redis Cache
- ✅ JWT Authentication
- ✅ Structured Logging
- ✅ API Documentation (Swagger)
- ✅ Unit Testing
- ✅ Docker Support
- ✅ Hot Reload Development (Air)
- ✅ Graceful Shutdown

## 📚 Auto Swagger Documentation

โปรเจกต์นี้ใช้ **auto-generated Swagger documentation** โดยจะ generate อัตโนมัติทุกครั้งที่:
- รัน `dev.bat` หรือ `make dev` (Development Mode)
- รัน `build.bat` หรือ `make build` (Production Build)

### 📖 เข้าถึง Swagger UI
- **URL:** http://localhost:8080/swagger/index.html
- **JSON:** http://localhost:8080/swagger/doc.json

### ➕ เพิ่ม API Documentation
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

## 🚀 Production Deployment

### Docker (แนะนำ)
```bash
docker-compose up -d
```

### Manual
```bash
# Copy environment file
cp .env.example .env

# Edit .env with your database settings
# Then run:
go run cmd/server/main.go
```

### 🧪 ทดสอบ API

เมื่อเซิร์ฟเวอร์รันแล้ว:

```bash
# ตรวจสอบสถานะ
curl http://localhost:8080/health

# เข้าสู่ระบบ
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"password123"}'
```

## �🛠️ การใช้งาน

### Prerequisites
- Go 1.21+
- PostgreSQL (สำหรับโหมดจริง)
- Redis (สำหรับโหมดจริง)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd go-template-structure
```

2. Install dependencies:
```bash
go mod download
```

3. Copy environment file:
```bash
cp .env.example .env
```

## 🛠️ การเริ่มต้นใช้งาน

### Prerequisites
- Go 1.21+
- PostgreSQL
- Redis

### Installation

1. Clone repository:
```bash
git clone <repository-url>
cd go-template-structure
```

2. Install dependencies:
```bash
go mod download
```

3. Setup environment:
```bash
cp .env.example .env
# แก้ไข .env ให้ตรงกับ database settings ของคุณ
```

4. รันโปรเจกต์:
```bash
# Development mode (แนะนำ)
dev.bat      # Windows
make dev     # Linux/Mac

# Production build
build.bat    # Windows
make build   # Linux/Mac
```

## 📖 API Documentation

เมื่อรันเซิร์ฟเวอร์แล้ว เข้าถึง API docs ได้ที่:
- **Swagger UI:** http://localhost:8080/swagger/index.html
- **Health Check:** http://localhost:8080/health

## 📚 เอกสารเพิ่มเติม

- [คู่มือพัฒนา (Development Guide)](docs/DEVELOPMENT.md)
- [คู่มือแก้ปัญหา (Troubleshooting)](docs/TROUBLESHOOTING.md)
- [คู่มือ Swagger](docs/SWAGGER_GUIDE.md)
