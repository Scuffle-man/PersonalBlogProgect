package main

import (
	"log"

	"PersonalBlogBackend/internal/api/handlers"
	"PersonalBlogBackend/internal/api/middleware"
	"PersonalBlogBackend/internal/config"
	"PersonalBlogBackend/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 设置运行模式
	gin.SetMode(gin.ReleaseMode)

	// 初始化数据库连接
	db := config.InitDB()

	// 初始化服务
	userService := service.NewUserService(db)
	articleService := service.NewArticleService(db)

	// 初始化处理器
	userHandler := handlers.NewUserHandler(userService)
	articleHandler := handlers.NewArticleHandler(articleService)

	// 初始化路由
	r := gin.Default()

	// 配置中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 公开路由
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)

	// 需要认证的路由
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 文章相关路由
		api.GET("/articles", articleHandler.GetArticles)
		api.POST("/articles", articleHandler.CreateArticle)
		api.GET("/articles/:id", articleHandler.GetArticle)
		api.PUT("/articles/:id", articleHandler.UpdateArticle)
		api.DELETE("/articles/:id", articleHandler.DeleteArticle)

		// 评论相关路由
		api.POST("/articles/:id/comments", commentHandler.CreateComment)
		api.GET("/articles/:id/comments", commentHandler.GetComments)
	}

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
