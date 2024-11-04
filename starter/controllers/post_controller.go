package controllers

import (
	"cms-backend/models"
	"cms-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPosts retrieves all posts with optional filtering
func GetPosts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var posts []models.Post

	title := c.Query("title")
	author := c.Query("author")

	query := db
	if title != "" {
		query = query.Where("title ILIKE ?", "%"+title+"%")
	}
	if author != "" {
		query = query.Where("author = ?", author)
	}

	// Use proper preloading for media relationships
	if err := query.Preload("Media").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, posts)
}

// GetPost retrieves a specific post by ID
func GetPost(c *gin.Context) {
    // TODO: Get database connection
    // Get database instance from Gin context
    
    // TODO: Handle ID parameter
    // 1. Get the ID from URL parameter
      // Example: /api/posts/123
    
    // 2. Convert string ID to uint


    // TODO: Retrieve post from database
    // 1. Define post variable and query database

    // 2. Return the post
}

// CreatePost creates a new post
func CreatePost(c *gin.Context) {
    // TODO: Get database connection
    // Get database instance from Gin context
    
    // TODO: Parse and validate input
    // 1. Define post variable to store incoming data
    
    // 2. Parse JSON request body into post struct

    // TODO: Validate required fields
    // Validate required fields

    // TODO: Create post in database
    // 1. Start database transaction
    
    // 2. Create the post

    // 3. Commit transaction

    // 4. Return created post
}

// UpdatePost updates an existing post
func UpdatePost(c *gin.Context) {
    // TODO: Get database connection
    // Get database instance from Gin context
    
    // TODO: Parse and validate ID
    // Get ID from URL parameter
      // Example: /api/posts/123

    // TODO: Check if post exists
    // Find existing post

    // TODO: Parse and validate update data
    // Define variable for update input

    // TODO: Update post fields
    // Update only the fields that are allowed to be updated

    // TODO: Save updates to database
    // 1. Start transaction
    
    // 2. Save the updated post

    // 3. Commit transaction

    // 4. Return updated post
}

// DeletePost deletes a post
func DeletePost(c *gin.Context) {
    // TODO: Get database connection
    // Get database instance from Gin context
    
    // TODO: Parse and validate ID
    // Get ID from URL parameter
	 // Example: /api/posts/123

    // TODO: Check if post exists
    // Find existing post

    // TODO: Delete post from database
    // 1. Start transaction
    
    // 2. Delete the post
    // Note: Consider if you want soft delete (recommended) or hard delete

    // 3. Commit transaction

    // 4. Return success message
}