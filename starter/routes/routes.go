package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
    //TODO: Add middleware to inject db into context
    // Middleware to inject db into context

    //TODO: Add routes for pages, posts, and media
    {
        // Page routes

        // Post routes

        // Media routes

    }
}
