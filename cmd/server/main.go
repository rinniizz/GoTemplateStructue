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
	"go-template-structure/internal/interfaces"
	"go-template-structure/internal/middleware"
	"go-template-structure/internal/repository"
	"go-template-structure/internal/service"
	"go-template-structure/pkg/database"
	"go-template-structure/pkg/logger"

	_ "go-template-structure/docs" // swagger docs

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	// Initialize primary database connection (required)
	db, err := database.NewPostgresConnection(cfg.Database)
	if err != nil {
		logger.Fatal("Failed to connect to database:", err)
	}
	userRepo := repository.NewUserRepository(db)
	logger.Info("Connected to PostgreSQL database")

	// Initialize Redis cache (optional but recommended)
	var redisClient interfaces.RedisInterface
	if client, err := database.NewRedisConnection(cfg.Redis); err != nil {
		logger.Warn("Failed to connect to Redis, caching disabled:", err)
	} else {
		redisClient = client
		logger.Info("Connected to Redis")
	}

	// Initialize services
	userService := service.NewUserService(userRepo, redisClient, cfg.JWT)
	authService := service.NewAuthService(userRepo, cfg.JWT)

	// Initialize handlers
	userHandler := handler.NewUserHandler(userService)
	authHandler := handler.NewAuthHandler(authService)

	// Setup router
	router := setupRouter(cfg, userHandler, authHandler)

	// Setup server
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler: router,
	}

	// Start server in goroutine
	go func() {
		logger.Info(fmt.Sprintf("🚀 Server is running on %s:%s", cfg.Server.Host, cfg.Server.Port))
		logger.Info("📚 Swagger UI: http://localhost:" + cfg.Server.Port + "/swagger/index.html")
		logger.Info("🏥 Health Check: http://localhost:" + cfg.Server.Port + "/health")

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server:", err)
		}
	}()

	// Graceful shutdown
	gracefulShutdown(srv)
}

func setupRouter(cfg *config.Config, userHandler *handler.UserHandler, authHandler *handler.AuthHandler) *gin.Engine {
	// Set Gin mode
	gin.SetMode(cfg.Server.GinMode)

	router := gin.New()

	// Start rate limiter cleanup routine
	middleware.CleanupVisitors()

	// Security Middlewares (ordered by priority)
	router.Use(middleware.RequestID())                                         // 1. Request tracking
	router.Use(middleware.PrometheusMetrics())                                 // 2. Prometheus metrics collection
	router.Use(middleware.SecurityHeaders())                                   // 3. Security headers (XSS, Clickjacking protection)
	router.Use(middleware.RateLimiter(cfg.RateLimit.RPS, cfg.RateLimit.Burst)) // 4. Rate limiting (configurable)
	router.Use(middleware.CORS())                                              // 5. CORS policy
	router.Use(middleware.Logger())                                            // 6. Request/Response logging
	router.Use(middleware.AuditLog())                                          // 7. Security audit logging
	router.Use(middleware.Recovery())                                          // 8. Panic recovery

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": fmt.Sprintf("Server is running in %s mode", cfg.Server.GinMode),
			"time":    time.Now().UTC(),
		})
	})

	// Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

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
