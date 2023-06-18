package controllers

import (
	"blog/src/models"
	"blog/src/repositories"
	"blog/src/services"
	"blog/src/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ArticleFilters struct {
	Title             string
	FullName          string
	Keywords          []string
	Subject           string
	YearOfPublication int
}

func homePageController(ctx *gin.Context) {
	articleRepository := repositories.NewArticleRepository(ctx)

	var filters ArticleFilters
	if err := ctx.ShouldBindQuery(&filters); err != nil {
		utils.RenderTemplateResponse(
			ctx,
			"index.html",
			map[string]interface{}{"Error": "Ошибка сервера"},
		)
		return
	}
	filters.Keywords = ctx.QueryArray("Keywords[]")
	fmt.Println(filters)

	var articles []*models.Article
	var err error

	if filters.Title != "" {
		articles, err = articleRepository.GetByTitle(filters.Title)
	} else if filters.FullName != "" {
		articles, err = articleRepository.GetByFullName(filters.FullName)
	} else if len(filters.Keywords) > 0 {
		articles, err = articleRepository.GetByKeywords(filters.Keywords)
	} else if filters.Subject != "" {
		articles, err = articleRepository.GetBySubject(filters.Subject)
	} else if filters.YearOfPublication > 0 {
		articles, err = articleRepository.GetByYearOfPublication(filters.YearOfPublication)
	} else {
		articles, err = articleRepository.Get()
	}

	if err != nil {
		utils.RenderTemplateResponse(
			ctx,
			"index.html",
			map[string]interface{}{"Error": "Ошибка сервера"},
		)
	}

	aviableKeywords, err := articleRepository.GetAviableKeywords()
	if err != nil {
		utils.RenderTemplateResponse(
			ctx,
			"index.html",
			map[string]interface{}{"Error": "Ошибка сервера"},
		)
	}

	user := services.GetRequestUser(ctx)
	utils.RenderTemplateResponse(ctx,
		"index.html",
		map[string]interface{}{
			"User":            user,
			"Articles":        articles,
			"AviableKeywords": aviableKeywords,
			"Filters":         filters,
		},
	)
}
