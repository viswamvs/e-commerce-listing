package handlers

import (
	"e-commerce-listing/dtos"
	"e-commerce-listing/services/product"
	"e-commerce-listing/utils/context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func GetProducts(c *context.Context) {
	cacheKey := "products"

	val, err := c.Redis.Get(c.Request.Context(), cacheKey).Result()
	if err == redis.Nil {
		products, err := product.NewProductService().GetProducts(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		jsonData, _ := json.Marshal(products)
		c.Redis.Set(c.Request.Context(), cacheKey, jsonData, 10*time.Minute)

		c.JSON(http.StatusOK, products)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var products []*dtos.Product
	json.Unmarshal([]byte(val), &products)
	c.JSON(http.StatusOK, products)
}

func GetProduct(c *context.Context) {
	productID := c.Param("id")
	cacheKey := fmt.Sprintf("product:%s", productID)

	val, err := c.Redis.Get(c.Request.Context(), cacheKey).Result()
	if err == redis.Nil {
		product, err := product.NewProductService().GetProduct(c, productID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		jsonData, _ := json.Marshal(product)
		c.Redis.Set(c.Request.Context(), cacheKey, jsonData, 10*time.Minute)

		c.JSON(http.StatusOK, product)
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var product dtos.Product
	json.Unmarshal([]byte(val), &product)
	c.JSON(http.StatusOK, product)
}

func SaveProduct(c *context.Context) {
	req := &dtos.Product{}

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = product.NewProductService().SaveProducts(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if req.Id > 0 {
		cacheKey := fmt.Sprintf("product:%d", req.Id)

		productJSON, err := json.Marshal(req)
		if err == nil {
			c.Redis.Set(c.Request.Context(), cacheKey, productJSON, 10*time.Minute)
		}
		c.Redis.Del(c.Request.Context(), "products")
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func DeleteProduct(c *context.Context) {
	productID := c.Param("id")
	cacheKey := fmt.Sprintf("product:%s", productID)

	err := product.NewProductService().DeleteProduct(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Redis.Del(c.Request.Context(), cacheKey)
	c.Redis.Del(c.Request.Context(), "products")

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
