package repository

import "gorm.io/gorm"

type Userrepository struct {
	db *gorm.DB
}

func NewUserRepository(DB *gorm.DB) *Userrepository {
	return &Userrepository{db: DB}
}
