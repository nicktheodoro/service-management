package test

import (
	"fmt"
	models "service-management/internal/pkg/models/budgets"
	"service-management/internal/pkg/repositories"
	"testing"
	"time"
)

var BudgetItemTest models.BudgetItem

func TestAddBudgetItem(t *testing.T) {
	SetupTests()
	SetupBudgetTests()

	budget := models.Budget{
		Description: "Internal painting",
		Discount:    0,
		Surcharge:   0,
		Note:        "",
		Status:      "REQUESTED",
		CustomerID:  CustomerTest.ID,
	}
	budgetRep := repositories.GetBudgetRepository()
	if err := budgetRep.Create(&budget); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	BudgetTest = budget

	modelToAdd := models.BudgetItem{
		BudgetID:   BudgetTest.ID,
		ProviderID: ProviderTest.ID,
		ServiceID:  ServiceTest.ID,
	}
	r := repositories.GetBudgetItemRepository()
	if err := r.Create(&modelToAdd); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(BudgetTest.Items) > 0 {
		t.Fatalf("Expected budget items lenght 0")
	}

	BudgetItemTest = modelToAdd
}

func TestUpdateBudgetItem(t *testing.T) {
	r := repositories.GetBudgetItemRepository()

	modelToUpdate, err := r.GetByID(fmt.Sprint(BudgetTest.ID), fmt.Sprint(BudgetItemTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	modelToUpdate.UpdatedAt = time.Now()
	modelToUpdate.CreatedAt = time.Now()

	if err := r.Update(modelToUpdate); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expect := BudgetItemTest.UpdatedAt != modelToUpdate.UpdatedAt ||
		BudgetItemTest.CreatedAt != modelToUpdate.CreatedAt

	if !expect {
		t.Fatalf("UpdateBudget test failed")
	}
}

func TestGetAllBudgetItems(t *testing.T) {
	r := repositories.GetBudgetItemRepository()
	if _, err := r.GetAll(fmt.Sprint(BudgetTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetBudgetItemById(t *testing.T) {
	r := repositories.GetBudgetItemRepository()
	if _, err := r.GetByID(fmt.Sprint(BudgetTest.ID), fmt.Sprint(BudgetItemTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDeleteBudgetItem(t *testing.T) {
	r := repositories.GetBudgetItemRepository()

	modelToDelete, err := r.GetByID(fmt.Sprint(BudgetTest.ID), fmt.Sprint(BudgetItemTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := r.Delete(modelToDelete); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
