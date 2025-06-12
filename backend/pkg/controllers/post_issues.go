package controllers

import (
	"github.com/dev-jiemu/itops-assignment/backend/pkg/managers"
	"github.com/dev-jiemu/itops-assignment/backend/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type CreateIssueReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserId      uint   `json:"userId"`
}

func CreateIssue(c *gin.Context) {
	var req CreateIssueReq

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInfo := managers.UserManager.GetUserInfo(req.UserId)
	if userInfo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "user is not found"})
		return
	}

	now := time.Now()

	// 생성은 PENDING 으로
	newIssue := models.Issue{
		ID:          uint(managers.IssueManager.GetIssueCount() + 1),
		Title:       req.Title,
		Description: req.Description,
		Status:      models.ISSUE_STATUS_PENDING,
		User:        userInfo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	managers.IssueManager.AddIssue(newIssue)
	c.JSON(http.StatusCreated, newIssue)
}
