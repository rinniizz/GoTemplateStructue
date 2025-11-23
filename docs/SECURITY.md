# 🔒 Security Features Guide

เอกสารคำแนะนำการใช้งานฟีเจอร์ความปลอดภัยใน Go Template Structure

---

## 📋 ฟีเจอร์ความปลอดภัยที่มีอยู่

### 1️⃣ Rate Limiting (จำกัดจำนวน Request)

**ป้องกัน:** DDoS, Brute Force Attack

**การตั้งค่า:**
```go
// ใน cmd/server/main.go
router.Use(middleware.RateLimiter(10, 20))
// 10 = จำนวน requests ต่อวินาที
// 20 = burst size (อนุญาตให้พุ่งได้สูงสุด 20 requests)
```

**การปรับแต่ง:**
- **API ปกติ:** `RateLimiter(10, 20)` - 10 req/sec
- **API ที่ใช้บ่อย:** `RateLimiter(50, 100)` - 50 req/sec
- **API ที่อันตราย (Login, Register):** `RateLimiter(3, 5)` - 3 req/sec

**ตัวอย่างการใช้เฉพาะ endpoint:**
```go
auth := v1.Group("/auth")
auth.Use(middleware.RateLimiter(3, 5)) // จำกัด login/register
{
    auth.POST("/login", authHandler.Login)
    auth.POST("/register", authHandler.Register)
}
```

---

### 2️⃣ Request ID Tracking

**ประโยชน์:** ติดตาม request, debug ปัญหา, audit trail

**การใช้งานใน code:**
```go
func (h *Handler) SomeHandler(c *gin.Context) {
    requestID, _ := c.Get("RequestID")
    
    logger.Info("Processing request", map[string]interface{}{
        "request_id": requestID,
        "user_id": userID,
    })
}
```

**ดูใน Response Header:**
```bash
curl -I http://localhost:8080/api/v1/users
# จะเห็น: X-Request-ID: 550e8400-e29b-41d4-a716-446655440000
```

---

### 3️⃣ Security Headers

**ป้องกัน:** XSS, Clickjacking, MIME Sniffing

**Headers ที่ถูกเพิ่มอัตโนมัติ:**
- `X-Content-Type-Options: nosniff` - ป้องกัน MIME type sniffing
- `X-Frame-Options: DENY` - ป้องกัน Clickjacking
- `X-XSS-Protection: 1; mode=block` - เปิด XSS protection
- `Content-Security-Policy: default-src 'self'` - จำกัด resource loading
- `Referrer-Policy: strict-origin-when-cross-origin`
- `Permissions-Policy: geolocation=(), microphone=(), camera=()`

**การปรับแต่ง CSP:**
```go
// ใน internal/middleware/security_headers.go
c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'")
```

---

### 4️⃣ Audit Logging

**บันทึก:** ทุกการกระทำที่สำคัญ

**สิ่งที่ถูกบันทึก:**
- ✅ POST, PUT, DELETE requests (การแก้ไขข้อมูล)
- ✅ Login, Register, Refresh token
- ✅ Error responses (4xx, 5xx)
- ✅ User ID, Email, IP, User Agent
- ✅ Request ID สำหรับ tracking
- ✅ Latency (เวลาที่ใช้ประมวลผล)

**ตัวอย่าง Log:**
```json
{
  "level": "info",
  "msg": "Audit Log",
  "request_id": "550e8400-e29b-41d4-a716-446655440000",
  "method": "POST",
  "path": "/api/v1/auth/login",
  "status": 200,
  "latency_ms": 45,
  "ip": "192.168.1.100",
  "email": "user@example.com",
  "user_agent": "Mozilla/5.0..."
}
```

---

### 5️⃣ IP Filtering

**3 รูปแบบ:**

#### A. IP Whitelist (อนุญาตเฉพาะ IP ที่กำหนด)
```go
// สำหรับ Admin routes
adminRoutes := v1.Group("/admin")
adminRoutes.Use(middleware.IPWhitelist([]string{
    "192.168.1.100",
    "10.0.0.1",
}))
```

#### B. IP Blacklist (บล็อก IP ที่กำหนด)
```go
router.Use(middleware.IPBlacklist([]string{
    "123.45.67.89",  // IP ที่ทำร้าย
    "98.76.54.32",
}))
```

