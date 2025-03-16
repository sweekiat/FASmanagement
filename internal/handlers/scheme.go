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