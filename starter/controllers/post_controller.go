// controllers/page_controller.go
package controllers

import (
	"cms-backend/models"
	"cms-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetPages retrieves all pages
func GetPages(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var pages []models.Page

	// TODO: Retrieve all pages from the database and assign to 'pages' variable

	// TODO: Handle any errors in fetching pages from the database
	c.JSON(http.StatusOK, pages)
}

// GetPage retrieves a specific page by ID
func GetPage(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var page models.Page

	// TODO: Retrieve a page by its ID from the database and assign to 'page' variable

	// TODO: Handle case where page with given ID is not found
	c.JSON(http.StatusOK, page)
}

// CreatePage creates a new page
func CreatePage(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var page models.Page

	// TODO: Bind JSON data from the request to the 'page' variable

	// TODO: Validate the 'page' data (e.g., check for required fields)

	// TODO: Save the new page to the database and handle any errors
	c.JSON(http.StatusCreated, page)
}

// UpdatePage updates an existing page by ID
func UpdatePage(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var page models.Page

	// TODO: Find the page by ID from the database

	// TODO: If the page is not found, return a 404 response

	// TODO: Bind updated JSON data from the request to the 'page' variable

	// TODO: Save the updated page to the database and handle any errors
	c.JSON(http.StatusOK, page)
}

// DeletePage deletes a page by ID
func DeletePage(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var page models.Page

	// TODO: Find the page by ID and delete it from the database

	// TODO: If the page is not found, return a 404 response

	// TODO: Return a success message upon deletion
	utils.SendError(c, http.StatusOK, "Page deleted")
}

