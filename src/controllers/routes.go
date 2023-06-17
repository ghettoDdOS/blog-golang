package controllers

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	router.Static("/static", "./static")

	router.GET("/", homeController)
	router.POST("/login", loginController)
	router.POST("/registration", registrationController)
}
