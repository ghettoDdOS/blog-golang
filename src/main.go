package main

import (
	"blog/src/config"
	"blog/src/controllers"
	"blog/src/middlewares"
	"blog/src/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	config.InitSettings()
	middlewares.InitMiddlewares(router)
	router.LoadHTMLGlob("src/templates/*")
	controllers.InitRoutes(router)
	models.InitConstraints()

	serverSettings := config.GetSettings().Server
	router.Run(fmt.Sprintf(":%d", serverSettings.Port))
}
