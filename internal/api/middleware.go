package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-API-Key")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set("request_id", requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)
		requestID := c.GetString("request_id")

		gin.DefaultWriter.Write([]byte(
			formatLog(
				requestID,
				c.Request.Method,
				c.Request.URL.Path,
				c.Writer.Status(),
				duration,
				c.ClientIP(),
			),
		))
	}
}

func APIKeyAuthMiddleware(validKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")

		if apiKey == "" {
			c.JSON(401, ErrorResponse{
				Success: false,
				Error:   "Missing API key",
				Code:    "UNAUTHORIZED",
			})
			c.Abort()
			return
		}

		if apiKey != validKey {
			c.JSON(403, ErrorResponse{
				Success: false,
				Error:   "Invalid API key",
				Code:    "FORBIDDEN",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				requestID := c.GetString("request_id")

				c.JSON(500, ErrorResponse{
					Success: false,
					Error:   "Internal server error",
					Code:    "INTERNAL_ERROR",
				})

				gin.DefaultWriter.Write([]byte(
					formatPanic(requestID, err),
				))

				c.Abort()
			}
		}()
		c.Next()
	}
}

func formatLog(requestID, method, path string, status int, duration time.Duration, ip string) string {
	return "[" + time.Now().Format("2006-01-02 15:04:05") + "] " +
		"RequestID=" + requestID + " " +
		"Method=" + method + " " +
		"Path=" + path + " " +
		"Status=" + string(rune(status)) + " " +
		"Duration=" + duration.String() + " " +
		"IP=" + ip + "\n"
}

func formatPanic(requestID string, err interface{}) string {
	return "[PANIC] RequestID=" + requestID + " Error=" + string(rune(err.(int))) + "\n"
}
