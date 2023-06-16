package main

import (
	"blog/config"
	"blog/middlewares"
	"blog/models"
	"blog/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	config.InitSettings()

	router.Use(middlewares.Database())

	router.GET("/", func(ctx *gin.Context) {
		settings := config.GetSettings()

		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"settings": settings,
		})
	})

	router.POST("/registration", func(ctx *gin.Context) {
		var user models.User
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	})

	pass1 := utils.MakePassword("penEg1342")
	pass2 := utils.MakePassword("penEg13421")
	fmt.Println(utils.CheckPassword(pass1, pass2))
	fmt.Println(utils.CheckPassword(pass1, pass1))
	router.Run(":8080")
}
