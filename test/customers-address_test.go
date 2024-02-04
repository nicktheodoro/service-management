package test

import (
	"fmt"
	models "service-management/internal/pkg/models/customers"
	"service-management/internal/pkg/repositories"
	"testing"
)

var CustomerAddressTest models.CustomerAddress

func TestAddCustomerAddress(t *testing.T) {
	SetupTests()
	customerRep := repositories.GetCustomerRepository()

	customer := models.Customer{
		Name:     "Test",
		Phone:    "+5521990170011",
		Document: "123.456.789-00",
	}

	if err := customerRep.Create(&customer); err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	modelToAdd := models.CustomerAddress{
		ZipCode:    "25955540",
		Street:     "Fortaleza",
		Number:     "139",
		District:   "Bom Retiro",
		City:       "Teresópolis",
		State:      "RJ",
		CustomerID: customer.ID,
	}

	r := repositories.GetCustomerAddressRepository()
	if err := r.Create(&modelToAdd); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	CustomerTest = customer
	CustomerAddressTest = modelToAdd
}

func TestUpdateCustomerAddress(t *testing.T) {
	customerID := fmt.Sprint(CustomerTest.ID)
	customerAddressID := fmt.Sprint(CustomerAddressTest.ID)
	r := repositories.GetCustomerAddressRepository()

	modelToUpdate, err := r.GetByID(customerID, customerAddressID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	modelToUpdate.ZipCode = "updated"
	modelToUpdate.Street = "updated"
	modelToUpdate.Number = "updated"
	modelToUpdate.District = "updated_document"
	modelToUpdate.City = "updated"
	modelToUpdate.State = "UP"

	if err := r.Update(modelToUpdate); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expect := CustomerAddressTest.ZipCode != modelToUpdate.ZipCode ||
		CustomerAddressTest.Street != modelToUpdate.Street ||
		CustomerAddressTest.Number != modelToUpdate.Number ||
		CustomerAddressTest.District != modelToUpdate.District ||
		CustomerAddressTest.City != modelToUpdate.City ||
		CustomerAddressTest.City != modelToUpdate.State

	if !expect {
		t.Fatalf("UpdateCustomer test failed")
	}
}

func TestGetAllCustomersAddress(t *testing.T) {
	customerID := fmt.Sprint(CustomerTest.ID)
	r := repositories.GetCustomerAddressRepository()

	if _, err := r.GetAll(customerID); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetCustomerAddressById(t *testing.T) {
	customerID := fmt.Sprint(CustomerTest.ID)
	customerAddressID := fmt.Sprint(CustomerAddressTest.ID)
	r := repositories.GetCustomerAddressRepository()

	if _, err := r.GetByID(customerID, customerAddressID); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDeleteCustomerAddress(t *testing.T) {
	customerID := fmt.Sprint(CustomerTest.ID)
	customerAddressID := fmt.Sprint(CustomerAddressTest.ID)
	r := repositories.GetCustomerAddressRepository()

	modelToDelete, err := r.GetByID(customerID, customerAddressID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := r.Delete(modelToDelete); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
