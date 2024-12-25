package main

import (
	"log"
	"net/http"
	"time"

	"PersonalBlogBackend/internal/api/handlers"
	"PersonalBlogBackend/internal/api/middleware"
	"PersonalBlogBackend/internal/config"
	"PersonalBlogBackend/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 设置运行模式
	gin.SetMode(gin.DebugMode)

	log.Println("正在初始化数据库...")
	// 初始化数据库连接
	db := config.InitDB()

	log.Println("正在初始化服务...")
	// 初始化服务
	userService := service.NewUserService(db)
	articleService := service.NewArticleService(db)
	commentService := service.NewCommentService(db)

	log.Println("正在初始化路由...")
	// 初始化处理器
	userHandler := handlers.NewUserHandler(userService)
	articleHandler := handlers.NewArticleHandler(articleService)
	commentHandler := handlers.NewCommentHandler(commentService)

	// 初始化路由
	r := gin.Default()

	// 配置 CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Vue 开发服务器默认端口
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

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
		// 用户相关路由
		api.DELETE("/users/me", userHandler.DeleteUser) // 删除当前登录用户

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

	log.Printf("服务器启动成功，监听端口: 8080")
	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
