package repositories

import (
	"blog/config"
	"blog/models"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserRepository struct {
	ctx *gin.Context
}

func NewUserRepository(ctx *gin.Context) *UserRepository {
	return &UserRepository{ctx: ctx}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
		CREATE (u:User {
			firstName: "%s",
			lastName: "%s",
			patronymic: "%s",
			job: "%s",
			email: "%s",
			password: "%s"
		})
	`
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())
	_, err := conn.DoQuery(fmt.Sprintf(
		query,
		user.FirstName,
		user.LastName,
		user.Patronymic,
		user.Job,
		user.Email,
		user.Password,
	))
	conn.CloseWithContext(r.ctx.Request.Context())
	return err
}

func (r *UserRepository) FindByID(id uint) (*models.User, error) {
	// query := "MATCH (n:User) WHERE ID(n) = %d RETURN n"

	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())
	result, err := conn.DoQuery("MATCH (n:User) RETURN n")
	if err != nil {
		return nil, err
	}
	fmt.Println(result.Next(context.Background()))
	// if res {
	// 	// var user models.User
	// 	record := result.Record()
	// 	fmt.Println(record.Values...)
	// 	// user.ID = record.GetByIndex(0).(int)
	// 	// user.Name = record.GetByIndex(1).(string)
	// 	// user.Email = record.GetByIndex(2).(string)

	// }
	conn.CloseWithContext(r.ctx.Request.Context())
	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	// query := "MATCH (n:User) WHERE ID(n) = %d RETURN n"

	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())
	result, err := conn.DoQuery("MATCH (n:User) RETURN n")
	if err != nil {
		return nil, err
	}
	fmt.Println(result.Next(context.Background()))
	// if res {
	// 	// var user models.User
	// 	record := result.Record()
	// 	fmt.Println(record.Values...)
	// 	// user.ID = record.GetByIndex(0).(int)
	// 	// user.Name = record.GetByIndex(1).(string)
	// 	// user.Email = record.GetByIndex(2).(string)

	// }
	conn.CloseWithContext(r.ctx.Request.Context())
	return nil, nil
}

// func (r *UserRepository) Update(user *models.User) error {
// 	return r.db.Save(user).Error
// }

// func (r *UserRepository) Delete(user *models.User) error {
// 	return r.db.Delete(user).Error
// }
