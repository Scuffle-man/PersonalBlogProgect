package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"PersonalBlogBackend/internal/service"
)

type ArticleHandler struct {
	articleService *service.ArticleService
}

func NewArticleHandler(articleService *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

// GetArticles 获取文章列表
func (h *ArticleHandler) GetArticles(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	articles, err := h.articleService.GetArticles(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, articles)
}

// CreateArticle 创建文章
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	log.Printf("开始创建文章")
	userID, exists := c.Get("userID")
	if !exists {
		log.Printf("未找到用户ID")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未找到用户信息"})
		return
	}
	log.Printf("用户ID: %v", userID)

	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("绑定JSON失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	article, err := h.articleService.CreateArticle(input.Title, input.Content, userID.(uint))
	if err != nil {
		log.Printf("创建文章失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("文章创建成功: %v", article.ID)
	c.JSON(http.StatusCreated, article)
}

// GetArticle 获取单篇文章
func (h *ArticleHandler) GetArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	article, err := h.articleService.GetArticleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "文章不存在"})
		return
	}

	c.JSON(http.StatusOK, article)
}

// UpdateArticle 更新文章
func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	var input struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")
	article, err := h.articleService.UpdateArticle(uint(id), input.Title, input.Content, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, article)
}

// DeleteArticle 删除文章
func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	userID := c.GetUint("userID")
	if err := h.articleService.DeleteArticle(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "文章已删除"})
} 