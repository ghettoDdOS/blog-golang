package controllers

import (
	"blog/src/models"
	"blog/src/repositories"
	"blog/src/services"
	"blog/src/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func addArticleController(ctx *gin.Context) {
	articleRepository := repositories.NewArticleRepository(ctx)
	user := services.GetRequestUser(ctx)
	var article models.Article
	if err := ctx.ShouldBind(&article); err != nil {
		utils.RenderTemplateResponse(
			ctx,
			"add-article.html",
			map[string]interface{}{"Error": "Ошибка сервера"},
		)
		return
	}

	err := articleRepository.Create(&article, user.Id)
	if err != nil {
		utils.RenderTemplateResponse(
			ctx,
			"add-article.html",
			map[string]interface{}{"Error": "Ошибка сервера"},
		)
	}
	ctx.Redirect(http.StatusFound, "/")
}

func addArticlePageController(ctx *gin.Context) {
	user := services.GetRequestUser(ctx)

	if user == nil {
		ctx.Redirect(http.StatusFound, "/login")
	}
	utils.RenderTemplateResponse(
		ctx,
		"add-article.html",
		map[string]interface{}{},
	)
}
