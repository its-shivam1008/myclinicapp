package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupPatientRouter() *gin.Engine {
	router := gin.Default()
	mock := NewPatientHandler(nil) 
	router.GET("/patients", mock.GetAllPatients)
	return router
}

func TestGetPatients(t *testing.T) {
	router := setupPatientRouter()
	req, _ := http.NewRequest("GET", "/patients", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
