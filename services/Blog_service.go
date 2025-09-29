package services

import (
	"fmt"

	"example.com/net-http-class/models"
	"example.com/net-http-class/utils"
	"gorm.io/gorm"
)

type Blogservice struct {
	DB *gorm.DB
}

func NewBlogservice(db *gorm.DB) *Blogservice {
	return &Blogservice{DB: db}
}

func (t *Blogservice) CreateBlogPost(token string, post *models.Blogrequest) error {
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return fmt.Errorf("invalid jwt token: %w", err)
	}

	User_id := claims.User_ID

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

func (t *Blogservice) FetchBlogPost(token string) ([]models.Blog, error) {
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return nil, fmt.Errorf("invalid jwt token: %w", err)
	}

	var Post []models.Blog

	user_id := claims.User_ID

	result := t.DB.Where("user_id = ?", user_id).Find(&Post)
	if result.Error != nil {
		return nil, fmt.Errorf("couldn't find user: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		// No posts found
		return nil, fmt.Errorf("no posts found for user %s", user_id)
	}

	return Post, nil
}
