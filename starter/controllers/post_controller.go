package controllers

import (
	"cms-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPosts(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var posts []models.Post
    // TODO: Implement logic to retrieve all posts with optional filtering
}

func GetPost(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    // TODO: Implement logic to retrieve a single post by ID
}

func CreatePost(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    var post models.Post
    // TODO: Implement logic to create a new post
    // Remember to handle the association with media
}

func UpdatePost(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    // TODO: Implement logic to update an existing post
}

func DeletePost(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    // TODO: Implement logic to delete a post
}

