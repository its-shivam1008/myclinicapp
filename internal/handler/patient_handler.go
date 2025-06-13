package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"myclinic/internal/models"
	"gorm.io/gorm"
)

type PatientHandler struct {
	DB *gorm.DB
}

func NewPatientHandler(db *gorm.DB) *PatientHandler {
	return &PatientHandler{DB: db}
}

func (h *PatientHandler) CreatePatient(c *gin.Context) {
	var patient models.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.DB.Create(&patient).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create"})
		return
	}
	c.JSON(http.StatusCreated, patient)
}

func (h *PatientHandler) GetAllPatients(c *gin.Context) {
	var patients []models.Patient
	if err := h.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch"})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func (h *PatientHandler) UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := h.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	h.DB.Save(&patient)
	c.JSON(http.StatusOK, patient)
}

func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Patient{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
