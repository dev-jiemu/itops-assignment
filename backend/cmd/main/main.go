package main

import (
	"github.com/dev-jiemu/itops-assignment/backend/pkg/controllers"
	"github.com/dev-jiemu/itops-assignment/backend/pkg/managers"
	"github.com/gin-gonic/gin"
)

func main() {
	// init manager
	managers.UserManager = managers.NewUserManager()
	managers.IssueManager = managers.NewIssueManager()

	// setting API Server
	router := gin.Default()

	router.GET("/issue", controllers.GetIssueList)
	router.GET("/issue/:id", controllers.GetIssueOne)
	router.POST("/issue", controllers.CreateIssue)
	router.PATCH("/issue/:id", controllers.PatchIssue)

	router.Run(":8080")
}
