package models

import (
	"time"
)

// User 用户模型
type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username" gorm:"unique;not null"`
	Email        string    `json:"email" gorm:"unique;not null"`
	PasswordHash string    `json:"-" gorm:"not null"` // json:"-" 确保不会在JSON响应中返回密码
	CreatedAt    time.Time `json:"created_at"`
}
