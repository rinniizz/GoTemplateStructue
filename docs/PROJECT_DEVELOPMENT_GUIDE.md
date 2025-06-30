# 📘 คู่มือแนวทางการพัฒนาโปรเจกต์ Go Template Structure

## 🎯 ภาพรวมโครงสร้างโปรเจกต์

### 📁 โครงสร้างไฟล์และหน้าที่

```
c:\git\GoTemplateStructue\
├── cmd/                         # 🚀 Entry Points
│   └── server/
│       └── main.go             # จุดเริ่มต้นแอป + Swagger config
├── internal/                    # 🔒 Private Application Code
│   ├── config/                 # ⚙️ Configuration Management
│   │   └── config.go           # โหลด ENV, validate config
│   ├── domain/                 # 🏛️ Business Entities
│   │   ├── user.go             # User entity + validation
│   │   └── auth.go             # Auth-related entities
│   ├── interfaces/             # 🔌 Contracts & Abstractions
│   │   ├── repository.go       # Repository interfaces
│   │   └── service.go          # Service interfaces
│   ├── repository/             # 💾 Data Access Layer
│   │   ├── user_repository.go  # User CRUD operations
│   │   └── postgres/           # PostgreSQL implementations
│   ├── service/                # 🧠 Business Logic Layer
│   │   ├── user_service.go     # User business logic
│   │   └── auth_service.go     # Authentication logic
│   ├── handler/                # 🌐 HTTP Layer
│   │   ├── handler.go          # Handler struct + routes
│   │   ├── user_handler.go     # User HTTP endpoints
│   │   └── auth_handler.go     # Auth HTTP endpoints
│   ├── middleware/             # 🛡️ HTTP Middlewares
│   │   ├── auth.go             # JWT validation
│   │   ├── cors.go             # CORS configuration
│   │   └── logger.go           # Request logging
│   └── mock/                   # 🎭 Mock Implementations
│       ├── user_service.go     # Mock user service
│       └── auth_service.go     # Mock auth service
├── pkg/                        # 📦 Public Libraries
│   ├── utils/                  # 🔧 Utility Functions
│   │   ├── response.go         # Standard API responses
│   │   ├── validation.go       # Input validation helpers
│   │   └── hash.go             # Password hashing
│   ├── logger/                 # 📊 Logging Package
│   │   └── logger.go           # Structured logging setup
│   └── database/               # 🗄️ Database Connections
│       ├── postgres.go         # PostgreSQL connection
│       └── redis.go            # Redis connection
├── docs/                       # 📚 Documentation
│   ├── docs.go                 # Auto-generated Swagger
│   ├── swagger.json            # Swagger JSON spec
│   ├── swagger.yaml            # Swagger YAML spec
│   └── *.md                    # Documentation files
├── test/                       # 🧪 Test Files
│   ├── integration/            # Integration tests
│   └── unit/                   # Unit tests
├── bin/                        # 📦 Build Output
├── tmp/                        # 🔄 Air temp files
├── .env.example               # 📋 Environment template
├── Dockerfile                 # 🐳 Docker configuration
├── docker-compose.yml         # 🐳 Docker compose
├── Makefile                   # 🔨 Build automation
├── dev.bat                    # 🔥 Windows dev script
├── build.bat                  # 🔨 Windows build script
├── test.bat                   # 🧪 Windows test script
└── README.md                  # 📖 Project overview
```

## 🛠️ การพัฒนาแต่ละส่วน

### 1. 🏛️ เพิ่ม Business Entity ใหม่

**ตัวอย่าง: เพิ่ม Product entity**

#### 1.1 สร้าง Domain Entity
```go
// internal/domain/product.go
package domain

import (
    "time"
    "github.com/go-playground/validator/v10"
)

type Product struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" validate:"required,min=2,max=100"`
    Description string    `json:"description" validate:"max=500"`
    Price       float64   `json:"price" validate:"required,min=0"`
    CategoryID  uint      `json:"category_id" validate:"required"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
    Name        string  `json:"name" validate:"required,min=2,max=100"`
    Description string  `json:"description" validate:"max=500"`
    Price       float64 `json:"price" validate:"required,min=0"`
    CategoryID  uint    `json:"category_id" validate:"required"`
}

type UpdateProductRequest struct {
    Name        *string  `json:"name,omitempty" validate:"omitempty,min=2,max=100"`
    Description *string  `json:"description,omitempty" validate:"omitempty,max=500"`
    Price       *float64 `json:"price,omitempty" validate:"omitempty,min=0"`
    CategoryID  *uint    `json:"category_id,omitempty"`
}

func (p *Product) Validate() error {
    validate := validator.New()
    return validate.Struct(p)
}
```

