package controllers

import (
	"blog/src/repositories"
	"blog/src/services"
	"blog/src/utils"

	"github.com/gin-gonic/gin"
)

func personalAreaPageController(ctx *gin.Context) {
	user := services.GetRequestUser(ctx)
	articleRepository := repositories.NewArticleRepository(ctx)

	articles, err := articleRepository.GetByUserId(user.Id)

	if err != nil {
		utils.RenderTemplateResponse(
			ctx,
			"personal-area.html",
			map[string]interface{}{"Error": "Ошибка сервера"},
		)
	}

	utils.RenderTemplateResponse(ctx,
		"personal-area.html",
		map[string]interface{}{
			"User":     user,
			"Articles": articles,
		},
	)
}
