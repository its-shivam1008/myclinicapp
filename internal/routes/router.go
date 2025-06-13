package routes

import (
    "myclinic/internal/handler"
    "myclinic/internal/middleware"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    patientHandler := handler.NewPatientHandler(db)

    r.POST("/login", handler.LoginHandler)

    auth := r.Group("/api")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.POST("/patients", patientHandler.CreatePatient)
        auth.GET("/patients", patientHandler.GetAllPatients)
        auth.PUT("/patients/:id", patientHandler.UpdatePatient)
        auth.DELETE("/patients/:id", patientHandler.DeletePatient)
    }
}
