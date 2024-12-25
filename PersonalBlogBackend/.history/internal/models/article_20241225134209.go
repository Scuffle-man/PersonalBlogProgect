package models

import (
	"time"
)

// Article 文章模型
type Article struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Content     string    `json:"content" gorm:"type:text;not null"`
	AuthorID    uint      `json:"author_id" gorm:"not null"`
	Author      User      `json:"author" gorm:"foreignKey:AuthorID"`
	PublishedAt time.Time `json:"published_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Comments    []Comment `json:"comments,omitempty" gorm:"foreignKey:ArticleID"`
} 