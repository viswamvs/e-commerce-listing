package routes

import (
	"e-commerce-listing/handlers"
	mw "e-commerce-listing/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.RouterGroup) {
	router.GET("/products", mw.AuthMiddleware(handlers.GetProducts))
	router.GET("/products/:id", mw.AuthMiddleware(handlers.GetProduct))
	router.POST("/products", mw.AuthMiddleware(handlers.SaveProduct))
	router.PUT("/products/:id", mw.AuthMiddleware(handlers.SaveProduct))
	router.DELETE("/products/:id", mw.AuthMiddleware(handlers.DeleteProduct))
}
