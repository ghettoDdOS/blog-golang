package controllers

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	router.Static("/static", "./src/static")

	router.GET("/", homePageController)

	router.GET("/login", loginPageController)
	router.POST("/login", loginController)
	router.POST("/logout", logoutController)

	router.GET("/registration", registrationPageController)
	router.POST("/registration", registrationController)

	router.GET("/add-article", addArticlePageController)
	router.POST("/add-article", addArticleController)
	router.POST("/delete-article/:articleId", deleteArticleController)

	router.GET("/personal-area", personalAreaPageController)
}
