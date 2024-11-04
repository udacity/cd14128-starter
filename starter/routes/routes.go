package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TODO: Import required packages:
// - controllers package for route handlers
// - gin-gonic/gin for web framework
// - gorm.io/gorm for database type

// InitializeRoutes sets up all API routes
func InitializeRoutes(router *gin.Engine, db *gorm.DB) {
    // TODO: Add database middleware that:
    // 1. Creates a middleware function that accepts a gin.Context parameter
    // 2. Uses router.Use() to register the middleware
    // 3. Inside the middleware:
    //    - Stores the db instance in the context using a "db" key
    //    - Calls Next() to continue to the next handler
    
    // TODO: Create API version group that:
    // 1. Uses router.Group() to create a new route group
    // 2. Sets the group prefix to "/api/v1"
    // 3. Stores the group in a variable for adding routes
    // Hint: This will be used to prefix all API routes
    
    // TODO: Within the api group, add the following route groups:
    {
        // TODO: Page Routes
        // GET    /api/v1/pages      - List all pages
        // GET    /api/v1/pages/:id  - Get single page
        // POST   /api/v1/pages      - Create new page
        // PUT    /api/v1/pages/:id  - Update existing page
        // DELETE /api/v1/pages/:id  - Delete page
        // Example:
        // api.GET("/pages", controllers.GetPages)
        
        // TODO: Post Routes
        // GET    /api/v1/posts      - List all posts
        // GET    /api/v1/posts/:id  - Get single post
        // POST   /api/v1/posts      - Create new post
        // PUT    /api/v1/posts/:id  - Update existing post
        // DELETE /api/v1/posts/:id  - Delete post
        // Example:
        // api.GET("/posts", controllers.GetPosts)
        
        // TODO: Media Routes
        // GET    /api/v1/media      - List all media
        // GET    /api/v1/media/:id  - Get single media
        // POST   /api/v1/media      - Upload new media
        // DELETE /api/v1/media/:id  - Delete media
        // Example:
        // api.GET("/media", controllers.GetMedia)
    }
}
