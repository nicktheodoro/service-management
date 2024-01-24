package controllers

import (
	"errors"
	"log"
	"net/http"
	models "service-management/internal/pkg/models/users"
	"service-management/internal/pkg/repositories"
	"service-management/pkg/crypto"

	http_err "service-management/pkg/http-err"

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

	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}

func GetUsers(c *gin.Context) {
	r := repositories.GetUserRepository()
	models, err := r.GetAll()
	if err != nil {
		log.Println(err)
		http_err.NewError(c, http.StatusNotFound, errors.New("users not found"))
		return
	}

	c.JSON(http.StatusOK, models)
}

func CreateUser(c *gin.Context) {
	r := repositories.GetUserRepository()
	var input UserInput
	_ = c.BindJSON(&input)
	model := models.User{
		Email:     input.Email,
		Firstname: input.Firstname,
		Lastname:  input.Lastname,
		Hash:      crypto.HashAndSalt([]byte(input.Password)),
		RoleID:    input.RoleID,
	}
	err := r.Create(&model)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, model)
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetUserRepository()

	var input UserInput
	_ = c.BindJSON(&input)
	model, err := r.GetByID(id)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
		return
	}

	model.Email = input.Email
	model.Lastname = input.Lastname
	model.Firstname = input.Firstname
	model.Hash = crypto.HashAndSalt([]byte(input.Password))
	model.RoleID = input.RoleID

	err = r.Update(model)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetUserRepository()

	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
		return
	}

	err = r.Delete(model)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "")
}
