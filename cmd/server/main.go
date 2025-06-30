package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-template-structure/internal/config"
	"go-template-structure/internal/handler"
	"go-template-structure/internal/middleware"
	"go-template-structure/internal/mock"
	"go-template-structure/internal/repository"
	"go-template-structure/internal/service"
	"go-template-structure/pkg/database"
	"go-template-structure/pkg/logger"

	_ "go-template-structure/docs" // swagger docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize logger
	logger.Init(cfg.LogLevel, cfg.LogFormat)
	logger.Info("Starting application...")

	// Check if we should use mock mode (when DB connection fails or MOCK_MODE=true)
	useMockMode := os.Getenv("MOCK_MODE") == "true"

	var userRepo repository.UserRepository
	var redisClient interface{}

	if !useMockMode {
		// Try to initialize real database
		db, err := database.NewPostgresConnection(cfg.Database)
		if err != nil {
			logger.Warn("Failed to connect to database, using mock mode:", err)
			useMockMode = true
		} else {
			userRepo = repository.NewUserRepository(db)
			logger.Info("Connected to PostgreSQL database")
		}

		// Try to initialize Redis
		if !useMockMode {
			redisClient, err = database.NewRedisConnection(cfg.Redis)
			if err != nil {
				logger.Warn("Failed to connect to Redis, using mock client:", err)
				redisClient = mock.NewMockRedisClient()
			} else {
				logger.Info("Connected to Redis")
			}
		}
	}

	// Use mock mode if needed
	if useMockMode {
		logger.Info("🔧 Running in MOCK MODE - no database required")
		userRepo = mock.NewMockUserRepository()
		redisClient = mock.NewMockRedisClient()
	}

	// Initialize services
	userService := service.NewUserService(userRepo, redisClient, cfg.JWT)
	authService := service.NewAuthService(userRepo, cfg.JWT)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	// Setup router
	router := setupRouter(cfg, userHandler, authHandler, useMockMode)

	// Setup server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}

	// Start server in goroutine
	go func() {
		mode := "PRODUCTION"
		if useMockMode {
			mode = "MOCK"
		}

		logger.Info(fmt.Sprintf("🚀 Server is running on %s:%s (%s MODE)", cfg.Server.Host, cfg.Server.Port, mode))
		logger.Info("📚 Swagger UI: http://localhost:" + cfg.Server.Port + "/swagger/index.html")
		logger.Info("🏥 Health Check: http://localhost:" + cfg.Server.Port + "/health")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server:", err)
		}
	}()

	// Graceful shutdown
	gracefulShutdown(srv)
}

func setupRouter(cfg *config.Config, userHandler *handler.UserHandler, authHandler *handler.AuthHandler, mockMode bool) *gin.Engine {
	// Set Gin mode
	gin.SetMode(cfg.Server.GinMode)

	router := gin.New()

	// Middlewares
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())

	// Health check
	router.GET("/health", func(c *gin.Context) {
		status := "ok"
		mode := "production"
		if mockMode {
			mode = "mock"
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  status,
			"mode":    mode,
			"message": fmt.Sprintf("Server is running in %s mode", mode),
			"time":    time.Now().UTC(),
		})
	})

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	v1 := router.Group("/api/v1")
	{
		// Auth routes
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
			auth.POST("/refresh", authHandler.RefreshToken)
		}

		// Protected routes
		protected := v1.Group("/")
		protected.Use(middleware.JWTAuth(cfg.JWT.Secret))
		{
			// User routes
			users := protected.Group("/users")
			{
				users.GET("/profile", userHandler.GetProfile)
				users.PUT("/profile", userHandler.UpdateProfile)
				users.GET("/", userHandler.GetUsers)
				users.GET("/:id", userHandler.GetUser)
				users.PUT("/:id", userHandler.UpdateUser)
				users.DELETE("/:id", userHandler.DeleteUser)
			}
		}
	}

	return router
}

func gracefulShutdown(srv *http.Server) {
	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("🛑 Shutting down server...")

	// The context is used to inform the server it has 30 seconds to finish
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown:", err)
	}

	logger.Info("✅ Server exited")
}
