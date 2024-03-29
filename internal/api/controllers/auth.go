package controllers

import (
	"errors"
	"log"
	"net/http"
	"service-management/internal/pkg/repositories"
	"service-management/pkg/crypto"
	http_err "service-management/pkg/http-err"
	"service-management/pkg/token"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login godoc
//
// @Description authenticate user and generate access token
// @Tags auth
// @Accept json
// @Produce json
// @Param input body LoginInput true "Login Input"
// @Success 200 {string} string "JWT Token"
// @Router /login [post]
func Login(c *gin.Context) {
	var loginInput LoginInput
	_ = c.BindJSON(&loginInput)
	s := repositories.GetUserRepository()

	user, err := s.GetByEmail(loginInput.Email)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
		return
	}

	if !crypto.ComparePasswords(user.Hash, []byte(loginInput.Password)) {
		http_err.NewError(c, http.StatusForbidden, errors.New("user and password not match"))
		return
	}

	token, _ := token.CreateToken(user.Email)
	c.JSON(http.StatusOK, token)
}
