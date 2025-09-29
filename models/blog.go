package models

import (
	"time"

	"github.com/google/uuid"
)

type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title     string    `json:"title"`
	Post      string    `json:"post"`
	UserID    uuid.UUID `gorm:"type:uuid;default:NULL"` // Foreign key field
	CreatedAt time.Time
}

type Blogrequest struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}
