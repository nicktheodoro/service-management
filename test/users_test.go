package test

import (
	"fmt"
	models "service-management/internal/pkg/models/users"
	"service-management/internal/pkg/repositories"
	"testing"
)

var userTest models.User

func TestAddUser(t *testing.T) {
	SetupTests()

	userToAdd := models.User{
		Firstname: "Antonio",
		Lastname:  "Paya",
		Email:     "antonio@teste.com",
		Hash:      "password",
		Role:      models.UserRole{RoleName: "user"},
	}
	r := repositories.GetUserRepository()
	if err := r.Create(&userToAdd); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	userTest = userToAdd
}

func TestUpdateUser(t *testing.T) {
	r := repositories.GetUserRepository()

	userToUpdate, err := r.GetByID(fmt.Sprint(userTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	userToUpdate.Firstname = "UpdatedFirstname"
	userToUpdate.Lastname = "UpdatedLastname"
	userToUpdate.Email = "updated.email@test.com"
	if err := r.Update(userToUpdate); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expect := userTest.Firstname != userToUpdate.Firstname || userTest.Lastname != userToUpdate.Lastname || userTest.Email != userToUpdate.Email
	if !expect {
		t.Fatalf("UpdateUser test failed")
	}
}

func TestGetAllUsers(t *testing.T) {
	r := repositories.GetUserRepository()
	if _, err := r.GetAll(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetUserById(t *testing.T) {
	r := repositories.GetUserRepository()
	if _, err := r.GetByID(fmt.Sprint(userTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	r := repositories.GetUserRepository()

	userToDelete, err := r.GetByID(fmt.Sprint(userTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := r.Delete(userToDelete); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = r.GetByID(fmt.Sprint(userTest.ID))
	if err == nil {
		t.Fatalf("Expected user to be deleted, but it still exists")
	}
}