#### C. IP Range Filter (กำหนด CIDR range)
```go
// อนุญาตเฉพาะ internal network
internalAPI := v1.Group("/internal")
internalAPI.Use(middleware.IPRangeFilter([]string{
    "192.168.0.0/16",  // Private network
    "10.0.0.0/8",      // Private network
}))
```

---

### 6️⃣ Input Validation & Sanitization

**ป้องกัน:** SQL Injection, XSS

**การใช้งาน:**

```go
import "go-template-structure/pkg/utils"

// ใน handler
func (h *Handler) CreateProduct(c *gin.Context) {
    var req CreateProductRequest
    c.ShouldBindJSON(&req)
    
    // Sanitize input
    req.Name = utils.SanitizeString(req.Name)
    req.Description = utils.SanitizeString(req.Description)
    
    // Validate email
    if !utils.ValidateEmail(req.Email) {
        c.JSON(400, gin.H{"error": "Invalid email format"})
        return
    }
    
    // Validate password strength
    if !utils.ValidatePassword(req.Password) {
        c.JSON(400, gin.H{
            "error": "Password must be at least 8 characters with uppercase, lowercase, number, and special character"
        })
        return
    }
    
    // Check for SQL injection
    if utils.ContainsSQLInjection(req.Name) {
        c.JSON(400, gin.H{"error": "Invalid input detected"})
        return
    }
    
    // Check for XSS
    if utils.ContainsXSS(req.Description) {
        c.JSON(400, gin.H{"error": "Invalid input detected"})
        return
    }
}
```

**Functions ที่มี:**
- `SanitizeString(input)` - ลบ HTML, SQL patterns
- `ValidateEmail(email)` - ตรวจสอบ email format
- `ValidatePassword(password)` - ตรวจสอบความแข็งแรง password
- `GetPasswordStrength(password)` - คืนค่าระดับ 1-5
- `ValidatePhone(phone)` - ตรวจสอบเบอร์โทร (Thai format)
- `ValidateURL(url)` - ตรวจสอบ URL format
- `ContainsSQLInjection(input)` - ตรวจจับ SQL injection
- `ContainsXSS(input)` - ตรวจจับ XSS patterns

---

## 🎯 Best Practices

### 1. การใช้ Rate Limiting

```go
// ❌ ไม่ดี - rate limit ทุก endpoint เท่ากัน
router.Use(middleware.RateLimiter(10, 20))

// ✅ ดี - แยก rate limit ตาม endpoint
router.Use(middleware.RateLimiter(50, 100)) // Global: ปกติ

auth := v1.Group("/auth")
auth.Use(middleware.RateLimiter(3, 5))      // Login/Register: เข้มงวด
{
    auth.POST("/login", authHandler.Login)
}

uploads := v1.Group("/uploads")
uploads.Use(middleware.RateLimiter(2, 3))   // Upload: จำกัดมาก
{
    uploads.POST("/", uploadHandler.Upload)
}
```

### 2. การ Validate Input

```go
// ✅ Validate ก่อน process เสมอ
func (h *Handler) UpdateUser(c *gin.Context) {
    var req UpdateUserRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }
    
    // 1. Sanitize
    req.Name = utils.SanitizeString(req.Name)
    
    // 2. Validate format
    if req.Email != "" && !utils.ValidateEmail(req.Email) {
        c.JSON(400, gin.H{"error": "Invalid email"})
        return
    }
    
    // 3. Check attacks
    if utils.ContainsSQLInjection(req.Name) || utils.ContainsXSS(req.Name) {
        c.JSON(400, gin.H{"error": "Invalid input"})
        return
    }
    
    // 4. Process
    user, err := h.service.UpdateUser(req)
    // ...
}
```

### 3. การใช้ Audit Logging

```go
// Audit log จะบันทึกอัตโนมัติ แต่สามารถเพิ่มข้อมูลได้
func (h *Handler) DeleteUser(c *gin.Context) {
    userID := c.Param("id")
    
    // เพิ่มข้อมูลสำคัญใน context
    c.Set("action", "delete_user")
    c.Set("target_user_id", userID)
    
    err := h.service.DeleteUser(userID)
    
    if err != nil {
        logger.Error("Failed to delete user", map[string]interface{}{
            "request_id": c.GetString("RequestID"),
            "user_id": userID,
            "error": err.Error(),
        })
    }
}
```

