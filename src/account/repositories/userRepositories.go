package repo

import (
	"fmt"
	"golang-user-service/src/account"
	"golang-user-service/src/account/models"

	"github.com/jinzhu/gorm"
)

type UserRepoImpl struct {
	DB *gorm.DB
}

func CreatePersonRepo(DB *gorm.DB) account.UserRepo {
	return &UserRepoImpl{DB}
}

func (e *UserRepoImpl) Create(user *models.User) (*models.User, error) {
	err := e.DB.Save(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.Create] error execute query %v \n", err)
		return nil, fmt.Errorf("failed insert data")
	}
	return user, nil
}

func (e *UserRepoImpl) ReadById(id string) (*models.User, error) {
	var user = models.User{}
	err := e.DB.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ReadById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}
	return &user, nil
}

func (e *UserRepoImpl) ReadByUsername(username string) (*models.User, error) {
	var user = models.User{}
	err := e.DB.Table("users").Where("username = ?", username).First(&user).Error
	if err != nil {
		fmt.Printf("[UserRepoImpl.ReadByUsername] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}
	return &user, nil
}
