package handlers

import (
	"FASManagementSystem/internal/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetAllApplicationsHandler(c *gin.Context) {
	applications, err := services.GetAllApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, applications)
}