### 4. IP Filtering สำหรับ Admin

```go
// ✅ ใช้ IP whitelist สำหรับ admin routes
admin := v1.Group("/admin")
admin.Use(middleware.IPWhitelist([]string{
    "192.168.1.100",  // Admin office IP
    "10.0.0.1",       // VPN IP
}))
admin.Use(middleware.JWTAuth(cfg.JWT.Secret)) // + JWT auth
{
    admin.GET("/users", adminHandler.GetAllUsers)
    admin.DELETE("/users/:id", adminHandler.DeleteUser)
}
```

---

## 🚨 การจัดการ Security Incidents

### 1. ตรวจจับ Rate Limit Exceeded

```bash
# ดู logs
tail -f logs/app.log | grep "Rate limit exceeded"

# บล็อก IP ถ้าพบการโจมตี
# แก้ไข main.go
router.Use(middleware.IPBlacklist([]string{
    "123.45.67.89",
}))
```

### 2. ตรวจสอบ SQL Injection Attempts

```bash
# ค้นหา injection attempts ใน logs
grep -i "Invalid input detected" logs/app.log

# จะเห็น IP ของผู้โจมตี
```

### 3. ตรวจสอบ Failed Login Attempts

```bash
# ดู audit logs
grep "login" logs/app.log | grep "status\":401"

# จะเห็น IP ที่พยายาม brute force
```

---

## 📊 Monitoring & Metrics

### สิ่งที่ควร Monitor

1. **Rate Limit Hits**
   - จำนวน requests ที่ถูกบล็อก
   - IP ที่ถูกบล็อกบ่อย

2. **Failed Authentication**
   - Login failures
   - Token refresh failures

3. **Suspicious Patterns**
   - SQL injection attempts
   - XSS attempts
   - Path traversal attempts

4. **Response Times**
   - Latency เฉลี่ย
   - Slow endpoints

### ตัวอย่าง Log Query (ถ้าใช้ ELK, Splunk)

```
# Count rate limit hits
level:warn AND message:"Rate limit exceeded" | stats count by ip

# Failed logins by IP
path:"/api/v1/auth/login" AND status:401 | stats count by ip

# SQL injection attempts
message:"Invalid input" AND message:"SQL" | stats count by ip
```

---

## 🔧 Configuration Tips

### Production Settings

```go
// High security for production
router.Use(middleware.RateLimiter(20, 40))           // เข้มงวดขึ้น
router.Use(middleware.IPBlacklist(blockedIPs))      // Block known attackers

// Enable HTTPS strict transport
c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")

// More restrictive CSP
c.Header("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self'")
```

### Development Settings

```go
// Relaxed for development
router.Use(middleware.RateLimiter(100, 200))  // เพิ่ม limit
// ปิด IP filtering
// ปิด HTTPS enforcement
```

---

## ✅ Security Checklist

ก่อน Deploy Production:

- [ ] ตั้งค่า Rate Limiting ที่เหมาะสม
- [ ] เปิด Security Headers
- [ ] ตรวจสอบ Audit Logging ทำงาน
- [ ] Validate input ทุก endpoint
- [ ] ใช้ HTTPS (TLS/SSL)
- [ ] ตั้งค่า IP Whitelist สำหรับ admin routes
- [ ] เปลี่ยน JWT Secret เป็น strong password
- [ ] ตั้งค่า CORS ให้ถูกต้อง
- [ ] Test rate limiting
- [ ] Setup monitoring & alerting

---

## 📚 เอกสารเพิ่มเติม

- [OWASP Top 10](https://owasp.org/www-project-top-ten/)
- [Gin Security Best Practices](https://github.com/gin-gonic/gin#dont-trust-all-proxies)
- [Go Security Checklist](https://github.com/guardrailsio/awesome-golang-security)

---

**หมายเหตุ:** ความปลอดภัยเป็นกระบวนการต่อเนื่อง ควร review และอัพเดทเป็นประจำ 🔒
