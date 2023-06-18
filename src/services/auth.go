package services

import (
	"blog/src/models"
	"blog/src/repositories"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetRequestUser(ctx *gin.Context, user *models.User) {
	session := sessions.Default(ctx)
	session.Set("userId", user.Id)
	session.Save()
}

func GetRequestUser(ctx *gin.Context) *models.User {
	session := sessions.Default(ctx)
	userId := session.Get("userId")

	if userId == nil {
		return nil
	}

	userRepository := repositories.NewUserRepository(ctx)
	user, err := userRepository.FindByID(userId.(int64))

	if err != nil {
		return nil
	}

	return user
}
