package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"PersonalBlogBackend/internal/service"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

// CreateComment 创建评论
func (h *CommentHandler) CreateComment(c *gin.Context) {
	articleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	var input struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("绑定评论JSON失败: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.GetUint("userID")
	log.Printf("用户 %d 正在为文章 %d 创建评论", userID, articleID)

	comment, err := h.commentService.CreateComment(input.Content, uint(articleID), userID)
	if err != nil {
		log.Printf("创建评论失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("评论创建成功: %v", comment.ID)
	c.JSON(http.StatusCreated, comment)
}

// GetComments 获取文章评论
func (h *CommentHandler) GetComments(c *gin.Context) {
	articleID, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的文章ID"})
		return
	}

	log.Printf("获取文章 %d 的评论列表", articleID)
	comments, err := h.commentService.GetCommentsByArticleID(uint(articleID))
	if err != nil {
		log.Printf("获取评论失败: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
} 