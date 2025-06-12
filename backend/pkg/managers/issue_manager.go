package managers

import (
	"github.com/dev-jiemu/itops-assignment/backend/pkg/models"
	"time"
)

type IssueList struct {
	Issues []models.Issue
}

var IssueManager *IssueList

func NewIssueManager() *IssueList {
	return &IssueList{
		Issues: make([]models.Issue, 0),
	}
}

func (v *IssueList) GetIssueList(status string) []models.Issue {
	if status == "" { // 값이 없으면 전체 리턴
		return v.Issues
	}

	// 필터를 적용하면 기존 배열보다 클 순 없으니까 capacity 적용
	filterIssue := make([]models.Issue, 0, len(v.Issues))

	for _, issue := range v.Issues {
		if issue.Status == status {
			filterIssue = append(filterIssue, issue)
		}
	}

	return filterIssue
}

// GetIssueCount : issue seqno
func (v *IssueList) GetIssueCount() int {
	return len(v.Issues)
}

func (v *IssueList) AddIssue(issue models.Issue) {
	v.Issues = append(v.Issues, issue)
}

func (v *IssueList) GetIssue(issueId uint) models.Issue {
	for _, issue := range v.Issues {
		if issue.ID == issueId {
			return issue
		}
	}

	return models.Issue{}
}

func (v *IssueList) UpdateIssue(title, status string, issueId uint, userInfo *models.User) models.Issue {
	for _, issue := range v.Issues {
		if issue.ID == issueId {
			now := time.Now()

			issue.Title = title
			issue.Status = status
			issue.UpdatedAt = now
			issue.User = userInfo

			return issue
		}
	}

	return models.Issue{}
}
