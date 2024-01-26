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

// GetUserRole godoc
//
// @Description get user role by ID
// @Tags user_roles
// @Produce json
// @Param id path integer true "User Role ID"
// @Success 200 {object} models.UserRole
// @Router /users/roles/{id} [get]
// @Security Authorization Token
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

// GetUserRoles godoc
//
// @Description get all user roles
// @Tags user_roles
// @Produce json
// @Success 200 {array} models.UserRole
// @Router /users/roles [get]
// @Security Authorization Token
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

// CreateUserRole godoc
//
// @Description create a new user role
// @Tags user_roles
// @Accept json
// @Produce json
// @Param input body UserRoleInput true "User Role Input"
// @Success 201 {object} models.UserRole
// @Router /users/roles [post]
// @Security Authorization Token
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

// UpdateUserRole godoc
//
// @Description update an existing user role by ID
// @Tags user_roles
// @Accept json
// @Produce json
// @Param id path integer true "User Role ID"
// @Param input body UserRoleInput true "User Role Input"
// @Success 200 {object} models.UserRole
// @Router /users/roles/{id} [put]
// @Security Authorization Token
func UpdateUserRole(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetUserRoleRepository()

	var input UserRoleInput
	_ = c.BindJSON(&input)
	model, err := r.GetByID(id)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user role not found"))
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

// DeleteUserRole godoc
//
// @Description delete an existing user role by ID
// @Tags user_roles
// @Param id path integer true "User Role ID"
// @Success 204
// @Router /users/roles/{id} [delete]
// @Security Authorization Token
func DeleteUserRole(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetUserRoleRepository()

	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("user role not found"))
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
