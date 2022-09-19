package services

import (
	"golang-user-service/src/account"
	"golang-user-service/src/account/models"
)

type UserServiceImpl struct {
	userRepo account.UserRepo
}

func CreatePersonUsecase(userRepo account.UserRepo) account.UserService {
	return &UserServiceImpl{userRepo}
}

func (e *UserServiceImpl) Create(user *models.User) (*models.User, error) {
	return e.userRepo.Create(user)
}

func (e *UserServiceImpl) ReadById(id string) (*models.User, error) {
	return e.userRepo.ReadById(id)
}

func (e *UserServiceImpl) ReadByUsername(username string) (*models.User, error) {
	return e.userRepo.ReadByUsername(username)
}
