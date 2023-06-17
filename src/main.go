package main

import (
	"blog/src/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	config.InitSettings()
	// middlewares.InitMiddlewares(router)
	// router.LoadHTMLGlob("templates/*")
	// controllers.InitRoutes(router)
	fmt.Println(config.GetSettings().Database.Uri)

	router.Run(":3001")
}
