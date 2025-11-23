package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type rateLimiter struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var (
	visitors = make(map[string]*rateLimiter)
	mu       sync.Mutex
)

// RateLimiter creates a rate limiting middleware
// rps: requests per second (e.g., 10 = 10 requests per second)
// burst: maximum burst size (e.g., 20 = allow burst of 20 requests)
func RateLimiter(rps int, burst int) gin.HandlerFunc {
	limit := rate.Limit(rps)
	return func(c *gin.Context) {
		ip := c.ClientIP()

		mu.Lock()
		limiter, exists := visitors[ip]
		if !exists {
			limiter = &rateLimiter{
				limiter:  rate.NewLimiter(limit, burst),
				lastSeen: time.Now(),
			}
			visitors[ip] = limiter
		}
		limiter.lastSeen = time.Now()
		mu.Unlock()

		if !limiter.limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"success": false,
				"message": "Rate limit exceeded. Please try again later.",
				"error":   "too_many_requests",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// CleanupVisitors removes old entries every 5 minutes to prevent memory leak
func CleanupVisitors() {
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			mu.Lock()
			for ip, limiter := range visitors {
				if time.Since(limiter.lastSeen) > 5*time.Minute {
					delete(visitors, ip)
				}
			}
			mu.Unlock()
		}
	}()
}
