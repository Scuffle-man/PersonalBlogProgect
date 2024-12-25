package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"PersonalBlogBackend/internal/models"
)

// Database 配置结构体
type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// InitDB 初始化数据库连接
func InitDB() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Printf("未找到 .env 文件: %v", err)
	}

	dbConfig := Database{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "3306"),
		User:     getEnv("DB_USER", "root"),
		Password: getEnv("DB_PASSWORD", ""),
		DBName:   getEnv("DB_NAME", "blog"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("数据库连接失败: %v", err)
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(&models.User{}, &models.Article{}, &models.Comment{})
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	return db
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
} 