package test

import (
	"fmt"
	models "service-management/internal/pkg/models/budgets"
	"service-management/internal/pkg/models/customers"
	"service-management/internal/pkg/models/providers"
	"service-management/internal/pkg/models/services"
	"service-management/internal/pkg/repositories"
	"testing"
)

var BudgetTest models.Budget

func SetupBudgetTests() error {
	customer := customers.Customer{
		Name:     "Test",
		Phone:    "+5521990170011",
		Document: "123.456.789-00",
	}
	customerRep := repositories.GetCustomerRepository()
	if err := customerRep.Create(&customer); err != nil {
		return err
	}

	provider := providers.Provider{
		Name:     "Antonio Paya",
		Email:    "antonio@teste.com",
		Phone:    "+5521990178696",
		Document: "123.456.789-00",
	}
	providerRep := repositories.GetProviderRepository()
	if err := providerRep.Create(&provider); err != nil {
		return err
	}

	service := services.Service{
		Description: "Interior painting",
		Value:       1000,
	}
	serviceRep := repositories.GetServiceRepository()
	if err := serviceRep.Create(&service); err != nil {
		return err
	}

	CustomerTest = customer
	ProviderTest = provider
	ServiceTest = service

	return nil
}

func TestAddBudget(t *testing.T) {
	SetupTests()
	SetupBudgetTests()

	modelToAdd := models.Budget{
		Description: "Internal painting",
		Discount:    0,
		Surcharge:   0,
		Note:        "",
		Status:      "REQUESTED",
		CustomerID:  CustomerTest.ID,
	}
	r := repositories.GetBudgetRepository()
	if err := r.Create(&modelToAdd); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(BudgetTest.Items) > 0 {
		t.Fatalf("Expected budget items lenght 0")
	}

	BudgetTest = modelToAdd
}

func TestUpdateBudget(t *testing.T) {
	r := repositories.GetBudgetRepository()

	modelToUpdate, err := r.GetByID(fmt.Sprint(BudgetTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	modelToUpdate.Description = "description_updated"
	modelToUpdate.Discount = 99
	modelToUpdate.Surcharge = 99
	modelToUpdate.Note = "note_updated"
	modelToUpdate.Status = "PROGESS"

	if err := r.Update(modelToUpdate); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expect := BudgetTest.Description != modelToUpdate.Description ||
		BudgetTest.Discount != modelToUpdate.Discount ||
		BudgetTest.Surcharge != modelToUpdate.Surcharge ||
		BudgetTest.Note != modelToUpdate.Note ||
		BudgetTest.Status != modelToUpdate.Status

	if !expect {
		t.Fatalf("UpdateBudget test failed")
	}

	for i, address := range BudgetTest.Items {
		expect = address != modelToUpdate.Items[i]
	}

	if !expect {
		t.Fatalf("UpdateCustomer adresses test failed")
	}
}

func TestGetAllBudgets(t *testing.T) {
	r := repositories.GetBudgetRepository()
	if _, err := r.GetAll(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetBudgetById(t *testing.T) {
	r := repositories.GetBudgetRepository()
	if _, err := r.GetByID(fmt.Sprint(BudgetTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDeleteBudget(t *testing.T) {
	r := repositories.GetBudgetRepository()

	modelToDelete, err := r.GetByID(fmt.Sprint(BudgetTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := r.Delete(modelToDelete); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
