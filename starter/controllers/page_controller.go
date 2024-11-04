package controllers

import (
	"cms-backend/models"
	"cms-backend/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TODO: Import required packages:
// - models package for Page struct
// - utils package for error handling
// - net/http for status codes
// - strconv for string to uint conversion
// - gin-gonic/gin for web framework
// - gorm.io/gorm for database operations

// GetPages retrieves all pages
func GetPages(c *gin.Context) {
    // Get database instance from context
    db := c.MustGet("db").(*gorm.DB)
	// Declare pages variable
    var pages []models.Page

    // TODO: Query all pages from database:
    // if err := db.Find(&pages).Error; err != nil {
    //     Handle error case here
    // }

    // TODO: Return success response:
}

// GetPage retrieves a specific page by ID
func GetPage(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "Invalid ID format",
		})
		return
	}

	var page models.Page
	if err := db.First(&page, uint(id)).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.HTTPError{
			Code:    http.StatusNotFound,
			Message: "Page not found",
		})
		return
	}
	c.JSON(http.StatusOK, page)
}

// CreatePage creates a new page
func CreatePage(c *gin.Context) {
	// TODO: Get database instance from context

	// TODO: Declare page variable

    // TODO: Bind JSON request body:

    // TODO: Start transaction:
    
    // TODO: Create page in database:
    
    // TODO: Commit transaction and return response:
}

// UpdatePage updates an existing page by ID
func UpdatePage(c *gin.Context) {
    // TODO: Get database instance from context
    
    // TODO: Convert string ID to uint:
    
    // TODO: Find existing page:

    // TODO: Bind JSON update data:

    // TODO: Update page fields:

    // TODO: Start transaction and save:

    // TODO: Return success response:
}

// DeletePage deletes a page by ID
func DeletePage(c *gin.Context) {
	// TODO: Get database instance from context
    
    // TODO: Convert string ID to uint:

    // TODO: Check if page exists:

    // TODO: Start transaction and delete:

    // TODO: Return success response:
}
