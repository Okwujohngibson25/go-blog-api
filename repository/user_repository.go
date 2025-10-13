package repository

import (
	"fmt"

	"example.com/net-http-class/models"
	"gorm.io/gorm"
)

type Userrepository struct {
	db *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *Userrepository {
	return &Userrepository{db: DB}
}

type UserRepository interface {
	Create(user *models.Users) error
	FindUserByEmail(email string) (*models.Users, error)
}

func (n *Userrepository) Create(user *models.Users) error {
	result := n.db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("user creation failed: %w", result.Error)
	}
	return nil
}

func (n *Userrepository) FindUserByEmail(email string) (*models.Users, error) {
	var user models.Users
	result := n.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("couldn't find user: %w", result.Error)
	}
	return &user, nil
}