#### 1.2 เพิ่ม Interface
```go
// internal/interfaces/repository.go (เพิ่มเข้าไป)
type ProductRepository interface {
    Create(product *domain.Product) error
    GetByID(id uint) (*domain.Product, error)
    GetAll(limit, offset int) ([]*domain.Product, error)
    Update(id uint, product *domain.Product) error
    Delete(id uint) error
    GetByCategory(categoryID uint, limit, offset int) ([]*domain.Product, error)
}

// internal/interfaces/service.go (เพิ่มเข้าไป)
type ProductService interface {
    CreateProduct(req *domain.CreateProductRequest) (*domain.Product, error)
    GetProduct(id uint) (*domain.Product, error)
    GetProducts(limit, offset int) ([]*domain.Product, error)
    UpdateProduct(id uint, req *domain.UpdateProductRequest) (*domain.Product, error)
    DeleteProduct(id uint) error
    GetProductsByCategory(categoryID uint, limit, offset int) ([]*domain.Product, error)
}
```

#### 1.3 สร้าง Repository Implementation
```go
// internal/repository/product_repository.go
package repository

import (
    "gorm.io/gorm"
    "your-project/internal/domain"
    "your-project/internal/interfaces"
)

type productRepository struct {
    db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepository {
    return &productRepository{db: db}
}

func (r *productRepository) Create(product *domain.Product) error {
    return r.db.Create(product).Error
}

func (r *productRepository) GetByID(id uint) (*domain.Product, error) {
    var product domain.Product
    err := r.db.First(&product, id).Error
    if err != nil {
        return nil, err
    }
    return &product, nil
}

func (r *productRepository) GetAll(limit, offset int) ([]*domain.Product, error) {
    var products []*domain.Product
    err := r.db.Limit(limit).Offset(offset).Find(&products).Error
    return products, err
}

func (r *productRepository) Update(id uint, product *domain.Product) error {
    return r.db.Model(&domain.Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *productRepository) Delete(id uint) error {
    return r.db.Delete(&domain.Product{}, id).Error
}

func (r *productRepository) GetByCategory(categoryID uint, limit, offset int) ([]*domain.Product, error) {
    var products []*domain.Product
    err := r.db.Where("category_id = ?", categoryID).Limit(limit).Offset(offset).Find(&products).Error
    return products, err
}
```

#### 1.4 สร้าง Service Implementation
```go
// internal/service/product_service.go
package service

import (
    "errors"
    "your-project/internal/domain"
    "your-project/internal/interfaces"
)

type productService struct {
    productRepo interfaces.ProductRepository
}

func NewProductService(productRepo interfaces.ProductRepository) interfaces.ProductService {
    return &productService{
        productRepo: productRepo,
    }
}

func (s *productService) CreateProduct(req *domain.CreateProductRequest) (*domain.Product, error) {
    // Validate request
    if err := validator.New().Struct(req); err != nil {
        return nil, err
    }

    product := &domain.Product{
        Name:        req.Name,
        Description: req.Description,
        Price:       req.Price,
        CategoryID:  req.CategoryID,
    }

    if err := s.productRepo.Create(product); err != nil {
        return nil, err
    }

    return product, nil
}

func (s *productService) GetProduct(id uint) (*domain.Product, error) {
    if id == 0 {
        return nil, errors.New("invalid product ID")
    }

    return s.productRepo.GetByID(id)
}

func (s *productService) GetProducts(limit, offset int) ([]*domain.Product, error) {
    if limit <= 0 {
        limit = 10
    }
    if limit > 100 {
        limit = 100
    }
    if offset < 0 {
        offset = 0
    }

    return s.productRepo.GetAll(limit, offset)
}

func (s *productService) UpdateProduct(id uint, req *domain.UpdateProductRequest) (*domain.Product, error) {
    // Validate request
    if err := validator.New().Struct(req); err != nil {
        return nil, err
    }

    // Check if product exists
    existingProduct, err := s.productRepo.GetByID(id)
    if err != nil {
        return nil, err
    }

    // Update fields if provided
    if req.Name != nil {
        existingProduct.Name = *req.Name
    }
    if req.Description != nil {
        existingProduct.Description = *req.Description
    }
    if req.Price != nil {
        existingProduct.Price = *req.Price
    }
    if req.CategoryID != nil {
        existingProduct.CategoryID = *req.CategoryID
    }

    if err := s.productRepo.Update(id, existingProduct); err != nil {
        return nil, err
    }

    return existingProduct, nil
}

func (s *productService) DeleteProduct(id uint) error {
    if id == 0 {
        return errors.New("invalid product ID")
    }

    // Check if product exists
    _, err := s.productRepo.GetByID(id)
    if err != nil {
        return err
    }

    return s.productRepo.Delete(id)
}

func (s *productService) GetProductsByCategory(categoryID uint, limit, offset int) ([]*domain.Product, error) {
    if categoryID == 0 {
        return nil, errors.New("invalid category ID")
    }

    if limit <= 0 {
        limit = 10
    }
    if limit > 100 {
        limit = 100
    }
    if offset < 0 {
        offset = 0
    }

    return s.productRepo.GetByCategory(categoryID, limit, offset)
}
```

