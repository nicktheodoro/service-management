package controllers

import (
	"errors"
	"log"
	"net/http"
	"service-management/internal/pkg/repositories"

	"github.com/antonioalfa22/go-rest-template/pkg/crypto"
	http_err "github.com/antonioalfa22/go-rest-template/pkg/http-err"
	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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

	token, _ := crypto.CreateToken(user.Email)
	c.JSON(http.StatusOK, token)
}
