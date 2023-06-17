package controllers

import (
	"blog/src/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeController(ctx *gin.Context) {
	settings := config.GetSettings()

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"settings": settings,
	})
}
