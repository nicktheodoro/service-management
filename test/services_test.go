package test

import (
	"fmt"
	models "service-management/internal/pkg/models/services"
	"service-management/internal/pkg/repositories"
	"testing"
)

var ServiceTest models.Service

func TestAddService(t *testing.T) {
	SetupTests()

	modelToAdd := models.Service{
		Description: "Interior painting",
		Value:       1000,
	}

	r := repositories.GetServiceRepository()
	if err := r.Create(&modelToAdd); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ServiceTest = modelToAdd
}

func TestUpdateService(t *testing.T) {
	r := repositories.GetServiceRepository()

	modelToUpdate, err := r.GetByID(fmt.Sprint(ServiceTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	modelToUpdate.Description = "updated"
	modelToUpdate.Value = 1993
	modelToUpdate.Note = "updated"

	if err := r.Update(modelToUpdate); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expect := ServiceTest.Description != modelToUpdate.Description ||
		ServiceTest.Value != modelToUpdate.Value ||
		ServiceTest.Note != modelToUpdate.Note

	if !expect {
		t.Fatalf("UpdateCustomer test failed")
	}
}

func TestGetAllServices(t *testing.T) {
	r := repositories.GetServiceRepository()
	if _, err := r.GetAll(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetServiceById(t *testing.T) {
	r := repositories.GetServiceRepository()
	if _, err := r.GetByID(fmt.Sprint(ServiceTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDeleteService(t *testing.T) {
	r := repositories.GetServiceRepository()

	modelToDelete, err := r.GetByID(fmt.Sprint(ServiceTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := r.Delete(modelToDelete); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
