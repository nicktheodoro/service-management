package test

import (
	"fmt"
	models "service-management/internal/pkg/models/providers"
	"service-management/internal/pkg/repositories"
	"testing"
)

var ProviderTest models.Provider

func TestAddProvider(t *testing.T) {
	SetupTests()

	modelToAdd := models.Provider{
		Name:     "Antonio Paya",
		Email:    "antonio@teste.com",
		Phone:    "+5521990178696",
		Document: "123.456.789-00",
	}
	r := repositories.GetProviderRepository()
	if err := r.Create(&modelToAdd); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ProviderTest = modelToAdd
}

func TestUpdateProvider(t *testing.T) {
	r := repositories.GetProviderRepository()

	modelToUpdate, err := r.GetByID(fmt.Sprint(ProviderTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	modelToUpdate.Name = "Updated_name"
	modelToUpdate.Email = "updated.email@test.com"
	modelToUpdate.Phone = "updated_phone"
	modelToUpdate.Document = "updated_document"
	modelToUpdate.Note = "updated"

	if err := r.Update(modelToUpdate); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expect := ProviderTest.Name != modelToUpdate.Name ||
		ProviderTest.Email != modelToUpdate.Email ||
		ProviderTest.Phone != modelToUpdate.Phone ||
		ProviderTest.Document != modelToUpdate.Document ||
		ProviderTest.Note != modelToUpdate.Note

	if !expect {
		t.Fatalf("UpdateProvider test failed")
	}
}

func TestGetAllProviders(t *testing.T) {
	r := repositories.GetProviderRepository()
	if _, err := r.GetAll(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetProviderById(t *testing.T) {
	r := repositories.GetProviderRepository()
	if _, err := r.GetByID(fmt.Sprint(ProviderTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDeleteProvider(t *testing.T) {
	r := repositories.GetProviderRepository()

	modelToDelete, err := r.GetByID(fmt.Sprint(ProviderTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := r.Delete(modelToDelete); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
