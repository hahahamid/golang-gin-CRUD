package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hahahamid/go-crud/controllers"
	"github.com/hahahamid/go-crud/initializers"
)

func init() {

	initializers.LoadEnvironmentVaraible()
	initializers.ConnectDB()
}

func main() {
	r := gin.Default()

	r.POST("/add-posts", controllers.CreatePosts)
	r.GET("/get-posts", controllers.GetAllPosts)
	r.GET("/get-post/:id", controllers.GetPost)
	r.PUT("/update-post/:id", controllers.UpdatePost)
	r.DELETE("/delete-post/:id", controllers.DeletePost)

	r.Run() // listen and serve on 0.0.0.0:8080
}
