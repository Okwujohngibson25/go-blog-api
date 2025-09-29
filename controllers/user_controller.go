package controllers

import (
	"net/http"

	"example.com/net-http-class/models"
	"example.com/net-http-class/services"
	"github.com/gin-gonic/gin"
)

func Createnewuser(ctx *gin.Context, service *services.Userservicedependencies) {
	var user models.Users
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not bind data from req", "error": err.Error()})
		return
	}

	err = service.Createuser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create user", "error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Account created"})
}

func Loginuser(ctx *gin.Context, service *services.Userservicedependencies) {
	var user models.Users
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not bind data from req", "error": err.Error()})
		return
	}

	token, err := service.Loginuser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not Login user", "error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logged in", "token": token})
}
