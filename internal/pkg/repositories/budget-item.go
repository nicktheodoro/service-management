package repositories

import (
	models "service-management/internal/pkg/models/budgets"
	"strconv"
)

type BudgetItemRepository struct{}

var budgetItemRep *BudgetItemRepository

func GetBudgetItemRepository() *BudgetItemRepository {
	if budgetItemRep == nil {
		budgetItemRep = &BudgetItemRepository{}
	}
	return budgetItemRep
}

func (r *BudgetItemRepository) GetAll(budgetID string) (*[]models.BudgetItem, error) {
	var m []models.BudgetItem
	w := models.BudgetItem{}
	w.BudgetID, _ = strconv.ParseUint(budgetID, 10, 64)

	err := Find(&w, &m, []string{"Provider", "Service"}, "id asc")
	return &m, err
}

func (r *BudgetItemRepository) GetByID(budgetID, id string) (*models.BudgetItem, error) {
	var m models.BudgetItem
	w := models.BudgetItem{}
	w.ID, _ = strconv.ParseUint(id, 10, 64)
	w.BudgetID, _ = strconv.ParseUint(budgetID, 10, 64)
	_, err := First(&w, &m, []string{"Provider", "Service"})

	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *BudgetItemRepository) Create(model *models.BudgetItem) error {
	return Create(&model)
}

func (r *BudgetItemRepository) Update(model *models.BudgetItem) error {
	w := models.BudgetItem{}
	w.ID = model.ID
	return Updates(&w, &model)
}

func (r *BudgetItemRepository) Delete(model *models.BudgetItem) error {
	return Delete(model)
}
