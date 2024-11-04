package models

// TODO: Import the required "time" package for timestamp fields

// TODO: Create a Page struct that will represent pages in our CMS
// This struct should include fields for:
// - ID (unsigned integer, primary key)
// - Title (string, required, with max length)
// - Content (text field, required)
// - CreatedAt (timestamp for creation date)
// - UpdatedAt (timestamp for last update)

type Page struct {
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
    
    // TODO: Add CreatedAt field using time.Time with:
    // - gorm tag for automatic timestamp on creation
    // - json tag for serialization
    
    // TODO: Add UpdatedAt field using time.Time with:
    // - gorm tag for automatic timestamp on updates
    // - json tag for serialization
}
