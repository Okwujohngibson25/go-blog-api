package controllers

import (
	"net/http"

	"example.com/net-http-class/models"
	"example.com/net-http-class/services"
	"github.com/gin-gonic/gin"
)

func CreateBlogpost(ctx *gin.Context, Blogservice *services.Blogservice) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(401, gin.H{"error": "missing token"})
		return
	}

	var input models.Blogrequest

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request", "error": err})
		return
	}

	err = Blogservice.CreateBlogPost(authHeader, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create blog post", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "blog post created"})
}

func FetchBlogPost(ctx *gin.Context, Blogservice *services.Blogservice) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(401, gin.H{"error": "missing token"})
		return
	}

	post, err := Blogservice.FetchBlogPost(token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not Fetch blog post for user", "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "post found", "post": post})
}
