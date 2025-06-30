# Go Template Structure

เทมเพลตโครงสร้าง Golang แบบ Enterprise ที่รวม Best Practices ทั้งหมด

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
│   └── middleware/       # HTTP middlewares
├── pkg/                   # Public libraries
│   ├── utils/            # Utility functions
│   ├── logger/           # Logging package
│   └── database/         # Database connections
├── api/                   # API definitions
│   └── swagger/          # Swagger documentation
├── scripts/               # Build and deployment scripts
├── docs/                  # Documentation
├── test/                  # Test files
├── .env.example          # Environment variables example
├── Dockerfile            # Docker configuration
├── docker-compose.yml    # Docker compose
├── Makefile              # Build automation
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
- ✅ Graceful Shutdown

## � Quick Start

### 🔧 Mock Mode พร้อม Swagger (แนะนำ)

**Windows:**
```bash
# รัน generate swagger และเริ่มเซิร์ฟเวอร์
run-with-swagger.bat
```

### 🔧 Mock Mode แบบง่าย (ไม่ต้องมี Database)

**Windows:**
```bash
# 1. ทดสอบ build ก่อน
test-build.bat

# 2. รัน Mock Mode
run-mock.bat
```

**Linux/Mac:**
```bash
# 1. Generate Swagger
swag init -g cmd/server/main.go -o docs

# 2. รัน Mock Mode
MOCK_MODE=true go run cmd/server/main.go
```

### 🐳 Production Mode (มี Database)

**Docker (แนะนำ):**
```bash
docker-compose up -d
```

**Local:**
```bash
# ต้องติดตั้ง PostgreSQL และ Redis ก่อน
run-production.bat  # Windows
# หรือ
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

### 🚀 วิธีการรัน

#### วิธีที่ 1: รันแบบ Mock Mode (ไม่ต้องมี Database)
```bash
# ตั้งค่า MOCK_MODE=true ในไฟล์ .env หรือ
set MOCK_MODE=true
go run cmd/server/main.go
```

#### วิธีที่ 2: รันด้วย Docker (รวม Database)
```bash
docker-compose up -d
```

#### วิธีที่ 3: รันแบบ Local (ต้องมี PostgreSQL และ Redis)
```bash
make run
```

#### วิธีที่ 4: รันแบบ Development (Hot Reload)
```bash
make dev
```

## � Swagger API Documentation

โปรเจกต์นี้ใช้ **Swagger UI แบบ Auto-Generated** จาก code comments (ไม่ใช่สร้างเอง)

### 🔄 Generate Swagger Documentation

**Windows:**
```bash
# Generate swagger docs จาก code comments
generate-swagger.bat
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
- `run-with-swagger.bat` - 🚀 รันพร้อม Generate Swagger (แนะนำ)
- `run-mock.bat` - 🔧 รันแบบ Mock Mode
- `run-production.bat` - 🏭 รันแบบ Production Mode
- `test-build.bat` - 🧪 ทดสอบการ Build
- `generate-swagger.bat` - 📚 Generate Swagger เท่านั้น
- `test-api.bat` - 🧪 ทดสอบ API endpoints ด้วย curl

### Makefile Commands (Linux/Mac)
- `make run` - รันเซิร์ฟเวอร์
- `make mock` - รันแบบ Mock Mode
- `make build` - สร้าง binary
- `make test` - รันเทส
- `make lint` - ตรวจสอบ code quality
- `make swagger` - สร้าง swagger docs
- `make clean` - ลบไฟล์ที่ไม่จำเป็น
