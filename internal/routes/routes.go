package routes

import (
	"FASManagementSystem/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/api/applicants", handlers.GetApplicantsHandler)
	router.POST("/api/applicants", handlers.CreateNewApplicantHandler)
	router.DELETE("/api/applicants/:id", handlers.DeleteApplicantHandler)
	router.GET("/api/schemes", handlers.GetAllSchemesHandler)
	router.GET("/api/applications", handlers.GetAllApplicationsHandler)
	router.GET("/api/schemes/elligible", handlers.GetAllElligibleSchemesHandler)
	router.POST("/api/applications", handlers.CreateNewApplicationHandler)

}
