package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "FASManagementSystem/internal/services" 
)


func GetApplicantsHandler(c *gin.Context) {
    applicants, err := services.GetAllApplicantsWithHousehold()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, applicants)
}