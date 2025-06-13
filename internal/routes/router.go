package routes

import (
    "myclinic/internal/handler"
    "myclinic/internal/middleware"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    patientHandler := handler.NewPatientHandler(db)
    doctorHandler := handler.NewDoctorHandler(db)

    r.POST("/login", handler.LoginHandler)

    auth := r.Group("/api")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.POST("/patients", patientHandler.CreatePatient)
        auth.GET("/patients", patientHandler.GetAllPatients)
        auth.PUT("/patients/:id", patientHandler.UpdatePatient)
        auth.DELETE("/patients/:id", patientHandler.DeletePatient)

        auth.GET("/doctor/patients", doctorHandler.GetAllPatients)
		auth.GET("/doctor/patients/:id", doctorHandler.GetPatientByID)
		auth.PUT("/doctor/patients/:id", doctorHandler.UpdatePatient)
    }
}
