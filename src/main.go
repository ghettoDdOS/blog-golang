package main

import (
	"blog/src/config"
	"blog/src/controllers"
	"blog/src/middlewares"
	"blog/src/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.InitSettings()
	middlewares.InitMiddlewares(router)
	controllers.InitRoutes(router)
	config.InitDatabase()
	models.InitConstraints()

	router.Run(":8080")
}
