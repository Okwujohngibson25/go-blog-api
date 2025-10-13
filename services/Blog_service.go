package services

import (
	"fmt"

	"example.com/net-http-class/models"
	"example.com/net-http-class/repository"
	"github.com/google/uuid"
)

type BlogService struct {
	blogRepo repository.BlogRepository
}

func NewBlogService(blogRepo repository.BlogRepository) *BlogService {
	return &BlogService{blogRepo: blogRepo}
}

func (t *BlogService) CreateBlogPost(userid uuid.UUID, post *models.Blogrequest) error {
	User_id := userid

	blog := models.Blog{
		Title:  post.Title,
		Post:   post.Post,
		UserID: User_id,
	}

	err := t.blogRepo.Create(&blog)
	if err != nil {
		return fmt.Errorf("failed to create blog in service: %w", err)
	}
	return nil
}

func (t *BlogService) FetchBlogPost(userid uuid.UUID) ([]models.Blog, error) {
	Posts, err := t.blogRepo.FindBlogPost(userid)
	if err != nil {
		return nil, fmt.Errorf("could not find posts: %w", err)
	}

	return Posts, nil
}

func (t *BlogService) FetchBlogPostById(ID, userID uuid.UUID) (*models.Blog, error) {
	post, err := t.blogRepo.FindBlogPostByID(ID, userID)
	if err != nil {
		return nil, fmt.Errorf("could not find post: %w", err)
	}
	return post, nil
}
