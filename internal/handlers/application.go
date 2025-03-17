package handlers

import (
	"FASManagementSystem/internal/models"
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
func CreateNewApplicationHandler(c *gin.Context) {
	var newApplication models.Application
	err := c.BindJSON(&newApplication)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.CreateNewApplication(newApplication)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newApplication)
}
