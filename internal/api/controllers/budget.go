package controllers

import (
	"errors"
	"log"
	"net/http"
	models "service-management/internal/pkg/models/budgets"
	"service-management/internal/pkg/repositories"

	http_err "service-management/pkg/http-err"

	"github.com/gin-gonic/gin"
)

type BudgetInput struct {
	Description string              `json:"description" binding:"required"`
	Note        string              `json:"note"`
	Discount    int64               `json:"discount"`
	Surcharge   int64               `json:"surcharge"`
	Status      models.BudgetStatus `json:"status" binding:"required"`
	CustomerID  uint64              `json:"customer_id" binding:"required"`
	Items       []models.BudgetItem `json:"items"`
}

type BudgetUpdateStatusInput struct {
	Status models.BudgetStatus `json:"status" binding:"required"`
}

// GetBudget godoc
// @Description get all budgets
// @Tags budgets
// @Produce json
// @Success 200 {array} budgets.Budget
// @Router /budgets [get]
// @Security Authorization Token
func GetBudget(c *gin.Context) {
	r := repositories.GetBudgetRepository()
	id := c.Param("id")

	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("budget not found"))
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}

// GetBudget godoc
//
//	@Description get Budget by ID
//	@Tags budgets
//	@Produce json
//	@Param id path integer true "Budget ID"
//	@Success 200 {object} budgets.Budget
//	@Router /budgets/{id} [get]
//	@Security Authorization Token
func GetBudgets(c *gin.Context) {
	r := repositories.GetBudgetRepository()
	models, err := r.GetAll()
	if err != nil {
		log.Println(err)
		http_err.NewError(c, http.StatusNotFound, errors.New("budgets not found"))
		return
	}

	c.JSON(http.StatusOK, models)
}

// CreateBudget godoc
// @Description create a new Budget
// @Tags budgets
// @Accept json
// @Produce json
// @Param input body BudgetInput true "Budget Input"
// @Success 201 {object} budgets.Budget
// @Router /budgets [post]
// @Security Authorization Token
func CreateBudget(c *gin.Context) {
	r := repositories.GetBudgetRepository()
	var input BudgetInput
	_ = c.BindJSON(&input)
	modelToAdd := models.Budget{
		Description: input.Description,
		Discount:    input.Discount,
		Surcharge:   input.Surcharge,
		Note:        input.Note,
		Status:      input.Status,
		CustomerID:  input.CustomerID,
		Items:       input.Items,
	}

	err := r.Create(&modelToAdd)
	if err != nil {
		http_err.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusCreated, modelToAdd)
}

// UpdateBudget godoc
// @Description update an existing budget by ID
// @Tags budgets
// @Accept json
// @Produce json
// @Param id path integer true "Budget ID"
// @Param input body BudgetInput true "Budget Input"
// @Success 200 {object} budgets.Budget
// @Router /budgets/{id} [put]
// @Security Authorization Token
func UpdateBudget(c *gin.Context) {
	id := c.Params.ByName("id")
	var input BudgetInput
	_ = c.BindJSON(&input)

	r := repositories.GetBudgetRepository()
	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("budget not found"))
		log.Println(err)
		return
	}

	model.Description = input.Description
	model.Note = input.Note
	model.Status = input.Status
	model.CustomerID = input.CustomerID
	model.Items = input.Items
	err = r.Update(model)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}

// DeleteBudget godoc
// @Description delete an existing Budget by ID
// @Tags budgets
// @Param id path integer true "Budget ID"
// @Success 204
// @Router /budgets/{id} [delete]
// @Security Authorization Token
func DeleteBudget(c *gin.Context) {
	id := c.Params.ByName("id")
	r := repositories.GetBudgetRepository()

	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("budget not found"))
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

// UpdateBudget godoc
// @Description update an existing budget status
// @Tags budgets
// @Accept json
// @Produce json
// @Param id path integer true "Budget ID"
// @Param input body BudgetUpdateStatusInput true "Budget Update Status Input"
// @Success 200 {object} budgets.Budget
// @Router /budgets/{id}/status [patch]
// @Security Authorization Token
func UpdateBudgetStatus(c *gin.Context) {
	id := c.Params.ByName("id")
	var input BudgetUpdateStatusInput
	_ = c.BindJSON(&input)

	r := repositories.GetBudgetRepository()
	model, err := r.GetByID(id)
	if err != nil {
		http_err.NewError(c, http.StatusNotFound, errors.New("budget not found"))
		log.Println(err)
		return
	}

	model.Status = input.Status
	err = r.Update(model)

	if err != nil {
		http_err.NewError(c, http.StatusNotFound, err)
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, model)
}
