# Go Template Structure

เทมเพลตโครงสร้าง Golang แบบ E## ⚡ Quick Start

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
- �🔧 Mock mode (ไม่ต้องติดตั้ง Database)
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
- `make lint` - ตรวจสอบ code quality
- `make swagger` - สร้าง swagger docs
- `make clean` - ลบไฟล์ที่ไม่จำเป็น ที่รวม Best Practices ทั้งหมด

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
│   ├── mock/             # Mock implementations
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
│   ├── MOCK_MODE.md      # Mock mode documentation
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
- ✅ Mock Mode (ไม่ต้องมี Database)
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

## 🔧 Mock Mode

โปรเจกต์นี้รองรับ **Mock Mode** ที่ไม่ต้องติดตั้ง Database:
- `dev.bat` และ `make dev` จะเปิด Mock Mode อัตโนมัติ
- Mock Mode จะใช้ในหน่วยข้อมูลแทน Database จริง
- เหมาะสำหรับการพัฒนาและทดสอบ

## 🐳 Production Deployment

### Docker (แนะนำ)
```bash
docker-compose up -d
```

### Manual (ต้องติดตั้ง PostgreSQL และ Redis ก่อน)
```bash
# Copy environment file
cp .env.example .env

# Edit .env with your database settings
# Then run:
make run    # Linux/Mac
# หรือ
go run cmd/server/main.go  # Windows
```

## 🧪 ทดสอบ API
make run           # Linux/Mac
```

### 🧪 ทดสอบ API

เมื่อเซิร์ฟเวอร์รันแล้ว:

```bash
# ตรวจสอบสถานะ
curl http://localhost:8080/health

# เข้าสู่ระบบ (Mock Mode)
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

###  Development Features:
- **Hot Reload**: เปลี่ยนไฟล์แล้ว restart อัตโนมัติ (Air)
- **Mock Mode**: ไม่ต้องมี database
- **Debug Logging**: ข้อมูล log ละเอียด
- **Auto Swagger**: generate swagger docs อัตโนมัติ
- **CORS Enabled**: เรียก API จาก frontend ได้

## 🛠️ การใช้งาน

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

### 🚀 วิธีการรัน

#### วิธีที่ 1: รันแบบ Development (แนะนำ)
```bash
# Windows
dev.bat

# Linux/Mac
make dev
```

#### วิธีที่ 2: รันแบบ Production
```bash
# สร้าง binary ก่อน
build.bat  # Windows
make build # Linux/Mac

# รัน binary
bin/gotemplate.exe  # Windows
./bin/gotemplate    # Linux/Mac
```
```

**Linux/Mac:**
```bash
# ติดตั้ง swag CLI (ครั้งเดียว)
go install github.com/swaggo/swag/cmd/swag@latest

# Generate swagger docs
swag init -g cmd/server/main.go --output docs --parseDependency --parseInternal

# หรือใช้ Makefile
make swagger
```

### 📖 เข้าถึง Swagger UI

เมื่อเซิร์ฟเวอร์รันแล้ว สามารถเข้าถึง Swagger UI ได้ที่:
- **Swagger UI:** http://localhost:8080/swagger/index.html
- **JSON Schema:** http://localhost:8080/swagger/doc.json

### 🚀 Quick Start พร้อม Swagger

**วิธีที่ 1: ใช้ batch script (Windows)**
```bash
# รัน generate swagger และเริ่มเซิร์ฟเวอร์พร้อมกัน
run-with-swagger.bat
```

**วิธีที่ 2: Manual steps**
```bash
# 1. Generate swagger documentation
generate-swagger.bat

# 2. รันเซิร์ฟเวอร์ (Mock Mode)
run-mock.bat

# 3. เปิดบราวเซอร์ไปที่ http://localhost:8080/swagger/index.html
```

### ✨ API Endpoints (Auto-Generated)

Swagger UI จะแสดง API endpoints ทั้งหมดที่ถูก generate อัตโนมัติจาก code:

- **Authentication:**
  - `POST /api/v1/auth/register` - ลงทะเบียนผู้ใช้ใหม่
  - `POST /api/v1/auth/login` - เข้าสู่ระบบ
  - `POST /api/v1/auth/refresh` - รีเฟรช token

- **User Management:** (ต้อง authentication)
  - `GET /api/v1/users/profile` - ดูโปรไฟล์ตนเอง
  - `PUT /api/v1/users/profile` - แก้ไขโปรไฟล์ตนเอง
  - `GET /api/v1/users` - ดูรายชื่อผู้ใช้ทั้งหมด (พร้อม pagination)
  - `GET /api/v1/users/:id` - ดูข้อมูลผู้ใช้ตาม ID
  - `PUT /api/v1/users/:id` - แก้ไขข้อมูลผู้ใช้ตาม ID
  - `DELETE /api/v1/users/:id` - ลบผู้ใช้ตาม ID

- **System:**
  - `GET /health` - ตรวจสอบสถานะระบบ

**หมายเหตุ:** ข้อมูลทั้งหมดใน Swagger UI ถูก generate อัตโนมัติจาก swagger comments ใน source code

## 🧪 Testing

```bash
make test
```

## 📦 Build

```bash
make build
```

## 🔧 Available Commands

### Windows Batch Scripts
- `run-dev.bat` - 🔥 **Development Mode พร้อม Air Hot Reload (แนะนำ)**
- `run-with-swagger.bat` - � รันพร้อม Generate Swagger
- `run-dev-simple.bat` - 🛠️ Development Mode แบบง่าย (ไม่มี hot reload)
- `run-mock.bat` - 🔧 รันแบบ Mock Mode
- `run-production.bat` - 🏭 รันแบบ Production Mode
- `test-build.bat` - 🧪 ทดสอบการ Build
- `generate-swagger.bat` - 📚 Generate Swagger เท่านั้น
- `test-api.bat` - 🧪 ทดสอบ API endpoints ด้วย curl

### Makefile Commands (Linux/Mac)
- `make dev` - 🔥 **Development Mode พร้อม Air Hot Reload (แนะนำ)**
- `make dev-simple` - 🛠️ Development Mode แบบง่าย (ไม่มี hot reload)
- `make run` - รันเซิร์ฟเวอร์
- `make mock` - รันแบบ Mock Mode
- `make build` - สร้าง binary
- `make test` - รันเทส
- `make lint` - ตรวจสอบ code quality
- `make swagger` - สร้าง swagger docs
- `make clean` - ลบไฟล์ที่ไม่จำเป็น
