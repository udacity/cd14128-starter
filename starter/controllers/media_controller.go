package controllers

import (
    "cms-backend/models"
    "cms-backend/utils"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func GetMedia(c *gin.Context) {
    // TODO: Implement logic to retrieve all media
}

func GetMediaByID(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
    // TODO: Implement logic to retrieve media by ID
}

func CreateMedia(c *gin.Context) {
    // TODO: Implement logic to create new media
}

func DeleteMedia(c *gin.Context) {
    // TODO: Implement logic to delete media
}

