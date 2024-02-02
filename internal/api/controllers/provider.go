package controllers

import (
	"errors"
	"log"
	"net/http"
	models "service-management/internal/pkg/models/providers"
	"service-management/internal/pkg/repositories"

	http_err "service-management/pkg/http-err"

	"github.com/gin-gonic/gin"
)

type ProviderInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Document string `json:"document" binding:"required"`
	Note     string `json:"note" form:"note"`
}

// GetProvider godoc
// @Description get all providers
// @Tags providers
// @Produce json
// @Success 200 {array} providers.Provider
// @Router /providers [get]
// @Security Authorization Token
func GetProvider(c *gin.Context) {
	r := repositories.GetProviderRepository()
	id := c.Param("id")

	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("provider not found"))
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}

// GetProvider godoc
//
//	@Description get Provider by ID
//	@Tags providers
//	@Produce json
//	@Param id path integer true "Provider ID"
//	@Success 200 {object} providers.Provider
//	@Router /providers/{id} [get]
//	@Security Authorization Token
func GetProviders(c *gin.Context) {
	r := repositories.GetProviderRepository()
	models, err := r.GetAll()
	if err != nil {
		log.Println(err)
		http_err.NewError(c, http.StatusNotFound, errors.New("providers not found"))
		return
	}

	c.JSON(http.StatusOK, models)
}

// CreateProvider godoc
// @Description create a new Provider
// @Tags providers
// @Accept json
// @Produce json
// @Param input body ProviderInput true "Provider Input"
// @Success 201 {object} providers.Provider
// @Router /providers [post]
// @Security Authorization Token
func CreateProvider(c *gin.Context) {
	r := repositories.GetProviderRepository()
	var input ProviderInput
	_ = c.BindJSON(&input)
	model := models.Provider{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Document: input.Document,
		Note:     input.Note,
	}
	err := r.Create(&model)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, model)
}

// UpdateProvider godoc
// @Description update an existing provider by ID
// @Tags providers
// @Accept json
// @Produce json
// @Param id path integer true "Provider ID"
// @Param input body ProviderInput true "Provider Input"
// @Success 200 {object} providers.Provider
// @Router /providers/{id} [put]
// @Security Authorization Token
func UpdateProvider(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetProviderRepository()

	var input ProviderInput
	_ = c.BindJSON(&input)
	model, err := r.GetByID(id)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("provider not found"))
		log.Println(err)
		return
	}

	model.Name = input.Name
	model.Email = input.Email
	model.Document = input.Document
	model.Note = input.Note
	err = r.Update(model)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}

// DeleteProvider godoc
// @Description delete an existing Provider by ID
// @Tags providers
// @Param id path integer true "Provider ID"
// @Success 204
// @Router /providers/{id} [delete]
// @Security Authorization Token
func DeleteProvider(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetProviderRepository()

	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("provider not found"))
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
