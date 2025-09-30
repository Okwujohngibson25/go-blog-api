package services

import (
	"fmt"

	"example.com/net-http-class/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Blogservice struct {
	DB *gorm.DB
}

func NewBlogservice(db *gorm.DB) *Blogservice {
	return &Blogservice{DB: db}
}

func (t *Blogservice) CreateBlogPost(userid uuid.UUID, post *models.Blogrequest) error {
	User_id := userid

	blog := models.Blog{
		Title:  post.Title,
		Post:   post.Post,
		UserID: User_id,
	}

	result := t.DB.Create(&blog)
	if result.Error != nil {
		return fmt.Errorf("user creation failed: %w", result.Error)
	}
	return nil
}

func (t *Blogservice) FetchBlogPost(userid uuid.UUID) ([]models.Blog, error) {
	var Post []models.Blog

	User_id := userid

	result := t.DB.Where("user_id = ?", User_id).Find(&Post)
	if result.Error != nil {
		return nil, fmt.Errorf("couldn't find user: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		// No posts found
		return nil, fmt.Errorf("no posts found for user %s", User_id)
	}

	return Post, nil
}
