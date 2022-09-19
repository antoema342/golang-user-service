package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         uuid.UUID  `gorm:"column:id;type:uuid;primary_key;not null;default:uuid_generate_v4();" json:"id"`
	Username   string     `gorm:"column:username;unique" json:"username" binding:"required,email"`
	Password   string     `gorm:"column:password" json:"password" binding:"required,min=12"`
	Name       string     `gorm:"name:name" json:"name" binding:"required"`
	IsActive   bool       `gorm:"column:is_active" json:"is_active"`
	CreatedAt  time.Time  `gorm:"column:created_at" json:"created_at"`
	ModifiedAt time.Time  `gorm:"column:modified_at" json:"modified_at"`
	DeletedAt  *time.Time `gorm:"index" json:"deleted_at,omitempty"`
}

type SignIn struct {
	Username string `binding:"required,email"`
	Password string `binding:"required,min=12"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
