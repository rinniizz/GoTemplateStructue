package docs

// SimpleSwaggerJSON provides a basic swagger specification
const SimpleSwaggerJSON = `{
  "swagger": "2.0",
  "info": {
    "title": "Go Template API",
    "description": "เทมเพลต API สำหรับ Golang พร้อม Best Practices",
    "version": "1.0.0",
    "contact": {
      "name": "API Support",
      "email": "support@example.com"
    },
    "license": {
      "name": "MIT",
      "url": "https://opensource.org/licenses/MIT"
    }
  },
  "host": "localhost:8080",
  "basePath": "/api/v1",
  "schemes": ["http"],
  "consumes": ["application/json"],
  "produces": ["application/json"],
  "securityDefinitions": {
    "BearerAuth": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header",
      "description": "JWT Authorization header using the Bearer scheme. Example: 'Bearer {token}'"
    }
  },
  "paths": {
    "/health": {
      "get": {
        "summary": "Health Check",
        "description": "ตรวจสอบสถานะเซิร์ฟเวอร์",
        "tags": ["Health"],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "properties": {
                "status": {"type": "string"},
                "mode": {"type": "string"},
                "message": {"type": "string"},
                "time": {"type": "string"}
              }
            }
          }
        }
      }
    },
    "/auth/register": {
      "post": {
        "summary": "สมัครสมาชิก",
        "description": "สร้างบัญชีผู้ใช้ใหม่",
        "tags": ["Authentication"],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "email": {"type": "string", "format": "email"},
                "username": {"type": "string"},
                "password": {"type": "string", "minLength": 6},
                "first_name": {"type": "string"},
                "last_name": {"type": "string"}
              },
              "required": ["email", "username", "password", "first_name", "last_name"]
            }
          }
        ],
        "responses": {
          "201": {
            "description": "สมัครสมาชิกสำเร็จ",
            "schema": {
              "type": "object",
              "properties": {
                "success": {"type": "boolean"},
                "message": {"type": "string"},
                "data": {
                  "type": "object",
                  "properties": {
                    "user": {"$ref": "#/definitions/User"},
                    "access_token": {"type": "string"},
                    "refresh_token": {"type": "string"},
                    "expires_in": {"type": "integer"}
                  }
                }
              }
            }
          }
        }
      }
    },
    "/auth/login": {
      "post": {
        "summary": "เข้าสู่ระบบ",
        "description": "ยืนยันตัวตนและรับ token",
        "tags": ["Authentication"],
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "type": "object",
              "properties": {
                "email": {"type": "string", "format": "email"},
                "password": {"type": "string"}
              },
              "required": ["email", "password"]
            }
          }
        ],
        "responses": {
          "200": {
            "description": "เข้าสู่ระบบสำเร็จ"
          }
        }
      }
    },
    "/users/profile": {
      "get": {
        "summary": "ดูโปรไฟล์",
        "description": "ดูข้อมูลโปรไฟล์ของผู้ใช้ปัจจุบัน",
        "tags": ["Users"],
        "security": [{"BearerAuth": []}],
        "responses": {
          "200": {
            "description": "สำเร็จ"
          }
        }
      },
      "put": {
        "summary": "แก้ไขโปรไฟล์",
        "description": "แก้ไขข้อมูลโปรไฟล์ของผู้ใช้ปัจจุบัน",
        "tags": ["Users"],
        "security": [{"BearerAuth": []}],
        "responses": {
          "200": {
            "description": "สำเร็จ"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "properties": {
        "id": {"type": "integer"},
        "email": {"type": "string"},
        "username": {"type": "string"},
        "first_name": {"type": "string"},
        "last_name": {"type": "string"},
        "avatar": {"type": "string"},
        "is_active": {"type": "boolean"},
        "created_at": {"type": "string", "format": "date-time"},
        "updated_at": {"type": "string", "format": "date-time"}
      }
    }
  }
}`
