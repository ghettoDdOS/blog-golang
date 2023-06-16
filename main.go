package main

import (
	"blog/config"
	"blog/controllers"
	"blog/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	config.InitSettings()
	middlewares.InitMiddlewares(router)
	router.LoadHTMLGlob("templates/*")
	controllers.InitRoutes(router)

	router.Run(":3001")
}
