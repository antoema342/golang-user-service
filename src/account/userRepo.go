package account

import "golang-user-service/src/account/models"

type UserRepo interface {
	Create(user *models.User) (*models.User, error)
	ReadById(id string) (*models.User, error)
	ReadByUsername(username string) (*models.User, error)
}
