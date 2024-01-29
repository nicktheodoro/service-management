package test

import (
	"fmt"
	models "service-management/internal/pkg/models/customers"
	"service-management/internal/pkg/repositories"
	"testing"
)

var CustomerTest models.Customer

func TestAddCustomer(t *testing.T) {
	SetupTests()

	modelToAdd := models.Customer{
		Name:     "Antonio Paya",
		Email:    "antonio@teste.com",
		Phone:    "+5521990178696",
		Document: "123.456.789-00",
		Address: models.CustomerAddress{
			ZipCode:  "25955540",
			Street:   "Fortaleza",
			Number:   "139",
			District: "Bom Retiro",
			City:     "Teresópolis",
			State:    "RJ",
		},
	}
	r := repositories.GetCustomerRepository()
	if err := r.Create(&modelToAdd); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	CustomerTest = modelToAdd
}

func TestUpdateCustomer(t *testing.T) {
	r := repositories.GetCustomerRepository()

	modelToUpdate, err := r.GetByID(fmt.Sprint(CustomerTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	modelToUpdate.Name = "Updated_name"
	modelToUpdate.Email = "updated.email@test.com"
	modelToUpdate.Phone = "updated_phone"
	modelToUpdate.Document = "updated_document"
	modelToUpdate.Address = models.CustomerAddress{
		ZipCode:  "updated",
		Street:   "updated",
		Number:   "updated",
		District: "updated",
		City:     "updated",
		State:    "UP"}

	if err := r.Update(modelToUpdate); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expect := CustomerTest.Name != modelToUpdate.Name ||
		CustomerTest.Email != modelToUpdate.Email ||
		CustomerTest.Phone != modelToUpdate.Phone ||
		CustomerTest.Document != modelToUpdate.Document ||
		CustomerTest.Address != modelToUpdate.Address

	if !expect {
		t.Fatalf("UpdateCustomer test failed")
	}
}

func TestGetAllCustomers(t *testing.T) {
	r := repositories.GetCustomerRepository()
	if _, err := r.GetAll(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetCustomerById(t *testing.T) {
	r := repositories.GetCustomerRepository()
	if _, err := r.GetByID(fmt.Sprint(CustomerTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDeleteCustomer(t *testing.T) {
	r := repositories.GetCustomerRepository()

	modelToDelete, err := r.GetByID(fmt.Sprint(CustomerTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := r.Delete(modelToDelete); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
