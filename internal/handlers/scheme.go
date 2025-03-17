package handlers

import (
	"FASManagementSystem/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSchemesHandler(c *gin.Context) {
	schemes, err := services.GetAllSchemes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schemes)
}

func GetAllElligibleSchemesHandler(c *gin.Context) {
	applicantID := c.Query("applicant")
	if applicantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "applicant ID is required"})
		return
	}
	applicant, err := services.GetApplicantById(applicantID)
	// FIXME cant seem to get the applicant
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	schemes, err := services.GetElligibleSchemes(applicant)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, schemes)
}
