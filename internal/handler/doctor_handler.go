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

// @title MyClinic API
// @version 1.0
// @description API for managing patients and user authentication.
// @host localhost:8080
// @BasePath /
// GetAllPatientsForDoctor godoc
// @Summary Get all patients (Doctor access)
// @Tags Doctor
// @Produce json
// @Success 200 {array} models.Patient
// @Failure 500 {object} map[string]string
// @Router /doctor/patients [get]
func (h *DoctorHandler) GetAllPatientsForDoctor(c *gin.Context) {
	var patients []models.Patient
	if err := h.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch patients"})
		return
	}
	c.JSON(http.StatusOK, patients)
}

// GetPatientByID godoc
// @Summary Get patient by ID
// @Tags Patients
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} models.Patient
// @Failure 404 {object} map[string]string
// @Router /patients/{id} [get]
func (h *DoctorHandler) GetPatientByIDForDoctor(c *gin.Context) {
	id := c.Param("id")
	var patient models.Patient
	if err := h.DB.First(&patient, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Patient not found"})
		return
	}
	c.JSON(http.StatusOK, patient)
}

// UpdatePatientByDoctor godoc
// @Summary Doctor updates patient info
// @Tags Doctor
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param patient body models.Patient true "Updated info"
// @Success 200 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Router /doctor/patients/{id} [put]
func (h *DoctorHandler) UpdatePatientForDoctor(c *gin.Context) {
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
