package repository

import (
	"errors"
	"fmt"

	"example.com/net-http-class/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlogRepository interface {
	Create(blog *models.Blog) error
	FindBlogPost(userid uuid.UUID) ([]models.Blog, error)
	FindBlogPostByID(ID, userID uuid.UUID) (*models.Blog, error)
}

type blogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *blogRepository {
	return &blogRepository{db: db}
}

func (n *blogRepository) Create(blog *models.Blog) error {
	result := n.db.Create(blog)
	if result.Error != nil {
		return fmt.Errorf("user creation failed: %w", result.Error)
	}
	return nil
}

func (n *blogRepository) FindBlogPost(userid uuid.UUID) ([]models.Blog, error) {
	var Post []models.Blog

	result := n.db.Where("user_id = ?", userid).Find(&Post)
	if result.Error != nil {
		return nil, fmt.Errorf("could not find posts: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		// No posts found
		return nil, fmt.Errorf("no posts found for user %s", userid)
	}

	return Post, nil
}

func (n *blogRepository) FindBlogPostByID(ID, userID uuid.UUID) (*models.Blog, error) {
	var blog models.Blog
	err := n.db.Where("user_id = ? AND id = ?", userID, ID).First(&blog).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no Blogpost found")
		}
		return nil, fmt.Errorf("error fetching blogpost: %w", err)
	}

	return &blog, nil
}
