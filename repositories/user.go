package repositories

import (
	"blog/models"
	"fmt"

	"github.com/Epritka/morpheus/ogm"
)

type UserRepository struct {
	db *ogm.Driver
}

func NewUserRepository(db *ogm.Driver) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.User) error {
	query := `
		CREATE (u:User {
			id: %d,
			firstName: %s,
			lastName: %s,
			patronymic: %s,
			job: %s,
			email: %s,
			password: %s,
		})
	`
	conn := r.db.NewExecuter()
	_, err := conn.DoQuery(fmt.Sprintf(
		query,
		user.Id,
		user.FirstName,
		user.LastName,
		user.Patronymic,
		user.Job,
		user.Email,
		user.Password,
	))
	return err
}

// func (r *UserRepository) FindByID(id uint) (*models.User, error) {
// 	var user User
// 	err := r.db.First(&user, id).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *UserRepository) Update(user *models.User) error {
// 	return r.db.Save(user).Error
// }

// func (r *UserRepository) Delete(user *models.User) error {
// 	return r.db.Delete(user).Error
// }
