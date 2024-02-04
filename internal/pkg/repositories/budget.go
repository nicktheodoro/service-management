package repositories

import (
	models "service-management/internal/pkg/models/budgets"
	"strconv"
)

type BudgetRepository struct{}

var budgetRepository *BudgetRepository

func GetBudgetRepository() *BudgetRepository {
	if budgetRepository == nil {
		budgetRepository = &BudgetRepository{}
	}
	return budgetRepository
}

func (r *BudgetRepository) GetAll() (*[]models.Budget, error) {
	var m []models.Budget
	err := Find(&models.Budget{}, &m, []string{}, "id asc")
	return &m, err
}

func (r *BudgetRepository) GetByID(id string) (*models.Budget, error) {
	var m models.Budget
	where := models.Budget{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &m, []string{})

	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *BudgetRepository) Create(model *models.Budget) error {
	return Create(&model)
}

func (r *BudgetRepository) Update(model *models.Budget) error {
	where := models.Budget{}
	where.ID = model.ID
	return Updates(where, model)
}

func (r *BudgetRepository) Delete(model *models.Budget) error {
	return Delete(model)
}
