package controllers

import (
	"net/http"

	"example.com/net-http-class/models"
	"example.com/net-http-class/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateBlogpost(ctx *gin.Context, Blogservice *services.Blogservice) {
	userID := ctx.MustGet("userid").(uuid.UUID) // getting Userid passed inside gin.context header from middleware

	var input models.Blogrequest

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request", "error": err})
		return
	}

	err = Blogservice.CreateBlogPost(userID, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create blog post", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "blog post created"})
}

func FetchBlogPost(ctx *gin.Context, Blogservice *services.Blogservice) {
	userID := ctx.MustGet("userid").(uuid.UUID) // getting Userid passed inside gin.context header from middleware

	post, err := Blogservice.FetchBlogPost(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not Fetch blog post for user", "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "post found", "post": post})
}
