package routes

import (
	"github.com/gin-gonic/gin"

	"FASManagementSystem/internal/handlers"
)


func SetupRoutes(router *gin.Engine) {
        router.GET("/api/applicants", handlers.GetApplicantsHandler)
		router.POST("/api/applicants", handlers.GetApplicantsHandler)
		router.GET("/api/schemes", handlers.GetApplicantsHandler)
		
}