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

// @title MyClinic API
// @version 1.0
// @description API for managing patients and user authentication.
// @host localhost:8080
// @BasePath /
// CreatePatient godoc
// @Summary Create a new patient
// @Tags Patients
// @Accept json
// @Produce json
// @Param request body models.Patient true "Patient Info"
// @Success 201 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /patients [post]
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

// GetAllPatients godoc
// @Summary Get all patients
// @Tags Patients
// @Produce json
// @Success 200 {array} models.Patient
// @Failure 500 {object} map[string]string
// @Router /patients [get]
func (h *PatientHandler) GetAllPatients(c *gin.Context) {
	var patients []models.Patient
	if err := h.DB.Find(&patients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch"})
		return
	}
	c.JSON(http.StatusOK, patients)
}

// UpdatePatient godoc
// @Summary Update a patient
// @Tags Patients
// @Accept json
// @Produce json
// @Param id path int true "Patient ID"
// @Param patient body models.Patient true "Patient object"
// @Success 200 {object} models.Patient
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /patients/{id} [put]
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

// DeletePatient godoc
// @Summary Delete a patient
// @Tags Patients
// @Produce json
// @Param id path int true "Patient ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /patients/{id} [delete]
func (h *PatientHandler) DeletePatient(c *gin.Context) {
	id := c.Param("id")
	if err := h.DB.Delete(&models.Patient{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
