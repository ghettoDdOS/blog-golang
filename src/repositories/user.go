package repositories

import (
	"blog/src/config"
	"blog/src/models"
	"blog/src/utils"
	"errors"
	"fmt"

	"github.com/Epritka/morpheus/builder/entity"
	"github.com/gin-gonic/gin"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/dbtype"
)

type UserRepository struct {
	ctx   *gin.Context
	alias string
}

func NewUserRepository(ctx *gin.Context) *UserRepository {
	return &UserRepository{ctx: ctx, alias: "u"}
}

func (r *UserRepository) Create(user *models.User) error {
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())

	password, passwordErr := utils.MakePassword(user.Password)

	if passwordErr != nil {
		return passwordErr
	}

	conn.
		Create(entity.
			NewNode(r.alias).
			SetLables("User").SetProperties(
			map[string]any{
				"firstName":  user.FirstName,
				"lastName":   user.LastName,
				"patronymic": user.Patronymic,
				"job":        user.Job,
				"email":      user.Email,
				"password":   password,
			}),
		)

	_, err := conn.Do()

	conn.CloseWithContext(r.ctx.Request.Context())
	return err
}
func (r *UserRepository) findBy(query string) (*models.User, error) {
	db := config.GetDatabaseConnection()
	conn := db.NewExecuterWithContext(r.ctx.Request.Context())

	notFoundError := errors.New("user not found error")

	conn.
		Match(entity.NewNode(r.alias).SetLables("User")).
		Where(query).
		Return(entity.NewAlias(r.alias))

	records, err := conn.Do()
	if err != nil {
		return nil, err
	}

	if len(records) == 0 {
		return nil, notFoundError
	}

	node, found := records[0].Get(r.alias)

	if !found {
		return nil, notFoundError
	}

	userData := node.(dbtype.Node).Props
	user := models.User{
		BaseModel:  models.BaseModel{Id: node.(dbtype.Node).Id},
		FirstName:  userData["firstName"].(string),
		LastName:   userData["lastName"].(string),
		Patronymic: userData["patronymic"].(string),
		Job:        userData["job"].(string),
		Email:      userData["email"].(string),
		Password:   userData["password"].(string),
	}

	conn.CloseWithContext(r.ctx.Request.Context())
	return &user, nil
}

func (r *UserRepository) FindByID(id int64) (*models.User, error) {
	return r.findBy(fmt.Sprintf("ID(%s) = %d", r.alias, id))
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	return r.findBy(fmt.Sprintf(`%s.email = "%s"`, r.alias, email))
}