#### 1.5 สร้าง HTTP Handler
```go
// internal/handler/product_handler.go
package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "your-project/internal/domain"
    "your-project/internal/interfaces"
    "your-project/pkg/utils"
)

type ProductHandler struct {
    productService interfaces.ProductService
}

func NewProductHandler(productService interfaces.ProductService) *ProductHandler {
    return &ProductHandler{
        productService: productService,
    }
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product with the provided information
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      domain.CreateProductRequest  true  "Product information"
// @Success      201      {object}  utils.SuccessResponse{data=domain.Product}
// @Failure      400      {object}  utils.ErrorResponse
// @Failure      500      {object}  utils.ErrorResponse
// @Router       /api/v1/products [post]
// @Security     BearerAuth
func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var req domain.CreateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
        return
    }

    product, err := h.productService.CreateProduct(&req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create product", err.Error())
        return
    }

    utils.SuccessResponse(c, http.StatusCreated, "Product created successfully", product)
}

// GetProduct godoc
// @Summary      Get product by ID
// @Description  Get a single product by its ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  utils.SuccessResponse{data=domain.Product}
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Router       /api/v1/products/{id} [get]
func (h *ProductHandler) GetProduct(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
        return
    }

    product, err := h.productService.GetProduct(uint(id))
    if err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "Product not found", err.Error())
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "Product retrieved successfully", product)
}

// GetProducts godoc
// @Summary      Get all products
// @Description  Get a list of all products with pagination
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        limit   query     int  false  "Number of products to return (default: 10, max: 100)"
// @Param        offset  query     int  false  "Number of products to skip (default: 0)"
// @Success      200     {object}  utils.SuccessResponse{data=[]domain.Product}
// @Failure      500     {object}  utils.ErrorResponse
// @Router       /api/v1/products [get]
func (h *ProductHandler) GetProducts(c *gin.Context) {
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

    products, err := h.productService.GetProducts(limit, offset)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to retrieve products", err.Error())
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "Products retrieved successfully", products)
}

// UpdateProduct godoc
// @Summary      Update product
// @Description  Update an existing product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id       path      int                           true  "Product ID"
// @Param        product  body      domain.UpdateProductRequest   true  "Updated product information"
// @Success      200      {object}  utils.SuccessResponse{data=domain.Product}
// @Failure      400      {object}  utils.ErrorResponse
// @Failure      404      {object}  utils.ErrorResponse
// @Failure      500      {object}  utils.ErrorResponse
// @Router       /api/v1/products/{id} [put]
// @Security     BearerAuth
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
        return
    }

    var req domain.UpdateProductRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request payload", err.Error())
        return
    }

    product, err := h.productService.UpdateProduct(uint(id), &req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to update product", err.Error())
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "Product updated successfully", product)
}

// DeleteProduct godoc
// @Summary      Delete product
// @Description  Delete a product by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/products/{id} [delete]
// @Security     BearerAuth
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.ParseUint(idParam, 10, 32)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
        return
    }

    err = h.productService.DeleteProduct(uint(id))
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to delete product", err.Error())
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "Product deleted successfully", nil)
}
```

