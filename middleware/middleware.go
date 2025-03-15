package middleware

import (
	"e-commerce-listing/utils/context"
	"e-commerce-listing/utils/db"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(next func(*context.Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &context.Context{
			Context: c,
		}

		logAndGetContext(ctx)

		next(ctx)
	}
}

func logAndGetContext(c *context.Context) {
	c.DB = db.New()
}
