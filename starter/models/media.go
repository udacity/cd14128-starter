package models

import "time"

// This struct includes fields for:
// - ID (unsigned integer, primary key)
// - URL (string, required, with max length)
// - Type (string, for storing media type)
// - CreatedAt (timestamp for creation date)
// - UpdatedAt (timestamp for last update)

type Media struct {
	//ID field as uint with gorm tag for primary key and json tag
    ID        uint      `gorm:"primaryKey" json:"id"`

	//URL field as string with gorm tag for size limit (255) and not null constraint and json tag and binding tag to make it required
    URL       string    `gorm:"size:255;not null" json:"url" binding:"required"`
	
	//Type field as string with gorm tag for size limit (50) and json tag and binding tag to make it required
    Type      string    `gorm:"size:50" json:"type" binding:"required"`
	
	//CreatedAt field as time.Time with gorm tag for automatic timestamp on creation and json tag
    CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	
	//UpdatedAt field as time.Time with gorm tag for automatic timestamp on updates and json tag
    UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}