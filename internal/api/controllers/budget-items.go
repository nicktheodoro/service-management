package controllers

import (
	"errors"
	"log"
	"net/http"
	models "service-management/internal/pkg/models/budgets"
	"service-management/internal/pkg/repositories"
	"strconv"

	http_err "service-management/pkg/http-err"

	"github.com/gin-gonic/gin"
)

type BudgetItemInput struct {
	Note       string `json:"note" `
	ProviderID uint64 `json:"provider_id" binding:"required"`
	ServiceID  uint64 `json:"service_id" binding:"required"`
}

// GetBudgetItems godoc
//
// @Description get all budget items
// @Tags budgets_items
// @Produce json
// @Success 200 {array} budgets.BudgetItem
// @Router /budgets/{id}/items [get]
// @Security Authorization Token
func GetBudgetItems(c *gin.Context) {
	budgetID := c.Param("id")
	r := repositories.GetBudgetItemRepository()

	res, err := r.GetAll(budgetID)
	if err != nil {
		log.Println(err)
		http_err.NewError(c, http.StatusNotFound, errors.New("budget item not found"))
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetBudgetItem godoc
//
//	@Description get Budget Item by ID
//	@Tags budgets_items
//	@Produce json
//	@Param id path integer true "Budget Item ID"
//	@Success 200 {object} budgets.BudgetItem
//	@Router /budgets/{id}/items/{itemId} [get]
//	@Security Authorization Token
func GetBudgetItem(c *gin.Context) {
	budgetID := c.Params.ByName("id")
	itemID := c.Params.ByName("itemId")
	r := repositories.GetBudgetItemRepository()

	res, err := r.GetByID(budgetID, itemID)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("budget item not found"))
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// CreateBudgetItem godoc
//
// @Description create a new budget item
// @Tags budgets_items
// @Accept json
// @Produce json
// @Param input body BudgetItemInput true "Budget Item Input"
// @Success 201 {object} budgets.BudgetItem
// @Router /budgets/{id}/items [post]
// @Security Authorization Token
func CreateBudgetItem(c *gin.Context) {
	budgetID, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	r := repositories.GetBudgetItemRepository()

	var input BudgetItemInput
	_ = c.BindJSON(&input)
	modelToAdd := models.BudgetItem{
		BudgetID:   budgetID,
		ProviderID: input.ProviderID,
		ServiceID:  input.ServiceID,
	}
	err := r.Create(&modelToAdd)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, modelToAdd)
}

// UpdateBudgetItem godoc
//
// @Description update an existing budget item by ID
// @Tags budgets_items
// @Accept json
// @Produce json
// @Param id path integer true "Budget ID"
// @Param input body BudgetItemInput true "Budget Item Input"
// @Success 200 {object} budgets.BudgetItem
// @Router /budgets/{id}/items/{itemId} [put]
// @Security Authorization Token
func UpdateBudgetItem(c *gin.Context) {
	budgetID := c.Params.ByName("id")
	itemID := c.Params.ByName("itemId")
	r := repositories.GetBudgetItemRepository()

	var input BudgetItemInput
	_ = c.BindJSON(&input)
	modelToUpdate, err := r.GetByID(budgetID, itemID)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("budget item not found"))
		log.Println(err)
		return
	}

	modelToUpdate.ProviderID = input.ProviderID
	modelToUpdate.ServiceID = input.ServiceID

	err = r.Update(modelToUpdate)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, modelToUpdate)
}

// DeleteBudgetItem godoc
//
// @Description delete an existing budget item by ID
// @Tags budgets_items
// @Param id path integer true "Budget Item ID"
// @Success 204
// @Router /budgets/{id}/items/{itemId} [delete]
// @Security Authorization Token
func DeleteBudgetItem(c *gin.Context) {
	budgetID := c.Params.ByName("id")
	itemID := c.Params.ByName("itemId")
	r := repositories.GetBudgetItemRepository()

	modelToDelete, err := r.GetByID(budgetID, itemID)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("budget item not found"))
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
