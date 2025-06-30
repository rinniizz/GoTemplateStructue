# 🔧 Mock Mode Guide

โหมด Mock เป็นโหมดพิเศษที่ให้คุณรันแอปพลิเคชันได้โดยไม่ต้องติดตั้ง PostgreSQL หรือ Redis

## ✨ คุณสมบัติ Mock Mode

- 🚀 **รันได้ทันที** - ไม่ต้องติดตั้งฐานข้อมูล
- 💾 **จำลองข้อมูล** - มี User ตัวอย่างพร้อมใช้
- 🔗 **API ครบ** - ทดสอบ API ได้ทุกตัว
- 🏥 **Health Check** - แสดงสถานะ Mock Mode

## 🎯 วิธีการเปิดใช้

### 📋 วิธีที่ 1: ใช้ Batch File (Windows)
```bash
# รัน Mock Mode
run-mock.bat

# รัน Production Mode  
run-production.bat
```

### 📋 วิธีที่ 2: ใช้ Environment Variable
```bash
# Windows
set MOCK_MODE=true
go run cmd/server/main.go

# Linux/Mac
MOCK_MODE=true go run cmd/server/main.go
```

### 📋 วิธีที่ 3: ใช้ Makefile
```bash
# Linux/Mac
make mock

# Windows
make mock-win
```

### 📋 วิธีที่ 4: แก้ไขไฟล์ .env
```env
MOCK_MODE=true
```

## 🧪 ทดสอบ API ใน Mock Mode

### ตรวจสอบสถานะ
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "ok",
  "mode": "mock",
  "message": "Server is running in mock mode",
  "time": "2025-06-30T10:30:00Z"
}
```

### เข้าสู่ระบบด้วย User ตัวอย่าง
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@example.com",
    "password": "password123"
  }'
```

### สมัครสมาชิกใหม่
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newuser@example.com",
    "username": "newuser",
    "password": "password123",
    "first_name": "New",
    "last_name": "User"
  }'
```

## 📊 ข้อมูลตัวอย่างใน Mock Mode

### 👤 User ตัวอย่าง
- **Email:** admin@example.com
- **Password:** password123
- **Username:** admin
- **Role:** Admin User

## 🔄 การเปลี่ยนโหมด

### เปลี่ยนจาก Mock เป็น Production
1. ติดตั้ง PostgreSQL และ Redis
2. แก้ไขไฟล์ `.env` ให้ `MOCK_MODE=false`
3. ตั้งค่า Database connection
4. รันใหม่

### เปลี่ยนจาก Production เป็น Mock
1. แก้ไขไฟล์ `.env` ให้ `MOCK_MODE=true`
2. รันใหม่

## ⚠️ ข้อจำกัดของ Mock Mode

- ข้อมูลจะหายเมื่อปิดเซิร์ฟเวอร์
- ไม่มี Persistence
- ไม่รองรับ Advanced Database Features
- Redis Cache จะใช้ Memory แทน

## 🎨 สำหรับ Developer

Mock Mode เหมาะสำหรับ:
- 🧪 **การทดสอบ API**
- 📝 **การเขียนเอกสาร**
- 🎯 **การ Demo**
- 🔧 **การพัฒนา Frontend**
- 📱 **การทดสอบ Mobile App**
