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

	// 首先连接MySQL（不指定数据库）
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", 
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 创建数据库
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;", dbConfig.DBName)
	if err := db.Exec(createDBSQL).Error; err != nil {
		log.Fatalf("创建数据库失败: %v", err)
	}

	// 连接到指定的数据库
	dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.DBName,
	)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	// 自动迁移数据库表
	err = db.AutoMigrate(
		&models.User{},
		&models.Article{},
		&models.Comment{},
	)
	if err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	log.Println("数据库初始化成功")
	return db
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
} 