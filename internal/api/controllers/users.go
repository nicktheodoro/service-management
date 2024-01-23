package controllers

import (
	"errors"
	"log"
	"net/http"
	models "service-management/internal/pkg/models/users"
	"service-management/internal/pkg/repositories"
	"service-management/pkg/crypto"

	http_err "github.com/antonioalfa22/go-rest-template/pkg/http-err"
	"github.com/gin-gonic/gin"
)

type UserInput struct {
	Firstname string `json:"first_name" binding:"required"`
	Lastname  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	RoleID    uint64 `json:"role_id" binding:"required"`
}

func GetUser(c *gin.Context) {
	r := repositories.GetUserRepository()
	id := c.Param("id")

	user, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	r := repositories.GetUserRepository()
	resp, err := r.Get()
	if err != nil {
		log.Println(err)
		http_err.NewError(c, http.StatusNotFound, errors.New("users not found"))
		return
	}

	c.JSON(http.StatusOK, resp)
}

func CreateUser(c *gin.Context) {
	r := repositories.GetUserRepository()
	var userInput UserInput
	_ = c.BindJSON(&userInput)
	user := models.User{
		Email:     userInput.Email,
		Firstname: userInput.Firstname,
		Lastname:  userInput.Lastname,
		Hash:      crypto.HashAndSalt([]byte(userInput.Password)),
		RoleID:    userInput.RoleID,
	}
	err := r.Create(&user)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetUserRepository()

	var userInput UserInput
	_ = c.BindJSON(&userInput)
	user, err := r.GetByID(id)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
		return
	}

	user.Email = userInput.Email
	user.Lastname = userInput.Lastname
	user.Firstname = userInput.Firstname
	user.Hash = crypto.HashAndSalt([]byte(userInput.Password))
	user.RoleID = userInput.RoleID

	err = r.Update(user)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetUserRepository()

	user, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
		return
	}

	err = r.Delete(user)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "")
}
