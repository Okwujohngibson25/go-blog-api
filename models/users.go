package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Fullname  string    `json:"fullname"`
	Username  string    `json:"username"`
	Email     string    `json:"email" gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Password  string         `json:"password"`

	Blog []Blog `gorm:"foreignKey:UserID"` // One-to-many relation
}
