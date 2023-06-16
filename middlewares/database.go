package middlewares

import (
	"blog/config"

	"github.com/gin-gonic/gin"
)

func Database() gin.HandlerFunc {
	return func(c *gin.Context) {
		database := config.NewDatabaseConnection()
		c.Set("database", database)
		c.Next()
	}
}