#### 1.6 เพิ่ม Routes ใน Handler
```go
// internal/handler/handler.go (อัปเดต)
func (h *Handler) SetupRoutes() *gin.Engine {
    // ...existing code...

    // Product routes
    productHandler := NewProductHandler(h.ProductService)
    products := v1.Group("/products")
    {
        products.GET("", productHandler.GetProducts)
        products.GET("/:id", productHandler.GetProduct)
        
        // Protected routes
        products.Use(h.AuthMiddleware.RequireAuth())
        products.POST("", productHandler.CreateProduct)
        products.PUT("/:id", productHandler.UpdateProduct)
        products.DELETE("/:id", productHandler.DeleteProduct)
    }

    return r
}
```

#### 1.7 สร้าง Mock Implementation
```go
// internal/mock/product_service.go
package mock

import (
    "errors"
    "your-project/internal/domain"
    "your-project/internal/interfaces"
)

type mockProductService struct {
    products map[uint]*domain.Product
    nextID   uint
}

func NewMockProductService() interfaces.ProductService {
    return &mockProductService{
        products: make(map[uint]*domain.Product),
        nextID:   1,
    }
}

func (m *mockProductService) CreateProduct(req *domain.CreateProductRequest) (*domain.Product, error) {
    product := &domain.Product{
        ID:          m.nextID,
        Name:        req.Name,
        Description: req.Description,
        Price:       req.Price,
        CategoryID:  req.CategoryID,
    }
    
    m.products[m.nextID] = product
    m.nextID++
    
    return product, nil
}

func (m *mockProductService) GetProduct(id uint) (*domain.Product, error) {
    product, exists := m.products[id]
    if !exists {
        return nil, errors.New("product not found")
    }
    return product, nil
}

func (m *mockProductService) GetProducts(limit, offset int) ([]*domain.Product, error) {
    var products []*domain.Product
    count := 0
    skipped := 0
    
    for _, product := range m.products {
        if skipped < offset {
            skipped++
            continue
        }
        if count >= limit {
            break
        }
        products = append(products, product)
        count++
    }
    
    return products, nil
}

func (m *mockProductService) UpdateProduct(id uint, req *domain.UpdateProductRequest) (*domain.Product, error) {
    product, exists := m.products[id]
    if !exists {
        return nil, errors.New("product not found")
    }
    
    if req.Name != nil {
        product.Name = *req.Name
    }
    if req.Description != nil {
        product.Description = *req.Description
    }
    if req.Price != nil {
        product.Price = *req.Price
    }
    if req.CategoryID != nil {
        product.CategoryID = *req.CategoryID
    }
    
    return product, nil
}

func (m *mockProductService) DeleteProduct(id uint) error {
    _, exists := m.products[id]
    if !exists {
        return errors.New("product not found")
    }
    
    delete(m.products, id)
    return nil
}

func (m *mockProductService) GetProductsByCategory(categoryID uint, limit, offset int) ([]*domain.Product, error) {
    var products []*domain.Product
    count := 0
    skipped := 0
    
    for _, product := range m.products {
        if product.CategoryID != categoryID {
            continue
        }
        if skipped < offset {
            skipped++
            continue
        }
        if count >= limit {
            break
        }
        products = append(products, product)
        count++
    }
    
    return products, nil
}
```

#### 1.8 เพิ่มใน Dependency Injection
```go
// cmd/server/main.go (อัปเดต)
func main() {
    // ...existing code...

    if mockMode {
        // Mock services
        productService := mock.NewMockProductService()
        
        // Initialize handler with mock services
        handler := handler.NewHandler(
            userService,
            authService,
            productService, // เพิ่มใหม่
            authMiddleware,
        )
    } else {
        // Real services
        productRepo := repository.NewProductRepository(db)
        productService := service.NewProductService(productRepo)
        
        // Initialize handler with real services
        handler := handler.NewHandler(
            userService,
            authService,
            productService, // เพิ่มใหม่
            authMiddleware,
        )
    }

    // ...existing code...
}
```

### 2. 🧪 การเขียน Test

