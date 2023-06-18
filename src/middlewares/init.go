package middlewares

import (
	"blog/src/config"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func InitMiddlewares(router *gin.Engine) {
	store := cookie.NewStore([]byte(config.GetSettings().SecretKey))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: int(30 * 24 * 60 * 60),
	})
	router.Use(sessions.Sessions("mysession", store))

}
