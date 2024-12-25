package handlers

import (
	"log"
	"net/http"

	"PersonalBlogBackend/internal/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// Register 用户注册
func (h *UserHandler) Register(c *gin.Context) {
	log.Printf("收到注册请求")
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("注册请求参数错误: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("尝试创建用户: %s", input.Username)
	user, err := h.userService.CreateUser(input.Username, input.Email, input.Password)
	if err != nil {
		log.Printf("创建用户失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("用户创建成功: %s", user.Username)
	c.JSON(http.StatusCreated, user)
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.ValidateUser(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 生成 JWT token
	token, err := h.userService.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "生成令牌失败"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}

// DeleteUser 删除用户
func (h *UserHandler) DeleteUser(c *gin.Context) {
	// 获取当前登录用户的ID
	userID := c.GetUint("userID")
	
	// 调用服务删除用户
	if err := h.userService.DeleteUser(userID); err != nil {
		log.Printf("删除用户失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "用户已成功删除",
	})
}
