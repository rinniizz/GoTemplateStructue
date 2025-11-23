package middleware

import (
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

// IPWhitelist allows only specific IPs to access the endpoint
func IPWhitelist(allowedIPs []string) gin.HandlerFunc {
	allowedMap := make(map[string]bool)
	for _, ip := range allowedIPs {
		allowedMap[ip] = true
	}

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if !allowedMap[clientIP] {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Access denied from your IP address",
				"error":   "ip_not_allowed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// IPBlacklist blocks specific IPs from accessing the endpoint
func IPBlacklist(blockedIPs []string) gin.HandlerFunc {
	blockedMap := make(map[string]bool)
	for _, ip := range blockedIPs {
		blockedMap[ip] = true
	}

	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		if blockedMap[clientIP] {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Your IP address has been blocked",
				"error":   "ip_blocked",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// IPRangeFilter filters by CIDR range (e.g., "192.168.1.0/24")
func IPRangeFilter(allowedRanges []string) gin.HandlerFunc {
	var cidrs []*net.IPNet
	for _, cidr := range allowedRanges {
		_, ipNet, err := net.ParseCIDR(cidr)
		if err == nil {
			cidrs = append(cidrs, ipNet)
		}
	}

	return func(c *gin.Context) {
		clientIP := net.ParseIP(c.ClientIP())
		if clientIP == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Invalid IP address",
				"error":   "invalid_ip",
			})
			c.Abort()
			return
		}

		allowed := false
		for _, ipNet := range cidrs {
			if ipNet.Contains(clientIP) {
				allowed = true
				break
			}
		}

		if !allowed {
			c.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Access denied from your IP range",
				"error":   "ip_range_not_allowed",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
