package routes

import (
	"e-commerce-listing/handlers"
	mw "e-commerce-listing/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func ProductRoutes(router *gin.RouterGroup, redisClient *redis.Client) {
	router.GET("/products", mw.AuthMiddleware(handlers.GetProducts, redisClient))
	router.GET("/products/:id", mw.AuthMiddleware(handlers.GetProduct, redisClient))
	router.POST("/products", mw.AuthMiddleware(handlers.SaveProduct, redisClient))
	router.PUT("/products/:id", mw.AuthMiddleware(handlers.SaveProduct, redisClient))
	router.DELETE("/products/:id", mw.AuthMiddleware(handlers.DeleteProduct, redisClient))
}
