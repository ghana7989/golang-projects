package controllers

import (
	"github.com/ghana7989-go-crud/initializers"
	"github.com/ghana7989-go-crud/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// GET data off request body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	// Create a new post
	post := models.Post{
		Title: body.Title,
		Body:  body.Body,
	}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	// return the post
	c.JSON(200, post)
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	initializers.DB.Find(&posts)
	c.JSON(200, posts)
}

func PostsShow(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))
	c.JSON(200, post)
}

func PostsUpdate(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)
	post.Title = body.Title
	post.Body = body.Body
	initializers.DB.Save(&post)
	c.JSON(200, post)
}

func PostsDelete(c *gin.Context) {
	var post models.Post
	initializers.DB.First(&post, c.Param("id"))
	initializers.DB.Delete(&post)
	c.JSON(200, gin.H{"message": "Post deleted"})
}