#### 2.1 Unit Test สำหรับ Service
```go
// test/unit/product_service_test.go
package unit

import (
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "your-project/internal/domain"
    "your-project/internal/service"
)

// Mock repository
type mockProductRepository struct {
    mock.Mock
}

func (m *mockProductRepository) Create(product *domain.Product) error {
    args := m.Called(product)
    return args.Error(0)
}

func (m *mockProductRepository) GetByID(id uint) (*domain.Product, error) {
    args := m.Called(id)
    return args.Get(0).(*domain.Product), args.Error(1)
}

// เพิ่ม methods อื่นๆ...

func TestProductService_CreateProduct(t *testing.T) {
    // Arrange
    mockRepo := new(mockProductRepository)
    productService := service.NewProductService(mockRepo)
    
    req := &domain.CreateProductRequest{
        Name:        "Test Product",
        Description: "Test Description",
        Price:       99.99,
        CategoryID:  1,
    }
    
    expectedProduct := &domain.Product{
        ID:          1,
        Name:        req.Name,
        Description: req.Description,
        Price:       req.Price,
        CategoryID:  req.CategoryID,
    }
    
    mockRepo.On("Create", mock.AnythingOfType("*domain.Product")).Return(nil)
    
    // Act
    result, err := productService.CreateProduct(req)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, expectedProduct.Name, result.Name)
    assert.Equal(t, expectedProduct.Price, result.Price)
    mockRepo.AssertExpectations(t)
}

func TestProductService_CreateProduct_ValidationError(t *testing.T) {
    // Arrange
    mockRepo := new(mockProductRepository)
    productService := service.NewProductService(mockRepo)
    
    req := &domain.CreateProductRequest{
        Name:        "", // Invalid: empty name
        Description: "Test Description",
        Price:       99.99,
        CategoryID:  1,
    }
    
    // Act
    result, err := productService.CreateProduct(req)
    
    // Assert
    assert.Error(t, err)
    assert.Nil(t, result)
    mockRepo.AssertNotCalled(t, "Create")
}
```

#### 2.2 Integration Test สำหรับ API
```go
// test/integration/product_api_test.go
package integration

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "your-project/internal/domain"
    "your-project/internal/handler"
    "your-project/internal/mock"
)

func setupTestRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    
    // Use mock services
    productService := mock.NewMockProductService()
    productHandler := handler.NewProductHandler(productService)
    
    r := gin.New()
    v1 := r.Group("/api/v1")
    products := v1.Group("/products")
    {
        products.GET("", productHandler.GetProducts)
        products.POST("", productHandler.CreateProduct)
        products.GET("/:id", productHandler.GetProduct)
        products.PUT("/:id", productHandler.UpdateProduct)
        products.DELETE("/:id", productHandler.DeleteProduct)
    }
    
    return r
}

func TestCreateProduct_API(t *testing.T) {
    // Arrange
    router := setupTestRouter()
    
    payload := domain.CreateProductRequest{
        Name:        "Test Product",
        Description: "Test Description",
        Price:       99.99,
        CategoryID:  1,
    }
    
    jsonPayload, _ := json.Marshal(payload)
    
    // Act
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(jsonPayload))
    req.Header.Set("Content-Type", "application/json")
    router.ServeHTTP(w, req)
    
    // Assert
    assert.Equal(t, http.StatusCreated, w.Code)
    
    var response map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.Equal(t, "Product created successfully", response["message"])
    assert.NotNil(t, response["data"])
}

func TestGetProducts_API(t *testing.T) {
    // Arrange
    router := setupTestRouter()
    
    // Act
    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/api/v1/products", nil)
    router.ServeHTTP(w, req)
    
    // Assert
    assert.Equal(t, http.StatusOK, w.Code)
    
    var response map[string]interface{}
    err := json.Unmarshal(w.Body.Bytes(), &response)
    assert.NoError(t, err)
    assert.Equal(t, "Products retrieved successfully", response["message"])
    assert.NotNil(t, response["data"])
}
```

### 3. 🔄 Workflow การพัฒนา

#### 3.1 ขั้นตอนการเพิ่ม Feature ใหม่

1. **วางแผน Feature**
   ```bash
   # สร้าง branch ใหม่
   git checkout -b feature/product-management
   ```

2. **สร้าง Domain Entity**
   - เขียน `internal/domain/product.go`
   - เพิ่ม validation rules
   - เขียน unit tests สำหรับ domain

3. **สร้าง Interface**
   - เพิ่ม interfaces ใน `internal/interfaces/`
   - กำหนด contracts ที่ชัดเจน

4. **สร้าง Repository**
   - เขียน `internal/repository/product_repository.go`
   - เขียน unit tests สำหรับ repository

