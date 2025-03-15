package context

import (
	"e-commerce-listing/utils/db"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	DB *db.DBConn
}
