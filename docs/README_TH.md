# Go Template Structure

โครงสร้าง Enterprise-grade สำหรับ Golang API ที่รวม Best Practices ทั้งหมด

## 🏆 คุณสมบัติเด่น

✨ **Clean Architecture** - แยกชั้นงานอย่างชัดเจน  
🔐 **JWT Authentication** - ระบบยืนยันตัวตนที่ปลอดภัย  
📊 **PostgreSQL + Redis** - ฐานข้อมูลหลักและ Cache  
📝 **Swagger Documentation** - เอกสาร API อัตโนมัติ  
🧪 **Unit Testing** - ทดสอบครอบคลุม  
🐳 **Docker Support** - รองรับ Containerization  
⚡ **Gin Framework** - Web framework ที่รวดเร็ว  
📋 **Structured Logging** - บันทึกข้อมูลแบบมีโครงสร้าง  
🔄 **Graceful Shutdown** - ปิดระบบอย่างปลอดภัย  

## 🏗️ โครงสร้างโปรเจกต์

```
.
├── cmd/server/           # เซิร์ฟเวอร์หลัก
├── internal/
│   ├── config/          # การตั้งค่า
│   ├── domain/          # Business entities
│   ├── repository/      # Data layer
│   ├── service/         # Business logic
│   ├── handler/         # HTTP handlers
│   └── middleware/      # HTTP middlewares
├── pkg/
│   ├── utils/           # Utility functions
│   ├── logger/          # Logging package
│   └── database/        # Database connections
├── api/swagger/         # API documentation
├── test/                # Test files
├── scripts/             # Build scripts
└── docs/                # Documentation
```

## 🚀 เริ่มต้นใช้งาน

### 1. คัดลอกไฟล์ Environment
```bash
copy .env.example .env
```

### 2. รันด้วย Docker (แนะนำ)
```bash
docker-compose up -d
```

### 3. รันแบบ Local Development
```bash
# ติดตั้ง dependencies
go mod download

# รันเซิร์ฟเวอร์
make run
```

## 📚 API Documentation

เมื่อเซิร์ฟเวอร์รันแล้ว เข้าไปดู Swagger UI ได้ที่:
- http://localhost:8080/swagger/index.html

## 🔧 คำสั่งที่มีประโยชน์

```bash
make run          # รันเซิร์ฟเวอร์
make build        # สร้าง binary
make test         # รันเทส
make test-coverage # รันเทสพร้อม coverage
make lint         # ตรวจสอบ code quality
make swagger      # สร้าง swagger docs
make clean        # ลบไฟล์ที่ไม่จำเป็น
make docker-build # สร้าง Docker image
make docker-run   # รันด้วย Docker Compose
```

## 🎯 API Endpoints

### Authentication
- `POST /api/v1/auth/register` - สมัครสมาชิก
- `POST /api/v1/auth/login` - เข้าสู่ระบบ
- `POST /api/v1/auth/refresh` - ต่ออายุ token

### Users (ต้องยืนยันตัวตน)
- `GET /api/v1/users/profile` - ดูโปรไฟล์
- `PUT /api/v1/users/profile` - แก้ไขโปรไฟล์
- `GET /api/v1/users` - ดูรายการผู้ใช้
- `GET /api/v1/users/:id` - ดูผู้ใช้รายคน
- `PUT /api/v1/users/:id` - แก้ไขผู้ใช้
- `DELETE /api/v1/users/:id` - ลบผู้ใช้

## 🔨 สร้างจาก Template นี้

Template นี้ออกแบบมาให้ขยายได้ง่าย คุณสามารถ:

1. เพิ่ม Domain Models ใหม่ใน `internal/domain/`
2. สร้าง Repository ใหม่ใน `internal/repository/`
3. เพิ่ม Business Logic ใน `internal/service/`
4. สร้าง API Handlers ใน `internal/handler/`
5. เพิ่ม Middleware ใหม่ใน `internal/middleware/`

## 📊 ตัวอย่างการใช้งาน

### สมัครสมาชิก
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "username": "testuser",
    "password": "password123",
    "first_name": "Test",
    "last_name": "User"
  }'
```

### เข้าสู่ระบบ
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "password123"
  }'
```

## 🛠️ การพัฒนา

### ติดตั้ง Development Tools
```bash
make dev-deps
```

### รันเทสพร้อม Coverage
```bash
make test-coverage
```

### ตรวจสอบ Code Quality
```bash
make lint
```

## 📝 License

MIT License - ใช้งานได้อย่างอิสระ

---

🎉 **Happy Coding!** สร้างสรรค์ API ที่ยอดเยี่ยมด้วย Template นี้