5. **สร้าง Service**
   - เขียน `internal/service/product_service.go`
   - เขียน business logic
   - เขียน unit tests สำหรับ service

6. **สร้าง Handler**
   - เขียน `internal/handler/product_handler.go`
   - เพิ่ม Swagger comments
   - เขียน integration tests

7. **สร้าง Mock**
   - เขียน `internal/mock/product_service.go`
   - เพื่อใช้ใน development mode

8. **ทดสอบ**
   ```bash
   # รัน tests
   test.bat  # Windows
   make test # Linux/Mac
   
   # ทดสอบใน dev mode
   dev.bat   # Windows
   make dev  # Linux/Mac
   ```

9. **Generate Swagger**
   ```bash
   # Swagger จะ generate อัตโนมัติเมื่อรัน dev หรือ build
   # หรือ generate manual:
   make swagger
   ```

10. **Commit และ Push**
    ```bash
    git add .
    git commit -m "feat: add product management feature"
    git push origin feature/product-management
    ```

#### 3.2 การตรวจสอบคุณภาพ Code

##### A. ก่อน Commit
```bash
# ตรวจสอบ syntax และ format
go fmt ./...
go vet ./...

# รัน tests
go test ./...

# ตรวจสอบ coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out

# ใช้ linter (ถ้าติดตั้งแล้ว)
golangci-lint run
```

##### B. การทดสอบ Integration
```bash
# ทดสอบ build
test.bat # Windows
make test # Linux/Mac

# ทดสอบ development mode
dev.bat # Windows
make dev # Linux/Mac

# ทดสอบ production build
build.bat # Windows
make build # Linux/Mac

# ทดสอบ Docker
docker-compose up -d
```

##### C. การทดสอบ API
```bash
# ใช้ curl
curl -X GET http://localhost:8080/health
curl -X GET http://localhost:8080/api/v1/products

# ใช้ Swagger UI
# เปิด http://localhost:8080/swagger/index.html

# ทดสอบ authentication
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@example.com","password":"password123"}'
```

### 4. 📊 การ Debug และ Monitoring

#### 4.1 Logging
```go
// ใช้ logger ที่มีใน pkg/logger
import "your-project/pkg/logger"

func (s *productService) CreateProduct(req *domain.CreateProductRequest) (*domain.Product, error) {
    logger.Info("Creating new product", 
        logger.String("name", req.Name),
        logger.Float64("price", req.Price),
    )
    
    product, err := s.productRepo.Create(&domain.Product{
        Name:        req.Name,
        Description: req.Description,
        Price:       req.Price,
        CategoryID:  req.CategoryID,
    })
    
    if err != nil {
        logger.Error("Failed to create product", 
            logger.String("name", req.Name),
            logger.Error(err),
        )
        return nil, err
    }
    
    logger.Info("Product created successfully", 
        logger.Uint("id", product.ID),
        logger.String("name", product.Name),
    )
    
    return product, nil
}
```

#### 4.2 Error Handling
```go
// กำหนด custom errors
var (
    ErrProductNotFound = errors.New("product not found")
    ErrInvalidProductID = errors.New("invalid product ID")
    ErrProductNameRequired = errors.New("product name is required")
)

// ใช้ใน service
func (s *productService) GetProduct(id uint) (*domain.Product, error) {
    if id == 0 {
        return nil, ErrInvalidProductID
    }
    
    product, err := s.productRepo.GetByID(id)
    if err != nil {
        if errors.Is(err, gorm.ErrRecordNotFound) {
            return nil, ErrProductNotFound
        }
        return nil, err
    }
    
    return product, nil
}

// ใช้ใน handler
func (h *ProductHandler) GetProduct(c *gin.Context) {
    id, err := strconv.ParseUint(c.Param("id"), 10, 32)
    if err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
        return
    }
    
    product, err := h.productService.GetProduct(uint(id))
    if err != nil {
        switch {
        case errors.Is(err, service.ErrProductNotFound):
            utils.ErrorResponse(c, http.StatusNotFound, "Product not found", err.Error())
        case errors.Is(err, service.ErrInvalidProductID):
            utils.ErrorResponse(c, http.StatusBadRequest, "Invalid product ID", err.Error())
        default:
            utils.ErrorResponse(c, http.StatusInternalServerError, "Internal server error", err.Error())
        }
        return
    }
    
    utils.SuccessResponse(c, http.StatusOK, "Product retrieved successfully", product)
}
```

