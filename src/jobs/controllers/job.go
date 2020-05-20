package controllers

import (
	"github/pankaj070888/test-micorservices/src/jobs/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var jobModel = new(models.Job)

type JobController struct{}

func getJobs(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"code":    http.StatusOK,
			"message": "Welcome server 01",
		},
	)
}
