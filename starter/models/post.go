package models

// TODO: Import the required "time" package for timestamp fields

// TODO: Create a Post struct that will represent blog posts in our CMS
// This struct should include fields for:
// - ID (unsigned integer, primary key)
// - Title (string, required, with max length)
// - Content (text field, required)
// - Author (string, optional)
// - CreatedAt (timestamp for creation date)
// - UpdatedAt (timestamp for last update)
// - Media (slice of Media, representing a many-to-many relationship)

type Post struct {
    // TODO: Add ID field as uint with:
    // - gorm tag for primary key
    // - json tag for serialization
    
    // TODO: Add Title field as string with:
    // - gorm tags for size limit (255) and not null constraint
    // - json tag for serialization
    // - binding tag to make it required
    
    // TODO: Add Content field as string with:
    // - gorm tag specifying text type and not null constraint
    // - json tag for serialization
    // - binding tag to make it required
    
    // TODO: Add Author field as string with:
    // - gorm tag for size limit (100)
    // - json tag for serialization
    
    // TODO: Add CreatedAt field using time.Time with:
    // - gorm tag for automatic timestamp on creation
    // - json tag for serialization
    
    // TODO: Add UpdatedAt field using time.Time with:
    // - gorm tag for automatic timestamp on updates
    // - json tag for serialization
    
    // TODO: Add Media field as []Media with:
    // - gorm tag for many-to-many relationship (specify junction table name: post_media)
    // - json tag for serialization
}