package middlewares

import (
	"blog/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitMiddlewares(router *gin.Engine) {
	store := cookie.NewStore([]byte(config.GetSettings().SecretKey))
	router.Use(sessions.Sessions("session", store))
}
