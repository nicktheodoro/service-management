package controllers

import (
	"errors"
	"log"
	"net/http"
	models "service-management/internal/pkg/models/customers"
	"service-management/internal/pkg/repositories"

	http_err "service-management/pkg/http-err"

	"github.com/gin-gonic/gin"
)

type CustomerInput struct {
	Name     string                   `json:"name" binding:"required"`
	Email    string                   `json:"email" binding:"required"`
	Phone    string                   `json:"phone" binding:"required"`
	Document string                   `json:"document" binding:"required"`
	Address  []models.CustomerAddress `json:"address"`
}

// GetCustomer godoc
//
//	@Description get Customer by ID
//	@Tags customers
//	@Produce json
//	@Param id path integer true "Customer ID"
//	@Success 200 {object} customers.Customer
//	@Router /customers/{id} [get]
//	@Security Authorization Token
func GetCustomer(c *gin.Context) {
	r := repositories.GetCustomerRepository()
	id := c.Param("id")

	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("customer not found"))
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}

// GetCustomer godoc
//
// @Description get all customers
// @Tags customers
// @Produce json
// @Success 200 {array} customers.Customer
// @Router /customers [get]
// @Security Authorization Token
func GetCustomers(c *gin.Context) {
	r := repositories.GetCustomerRepository()
	models, err := r.GetAll()
	if err != nil {
		log.Println(err)
		http_err.NewError(c, http.StatusNotFound, errors.New("customers not found"))
		return
	}

	c.JSON(http.StatusOK, models)
}

// CreateCustomer godoc
//
// @Description create a new Customer
// @Tags customers
// @Accept json
// @Produce json
// @Param input body CustomerInput true "Customer Input"
// @Success 201 {object} customers.Customer
// @Router /customers [post]
// @Security Authorization Token
func CreateCustomer(c *gin.Context) {
	r := repositories.GetCustomerRepository()
	var input CustomerInput
	_ = c.BindJSON(&input)
	model := models.Customer{
		Name:     input.Name,
		Email:    input.Email,
		Phone:    input.Phone,
		Document: input.Document,
		Address:  input.Address,
	}
	err := r.Create(&model)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, model)
}

// UpdateCustomer godoc
//
// @Description update an existing customer by ID
// @Tags customers
// @Accept json
// @Produce json
// @Param id path integer true "Customer ID"
// @Param input body CustomerInput true "Customer Input"
// @Success 200 {object} customers.Customer
// @Router /customers/{id} [put]
// @Security Authorization Token
func UpdateCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetCustomerRepository()

	var input CustomerInput
	_ = c.BindJSON(&input)
	model, err := r.GetByID(id)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("customer not found"))
		log.Println(err)
		return
	}

	model.Name = input.Name
	model.Email = input.Email
	model.Document = input.Document
	model.Address = input.Address
	err = r.Update(model)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}

// DeleteCustomer godoc
//
// @Description delete an existing Customer by ID
// @Tags customers
// @Param id path integer true "Customer ID"
// @Success 204
// @Router /customers/{id} [delete]
// @Security Authorization Token
func DeleteCustomer(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetCustomerRepository()

	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("customer not found"))
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
