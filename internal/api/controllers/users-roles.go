package controllers

import (
	"errors"
	"log"
	"net/http"
	models "service-management/internal/pkg/models/users"
	"service-management/internal/pkg/repositories"
	http_err "service-management/pkg/http-err"

	"github.com/gin-gonic/gin"
)

type UserRoleInput struct {
	RoleName string `json:"role_name" binding:"required"`
}

func GetUserRole(c *gin.Context) {
	r := repositories.GetUserRoleRepository()
	id := c.Param("id")

	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user role not found"))
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}

func GetUserRoles(c *gin.Context) {
	r := repositories.GetUserRoleRepository()
	models, err := r.GetAll()
	if err != nil {
		log.Println(err)
		http_err.NewError(c, http.StatusNotFound, errors.New("user roles not found"))
		return
	}

	c.JSON(http.StatusOK, models)
}

func CreateUserRole(c *gin.Context) {
	r := repositories.GetUserRoleRepository()
	var input UserRoleInput
	_ = c.BindJSON(&input)
	model := models.UserRole{
		RoleName: input.RoleName,
	}
	err := r.Create(&model)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, model)
}

func UpdateUserRole(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetUserRoleRepository()

	var input UserRoleInput
	_ = c.BindJSON(&input)
	model, err := r.GetByID(id)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
		log.Println(err)
		return
	}

	model.RoleName = input.RoleName

	err = r.Update(model)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}

func DeleteUserRole(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetUserRoleRepository()

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
