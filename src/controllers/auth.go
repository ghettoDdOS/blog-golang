package controllers

import (
	"blog/src/models"
	"blog/src/repositories"
	"blog/src/services"
	"blog/src/utils"
	"fmt"

	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func loginController(ctx *gin.Context) {
	email := ctx.PostForm("Email")
	password := ctx.PostForm("Password")

	userRepository := repositories.NewUserRepository(ctx)
	user, err := userRepository.FindByEmail(email)
	if err != nil {
		utils.RenderTemplateResponse(
			ctx,
			"login.html",
			map[string]interface{}{"Error": "Пользователь не найден"},
		)
		return
	}
	if !utils.CheckPassword(password, user.Password) || user == nil {
		utils.RenderTemplateResponse(
			ctx,
			"login.html",
			map[string]interface{}{"Error": "Пользователь не найден"},
		)
		return
	}
	services.SetRequestUser(ctx, user)

	ctx.Redirect(http.StatusFound, "/")
}

func loginPageController(ctx *gin.Context) {
	user := services.GetRequestUser(ctx)

	if user != nil {
		ctx.Redirect(http.StatusFound, "/")
	}
	utils.RenderTemplateResponse(
		ctx,
		"login.html",
		map[string]interface{}{},
	)
}

func registrationController(ctx *gin.Context) {
	userRepository := repositories.NewUserRepository(ctx)
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		utils.RenderTemplateResponse(
			ctx,
			"registration.html",
			map[string]interface{}{"Error": "Ошибка сервера"},
		)
		return
	}
	err := userRepository.Create(&user)
	if err != nil {
		fmt.Println(err)
		utils.RenderTemplateResponse(
			ctx,
			"registration.html",
			map[string]interface{}{"Error": "Ошибка сервера"},
		)
	}
	ctx.Redirect(http.StatusFound, "/login")
}

func registrationPageController(ctx *gin.Context) {
	user := services.GetRequestUser(ctx)

	if user != nil {
		ctx.Redirect(http.StatusFound, "/")
	}
	utils.RenderTemplateResponse(
		ctx,
		"registration.html",
		map[string]interface{}{},
	)
}

func logoutController(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("userId")
	session.Save()

	ctx.Redirect(http.StatusFound, "/login")
}
