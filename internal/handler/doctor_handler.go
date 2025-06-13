package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"myclinic/internal/models"
	"gorm.io/gorm"
)

type DoctorHandler struct {
	DB *gorm.DB
}

func NewDoctorHandler(db *gorm.DB) *DoctorHandler {
	return &DoctorHandler{DB: db}
}

func (h *DoctorHandler) GetAllPatients(c *gin.Context) {
	var patients []models.Patient
	if err := h.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func (h *DoctorHandler) GetPatientByID(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := h.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func (h *DoctorHandler) UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var existing models.Patient

	if err := h.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}

	var input models.Patient
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	existing.Name = input.Name
	existing.Age = input.Age
	existing.Address = input.Address
	existing.Prescription = input.Prescription

	if err := h.DB.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update patient"})
		return
	}

	c.JSON(http.StatusOK, existing)
}
