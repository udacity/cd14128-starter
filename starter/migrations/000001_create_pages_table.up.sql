-- This migration creates the pages table

CREATE TABLE pages (
    -- id is the primary key for the table  
    id SERIAL PRIMARY KEY,
    -- title is the title of the page
    title VARCHAR(255) NOT NULL,
    -- content is the content of the page
    content TEXT NOT NULL,
    -- created_at is the timestamp when the page was created
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    -- updated_at is the timestamp when the page was last updated
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);