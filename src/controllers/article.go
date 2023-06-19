package controllers

import (
	"blog/src/models"
	"blog/src/repositories"
	"blog/src/services"
	"blog/src/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
		map[string]interface{}{
			"User": user,
		},
	)
}

func deleteArticleController(ctx *gin.Context) {
	articleRepository := repositories.NewArticleRepository(ctx)
	user := services.GetRequestUser(ctx)
	if user == nil {
		ctx.Redirect(http.StatusFound, "/login")
	}

	articleId, err := strconv.Atoi(ctx.Param("articleId"))
	if err != nil {
		log.Println(fmt.Errorf("failed to delete article: %s", err))
		return
	}
	err = articleRepository.DeleteById(articleId)
	if err != nil {
		log.Println(fmt.Errorf("failed to delete article: %s", err))
		return
	}

	ctx.Redirect(http.StatusFound, "/personal-area")
}

func userArticlesPageController(ctx *gin.Context) {
	user := services.GetRequestUser(ctx)

	if user == nil {
		ctx.Redirect(http.StatusFound, "/login")
	}

	var articleId int64
	var err error

	articleId, err = strconv.ParseInt(ctx.Param("articleId"), 10, 64)
	if err != nil {
		log.Println(fmt.Errorf("failed to parse id: %s", err))
		return
	}

	articleRepository := repositories.NewArticleRepository(ctx)
	author, err := articleRepository.GetAuthor(articleId)
	if err != nil || user == nil {
		log.Println(fmt.Errorf("failed to found user: %s", err))
		return
	}

	articles, err := articleRepository.GetByUserId(author.Id)
	if err != nil {
		log.Println(fmt.Errorf("failed to found articles: %s", err))
		return
	}

	utils.RenderTemplateResponse(
		ctx,
		"articles.html",
		map[string]interface{}{
			"User":     user,
			"Author":   author,
			"Articles": articles,
		},
	)
}
