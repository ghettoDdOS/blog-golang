package controllers

import (
	"blog/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func homeController(ctx *gin.Context) {
	settings := config.GetSettings()

	ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
		"settings": settings,
	})
}
