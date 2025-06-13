package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDB struct{}

func (m *MockDB) Where(query interface{}, args ...interface{}) *gorm.DB {
	return &gorm.DB{} 
}

func TestLoginHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	DB := &gorm.DB{} 

	router.POST("/login", LoginHandler(DB))

	loginBody := LoginRequest{
		Username: "testuser",
		Password: "password123",
		Role:     "receptionist",
	}

	body, _ := json.Marshal(loginBody)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
