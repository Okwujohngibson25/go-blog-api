package main

import (
	"log"

	"example.com/net-http-class/config"
	"example.com/net-http-class/controllers"
	"example.com/net-http-class/models"
	"example.com/net-http-class/services"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default() // gin initialization

	db := config.DBconnect()                               // db *gorm.db stored in db
	err := db.AutoMigrate(&models.Users{}, &models.Blog{}) // Automigration of tables
	if err != nil {
		log.Fatal("could not create migration", err)
	}
	// Services
	Userservice := services.NewUserservicedependencies(db)
	Blogservice := services.NewBlogservice(db)

	// routes.Routers(server, Userservice, Blogservice)

	server.POST("/create", func(ctx *gin.Context) {
		controllers.Createnewuser(ctx, Userservice)
	})

	server.POST("/login", func(ctx *gin.Context) {
		controllers.Loginuser(ctx, Userservice)
	})

	server.POST("/createpost", func(ctx *gin.Context) {
		controllers.CreateBlogpost(ctx, Blogservice)
	})

	server.GET("/fetchpost", func(ctx *gin.Context) {
		controllers.FetchBlogPost(ctx, Blogservice)
	})

	server.Run(":8080") // Runing on localhost:8080
}
