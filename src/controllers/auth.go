package controllers

import (
	"blog/src/models"
	"blog/src/repositories"
	"blog/src/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func loginController(ctx *gin.Context) {
	session := sessions.Default(ctx)
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")

	userRepository := repositories.NewUserRepository(ctx)
	user, err := userRepository.FindByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !utils.CheckPassword(password, user.Password) || user == nil {
		// FAil
		return
	}

	session.Set("user_id", user.Id)
	session.Save()
	return
}

func registrationController(ctx *gin.Context) {
	userRepository := repositories.NewUserRepository(ctx)
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// userRepository.Create(&user)
	userRepository.FindByID(1)

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}
