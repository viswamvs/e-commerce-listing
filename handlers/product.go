package handlers

import (
	"e-commerce-listing/dtos"
	"e-commerce-listing/services/product"
	"e-commerce-listing/utils/context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *context.Context) {

	products, err := product.NewProductService().GetProducts(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProduct(c *context.Context) {

	product, err := product.NewProductService().GetProduct(c, c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func SaveProduct(c *context.Context) {

	req := &dtos.Product{}

	err := c.BindJSON(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = product.NewProductService().SaveProducts(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func DeleteProduct(c *context.Context) {

	err := product.NewProductService().DeleteProduct(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
