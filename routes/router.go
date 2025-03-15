package routes

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {

	router := gin.New()

	v1 := router.Group("/v1")
	ProductRoutes(v1)

	return router
}
