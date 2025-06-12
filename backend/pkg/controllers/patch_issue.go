package controllers

import (
	"github.com/dev-jiemu/itops-assignment/backend/pkg/managers"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UpdateIssueReq struct {
	Title  string `json:"title"`
	Status string `json:"status"`
	UserId uint   `json:"user_id"`
}

func PatchIssue(c *gin.Context) {
	var req UpdateIssueReq

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

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// issue info check
	issueInfo := managers.IssueManager.GetIssue(uint(id))
	if issueInfo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue not found"})
		return
	}

	// user info check
	userInfo := managers.UserManager.GetUserInfo(req.UserId)
	if userInfo == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	updateIssue := managers.IssueManager.UpdateIssue(req.Title, req.Status, issueInfo.ID, userInfo)
	if updateIssue.ID == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Issue update failed"})
		return
	}

	c.JSON(http.StatusOK, updateIssue)
}
