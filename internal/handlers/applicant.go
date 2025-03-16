package handlers

import (
	"FASManagementSystem/internal/models"
	"FASManagementSystem/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetApplicantsHandler(c *gin.Context) {
	applicants, err := services.GetAllApplicantsWithHousehold()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, applicants)
}
func CreateNewApplicantHandler(c *gin.Context) {
	var newApplicant models.Applicant
	contentErr := c.ShouldBindJSON(&newApplicant)
	if contentErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": contentErr.Error()})
	}
	err := services.CreateNewApplicant(newApplicant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newApplicant)
}


