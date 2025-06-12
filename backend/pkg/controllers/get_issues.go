package controllers

import (
	"github.com/dev-jiemu/itops-assignment/backend/pkg/managers"
	"github.com/dev-jiemu/itops-assignment/backend/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetIssueList(c *gin.Context) {
	var status = c.Query("status")

	// status value check
	validStatus := isValidStatus(status)
	if !validStatus {
		c.JSON(http.StatusBadRequest, gin.H{"error": "status field not found"})
		return
	}

	issueList := managers.IssueManager.GetIssueList(status)
	c.JSON(http.StatusOK, issueList)
}

func GetIssueOne(c *gin.Context) {
	var issueSeq = c.Param("id")
	id, err := strconv.Atoi(issueSeq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if id == 0 { // 0으로 오는것도 비정상 처리
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	issueInfo := managers.IssueManager.GetIssue(uint(id))
	if issueInfo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		return
	}

	c.JSON(http.StatusOK, issueInfo)
}

func isValidStatus(status string) bool {
	if status == models.ISSUE_STATUS_PENDING || status == models.ISSUE_STATUS_IN_PROGRESS || status == models.ISSUE_STATUS_COMPLETED || status == models.ISSUE_STATUS_CANCELLED {
		return true
	}

	return false
}