#### 4.3 Performance Monitoring
```go
// เพิ่ม metrics ใน handler
import "time"

func (h *ProductHandler) GetProducts(c *gin.Context) {
    start := time.Now()
    defer func() {
        duration := time.Since(start)
        logger.Info("GetProducts completed",
            logger.Duration("duration", duration),
            logger.String("path", c.Request.URL.Path),
        )
    }()
    
    // ... handler logic ...
}
```

### 5. 🚀 Deployment และ Production

#### 5.1 การเตรียม Production Build
```bash
# Windows
build.bat

# Linux/Mac
make build

# Output: bin/gotemplate.exe (Windows) หรือ bin/gotemplate (Linux/Mac)
```

#### 5.2 Docker Deployment
```dockerfile
# Dockerfile มีการ multi-stage build
# Stage 1: Build
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/server/main.go

# Stage 2: Runtime
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs
CMD ["./main"]
```

```bash
# Build และ run ด้วย Docker
docker-compose up -d

# หรือ build manual
docker build -t go-template .
docker run -p 8080:8080 go-template
```

#### 5.3 Environment Variables สำหรับ Production
```bash
# .env
SERVER_PORT=8080
GIN_MODE=release
LOG_LEVEL=info

# Database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=myapp

# Redis
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT
JWT_SECRET=your-super-secret-key

# ปิด Mock Mode
MOCK_MODE=false
```

### 6. 🔧 Best Practices และ Guidelines

#### 6.1 Code Organization
- **Domain-Driven Design**: จัดกลุ่ม code ตาม business domain
- **Separation of Concerns**: แยก layer ให้ชัดเจน
- **Dependency Injection**: ใช้ interfaces สำหรับ loose coupling
- **Single Responsibility**: แต่ละ function/method ทำงานเดียว

#### 6.2 Naming Conventions
```go
// Package names: lowercase, single word
package user

// Struct names: PascalCase
type UserService struct {}

// Interface names: PascalCase, often with -er suffix
type UserRepository interface {}

// Function names: PascalCase (public), camelCase (private)
func CreateUser() {} // public
func validateUser() {} // private

// Constants: PascalCase or UPPER_CASE
const MaxRetries = 3
const DEFAULT_PAGE_SIZE = 10

// Variables: camelCase
var userService UserService
```

#### 6.3 Error Handling
```go
// สร้าง custom error types
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

// ใช้ wrapped errors
if err != nil {
    return fmt.Errorf("failed to create user: %w", err)
}

// ตรวจสอบ error types
if validationErr, ok := err.(ValidationError); ok {
    // handle validation error
}
```

#### 6.4 Testing Guidelines
```go
// Test naming: Test{FunctionName}_{Scenario}
func TestCreateUser_Success(t *testing.T) {}
func TestCreateUser_InvalidEmail(t *testing.T) {}
func TestCreateUser_DuplicateEmail(t *testing.T) {}

// ใช้ table-driven tests สำหรับ multiple scenarios
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        wantErr bool
    }{
        {"valid email", "test@example.com", false},
        {"invalid email", "invalid-email", true},
        {"empty email", "", true},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := ValidateEmail(tt.email)
            if (err != nil) != tt.wantErr {
                t.Errorf("ValidateEmail() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

#### 6.5 Security Considerations
- **Input Validation**: ตรวจสอบ input ทุกตัว
- **SQL Injection**: ใช้ prepared statements หรือ ORM
- **Authentication**: ใช้ JWT tokens
- **Authorization**: ตรวจสอบ permissions ทุก endpoint
- **Rate Limiting**: จำกัด requests per minute
- **CORS**: กำหนด allowed origins
- **HTTPS**: ใช้ TLS ใน production

### 7. 📈 Performance Optimization

#### 7.1 Database Optimization
```go
// ใช้ database indexes
type User struct {
    ID    uint   `gorm:"primaryKey"`
    Email string `gorm:"uniqueIndex"` // index สำหรับ unique email
    Name  string `gorm:"index"`       // index สำหรับ search
}

// ใช้ pagination
func (r *userRepository) GetUsers(limit, offset int) ([]*domain.User, error) {
    var users []*domain.User
    err := r.db.Limit(limit).Offset(offset).Find(&users).Error
    return users, err
}

