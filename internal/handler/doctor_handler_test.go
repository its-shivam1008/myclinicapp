package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupDoctorRouter() *gin.Engine {
	router := gin.Default()
	mock := NewDoctorHandler(nil) 
	router.GET("/doctor/patients", mock.GetAllPatientsForDoctor)
	router.PUT("/doctor/patients/:id", mock.UpdatePatientForDoctor)
	return router
}

func TestGetAllPatientsForDoctor(t *testing.T) {
	router := setupDoctorRouter()
	req, _ := http.NewRequest("GET", "/doctor/patients", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestUpdatePatientForDoctor(t *testing.T) {
	router := setupDoctorRouter()
	req, _ := http.NewRequest("PUT", "/doctor/patients/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
}
