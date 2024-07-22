package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	models "github.com/hahahamid/go-crud/Models"
	"github.com/hahahamid/go-crud/initializers"
)

func CreatePosts(c *gin.Context) {

	var body struct {
		Title    string
		Body     string
		Category string
		Author   string
	}

	c.Bind((&body))

	post := models.BlogPost{Title: body.Title, Body: body.Body, Category: body.Category, Author: body.Author}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		log.Fatal("Error")
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func GetAllPosts(c *gin.Context) {

	var posts []models.BlogPost
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.Status(400)
		log.Fatal("Error")
		return
	}

	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {

	id := c.Param("id")
	var post models.BlogPost
	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(400)
		log.Fatal(result.Error)
		return
	}

	c.JSON(200, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {

	id := c.Param("id")
	var body struct {
		Title string
		Body  string
	}

	c.Bind((&body))

	var post models.BlogPost
	result := initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.BlogPost{Title: body.Title, Body: body.Body})

	if result.Error != nil {
		c.Status(400)
		log.Fatal("Error")
		return
	}

	c.JSON(200, gin.H{
		"updated post": post,
	})
}

func DeletePost(c *gin.Context) {

	id := c.Param("id")
	result := initializers.DB.Delete(&models.BlogPost{}, id)

	if result.Error != nil {
		c.Status(400)
		log.Fatal("Error")
		return
	}

	c.JSON(200, gin.H{
		"message": "Post deleted successfully",
	})
}
