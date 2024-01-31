package controllers

import (
	"errors"
	"log"
	"net/http"
	models "service-management/internal/pkg/models/services"
	"service-management/internal/pkg/repositories"

	http_err "service-management/pkg/http-err"

	"github.com/gin-gonic/gin"
)

type ServiceInput struct {
	Description string               `json:"description" binding:"required"`
	Value       int64                `json:"value" binding:"required"`
	Note        string               `json:"note"`
	Status      models.ServiceStatus `json:"status"`
}

// GetServices godoc
//
// @Description get all services
// @Tags services
// @Produce json
// @Success 200 {array} services.Service
// @Router /services [get]
// @Security Authorization Token
func GetServices(c *gin.Context) {
	r := repositories.GetServiceRepository()
	res, err := r.GetAll()
	if err != nil {
		log.Println(err)
		http_err.NewError(c, http.StatusNotFound, errors.New("services not found"))
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetService godoc
//
//	@Description get Service by ID
//	@Tags services
//	@Produce json
//	@Param id path integer true "Service ID"
//	@Success 200 {object} services.Service
//	@Router /services/{id} [get]
//	@Security Authorization Token
func GetService(c *gin.Context) {
	r := repositories.GetServiceRepository()
	id := c.Param("id")

	res, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("service not found"))
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateService godoc
//
// @Description create a new service
// @Tags services
// @Accept json
// @Produce json
// @Param input body ServiceInput true "Service Input"
// @Success 201 {object} services.Service
// @Router /services [post]
// @Security Authorization Token
func CreateService(c *gin.Context) {
	r := repositories.GetServiceRepository()
	var input ServiceInput
	_ = c.BindJSON(&input)
	modelToAdd := models.Service{
		Description: input.Description,
		Value:       input.Value,
		Note:        input.Note,
		Status:      models.Requested,
	}
	err := r.Create(&modelToAdd)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, modelToAdd)
}

// UpdateService godoc
//
// @Description update an existing service by ID
// @Tags services
// @Accept json
// @Produce json
// @Param id path integer true "Service ID"
// @Param input body ServiceInput true "ServiceInput"
// @Success 200 {object} services.Service
// @Router /services/{id} [put]
// @Security Authorization Token
func UpdateService(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetServiceRepository()

	var input ServiceInput
	_ = c.BindJSON(&input)
	modelToUpdate, err := r.GetByID(id)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("service not found"))
		log.Println(err)
		return
	}

	modelToUpdate.Description = input.Description
	modelToUpdate.Value = input.Value
	modelToUpdate.Note = input.Note
	modelToUpdate.Status = input.Status

	err = r.Update(modelToUpdate)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, modelToUpdate)
}

// DeleteService godoc
//
// @Description delete an existing service by ID
// @Tags services
// @Param id path integer true "Service ID"
// @Success 204
// @Router /services/{id} [delete]
// @Security Authorization Token
func DeleteService(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetServiceRepository()

	modelToDelete, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("services not found"))
		log.Println(err)
		return
	}

	err = r.Delete(modelToDelete)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusNoContent, "")
}
