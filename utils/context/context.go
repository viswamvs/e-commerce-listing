package context

import (
	"e-commerce-listing/utils/db"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Context struct {
	*gin.Context
	DB    *db.DBConn
	Redis *redis.Client
}
