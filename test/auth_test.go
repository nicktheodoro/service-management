package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"service-management/internal/api/controllers"
	models "service-management/internal/pkg/models/users"
	"service-management/internal/pkg/repositories"
	"service-management/pkg/crypto"
	"testing"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func TestLogin(t *testing.T) {
	SetupTests()
	r := repositories.GetUserRepository()

	testUser := models.User{
		Firstname: "Test",
		Lastname:  "User",
		Email:     "testuser@test.com",
		Hash:      crypto.HashAndSalt([]byte("testpassword")), // Hash the password
		Role:      models.UserRole{RoleName: "user"},
	}

	if err := r.Create(&testUser); err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	loginInput := LoginInput{
		Email:    "testuser@test.com",
		Password: "testpassword",
	}

	w := performLoginRequest(loginInput)
	if w.Code != http.StatusOK {
		t.Fatalf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	var token string
	err := json.Unmarshal(w.Body.Bytes(), &token)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if token == "" {
		t.Fatalf("Expected a valid JWT token in the response")
	}
}

// Helper function to perform a login request
func performLoginRequest(loginInput LoginInput) *httptest.ResponseRecorder {
	router := gin.Default()

	router.POST("/login", controllers.Login)

	jsonInput, _ := json.Marshal(loginInput)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonInput))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}