// ใช้ select specific fields
func (r *userRepository) GetUserSummary(id uint) (*UserSummary, error) {
    var summary UserSummary
    err := r.db.Select("id, name, email").First(&summary, id).Error
    return &summary, err
}
```

#### 7.2 Caching
```go
// ใช้ Redis สำหรับ cache
import "github.com/go-redis/redis/v8"

type userService struct {
    userRepo interfaces.UserRepository
    redis    *redis.Client
}

func (s *userService) GetUser(id uint) (*domain.User, error) {
    // Try cache first
    cacheKey := fmt.Sprintf("user:%d", id)
    cached, err := s.redis.Get(ctx, cacheKey).Result()
    if err == nil {
        var user domain.User
        json.Unmarshal([]byte(cached), &user)
        return &user, nil
    }
    
    // Get from database
    user, err := s.userRepo.GetByID(id)
    if err != nil {
        return nil, err
    }
    
    // Cache result
    userJSON, _ := json.Marshal(user)
    s.redis.Set(ctx, cacheKey, userJSON, 5*time.Minute)
    
    return user, nil
}
```

### 8. 📝 Documentation

#### 8.1 Swagger Documentation
```go
// เขียน Swagger comments ที่ครบถ้วน
// @Summary      Short description
// @Description  Detailed description
// @Tags         group-name
// @Accept       json
// @Produce      json
// @Param        name  type      data-type  required  "description"
// @Success      200   {object}  ResponseType
// @Failure      400   {object}  ErrorResponse
// @Router       /path [method]
// @Security     BearerAuth
```

#### 8.2 Code Comments
```go
// Package user provides user management functionality.
// It includes user creation, authentication, and profile management.
package user

// UserService provides business logic for user operations.
// It handles user validation, password hashing, and business rules.
type UserService struct {
    repo interfaces.UserRepository
}

// CreateUser creates a new user with the provided information.
// It validates the input, checks for duplicate emails, hashes the password,
// and stores the user in the database.
//
// Parameters:
//   - req: CreateUserRequest containing user information
//
// Returns:
//   - *User: Created user object
//   - error: Any error that occurred during creation
func (s *UserService) CreateUser(req *CreateUserRequest) (*User, error) {
    // Implementation...
}
```

### 9. 🔍 Monitoring และ Troubleshooting

#### 9.1 Health Checks
```go
// เพิ่ม health check endpoints
func (h *Handler) HealthCheck(c *gin.Context) {
    health := map[string]interface{}{
        "status":    "ok",
        "timestamp": time.Now().Unix(),
        "version":   "1.0.0",
        "services": map[string]string{
            "database": "connected",
            "redis":    "connected",
        },
    }
    
    c.JSON(http.StatusOK, health)
}
```

#### 9.2 Metrics Collection
```go
// เพิ่ม metrics
import "github.com/prometheus/client_golang/prometheus"

var (
    requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name: "http_request_duration_seconds",
            Help: "Duration of HTTP requests",
        },
        []string{"method", "endpoint"},
    )
    
    requestCount = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "endpoint", "status"},
    )
)
```

## 🎉 สรุป

การพัฒนาโปรเจกต์นี้ตาม Clean Architecture pattern ทำให้:

1. **Code มีโครงสร้างชัดเจน** - แต่ละ layer มีหน้าที่เฉพาะ
2. **ง่ายต่อการ Test** - ใช้ dependency injection และ mocking
3. **Maintainable** - เปลี่ยนแปลง implementation ได้ง่าย
4. **Scalable** - เพิ่ม features ใหม่ได้โดยไม่กระทบของเก่า
5. **Developer Friendly** - มี hot reload, auto swagger, mock mode

**เครื่องมือสำคัญ:**
- `dev.bat` / `make dev` - Development mode พร้อม hot reload
- `build.bat` / `make build` - Production build
- `test.bat` / `make test` - รัน tests
- Swagger UI - API documentation
- Mock Mode - Development โดยไม่ต้องมี database

**Workflow:**
1. Design Domain Entity
2. Create Interface
3. Implement Repository
4. Implement Service  
5. Create Handler
6. Add Routes
7. Create Mock
8. Write Tests
9. Generate Swagger
10. Deploy

ตาม pattern นี้จะทำให้โปรเจกต์พร้อมสำหรับ enterprise-level development! 🚀
