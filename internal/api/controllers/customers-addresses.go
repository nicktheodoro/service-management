package controllers

import (
	"errors"
	"log"
	"net/http"
	models "service-management/internal/pkg/models/customers"
	"service-management/internal/pkg/repositories"
	"strconv"

	http_err "service-management/pkg/http-err"

	"github.com/gin-gonic/gin"
)

type CustomerAddressInput struct {
	ZipCode    string `json:"zip_code" binding:"required"`
	Street     string `json:"street" binding:"required"`
	Number     string `json:"number" binding:"required"`
	Complement string `json:"complement"`
	District   string `json:"district" binding:"required"`
	City       string `json:"city" binding:"required"`
	State      string `json:"state" binding:"required"`
}

// GetCustomerAddresses godoc
//
// @Description get all customer addresses
// @Tags customers_addresses
// @Produce json
// @Success 200 {array} customers.CustomerAddress
// @Router /customers/{id}/addresses [get]
// @Security Authorization Token
func GetCustomerAddresses(c *gin.Context) {
	customerID := c.Param("id")
	r := repositories.GetCustomerAddressRepository()

	res, err := r.GetAll(customerID)
	if err != nil {
		log.Println(err)
		http_err.NewError(c, http.StatusNotFound, errors.New("customer addresses not found"))
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetCustomerAddress godoc
//
//	@Description get Customer Address by ID
//	@Tags customers_addresses
//	@Produce json
//	@Param id path integer true "Customer Address ID"
//	@Success 200 {object} customers.CustomerAddress
//	@Router /customers/{id}/addresses/{addressId} [get]
//	@Security Authorization Token
func GetCustomerAddress(c *gin.Context) {
	customerID := c.Params.ByName("id")
	addressID := c.Params.ByName("addressId")
	r := repositories.GetCustomerAddressRepository()

	res, err := r.GetByID(customerID, addressID)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("costumer address not found"))
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateCustomerAddress godoc
//
// @Description create a new customer address
// @Tags customers_addresses
// @Accept json
// @Produce json
// @Param input body CustomerAddressInput true "Customer Address Input"
// @Success 201 {object} customers.CustomerAddress
// @Router /customers/{id}/addresses [post]
// @Security Authorization Token
func CreateCustomerAddress(c *gin.Context) {
	customerID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	r := repositories.GetCustomerAddressRepository()

	var input CustomerAddressInput
	_ = c.BindJSON(&input)
	modelToAdd := models.CustomerAddress{
		ZipCode:    input.ZipCode,
		Street:     input.Street,
		Number:     input.Number,
		Complement: input.Complement,
		District:   input.District,
		City:       input.City,
		State:      input.State,
		CustomerID: customerID,
	}
	err := r.Create(&modelToAdd)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, modelToAdd)
}

// UpdateCustomerAddress godoc
//
// @Description update an existing customer address by ID
// @Tags customers_addresses
// @Accept json
// @Produce json
// @Param id path integer true "Customer ID"
// @Param input body CustomerAddressInput true "Customer Address Input"
// @Success 200 {object} customers.CustomerAddress
// @Router /customers/{id}/addresses/{addressId} [put]
// @Security Authorization Token
func UpdateCustomerAddress(c *gin.Context) {
	customerID := c.Params.ByName("id")
	addressID := c.Params.ByName("addressId")
	r := repositories.GetCustomerAddressRepository()

	var input CustomerAddressInput
	_ = c.BindJSON(&input)
	modelToUpdate, err := r.GetByID(customerID, addressID)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("costumer address not found"))
		log.Println(err)
		return
	}

	modelToUpdate.ZipCode = input.ZipCode
	modelToUpdate.Street = input.Street
	modelToUpdate.Number = input.Number
	modelToUpdate.Complement = input.Complement
	modelToUpdate.District = input.District
	modelToUpdate.City = input.City
	modelToUpdate.State = input.State

	err = r.Update(modelToUpdate)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, modelToUpdate)
}

// DeleteCustomerAddress godoc
//
// @Description delete an existing customer address by ID
// @Tags customers_addresses
// @Param id path integer true "Customer Address ID"
// @Success 204
// @Router /customers/{id}/addresses/{addressId} [delete]
// @Security Authorization Token
func DeleteCustomerAddress(c *gin.Context) {
	customerID := c.Params.ByName("id")
	addressID := c.Params.ByName("addressId")
	r := repositories.GetCustomerAddressRepository()

	modelToDelete, err := r.GetByID(customerID, addressID)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("costumer address not found"))
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
