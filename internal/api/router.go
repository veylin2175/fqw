package api

import (
	"github.com/gin-gonic/gin"

	"fqw/internal/blockchain"
)

type RouterConfig struct {
	Client *blockchain.Client
	APIKey string
}

func SetupRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.New()

	router.Use(RecoveryMiddleware())
	router.Use(CORSMiddleware())
	router.Use(RequestIDMiddleware())
	router.Use(LoggingMiddleware())

	handler := NewHandler(cfg.Client)

	router.GET("/health", handler.HealthCheck)

	v1 := router.Group("/api/v1")
	{
		verification := v1.Group("/verify")
		{
			verification.POST("", handler.VerifyDocument)
			verification.GET("/:tokenId", handler.CheckVerification)
			verification.DELETE("/:tokenId", handler.RevokeVerification)
		}

		v1.GET("/system/info", handler.GetSystemInfo)

		admin := v1.Group("/admin")
		admin.Use(APIKeyAuthMiddleware(cfg.APIKey))
		{
			admin.POST("/issuer", handler.AddIssuer)
		}
	}

	return router
}
