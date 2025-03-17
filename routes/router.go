package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func GetRouter(redisClient *redis.Client) *gin.Engine {

	router := gin.New()

	v1 := router.Group("/v1")
	ProductRoutes(v1, redisClient)

	return router
}
