package middleware

import (
	"e-commerce-listing/utils/context"
	"e-commerce-listing/utils/db"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func AuthMiddleware(next func(*context.Context), redisClient *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := &context.Context{
			Context: c,
			Redis:   redisClient,
		}

		logAndGetContext(ctx)

		next(ctx)
	}
}

func logAndGetContext(c *context.Context) {
	c.DB = db.New()
}
