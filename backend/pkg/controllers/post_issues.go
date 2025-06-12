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

	// 담당자가 없는 경우 : PENDING
	// 있는 경우 : IN_PROGRESS
	userInfo := managers.UserManager.GetUserInfo(req.UserId)
	//if userInfo.ID == 0 {
	//	c.JSON(http.StatusNotFound, gin.H{"error": "user is not found"})
	//	return
	//}

	now := time.Now()

	newIssue := models.Issue{
		ID:          uint(managers.IssueManager.GetIssueCount() + 1),
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if userInfo.ID == 0 {
		newIssue.Status = models.ISSUE_STATUS_PENDING
	} else {
		newIssue.Status = models.ISSUE_STATUS_IN_PROGRESS
		newIssue.User = userInfo
	}

	managers.IssueManager.AddIssue(newIssue)
	c.JSON(http.StatusCreated, newIssue)
}
