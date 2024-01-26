package test

import (
	"fmt"
	models "service-management/internal/pkg/models/users"
	"service-management/internal/pkg/repositories"
	"testing"
)

var userRoleTest models.UserRole

func TestCreateUserRole(t *testing.T) {
	SetupTests()

	roleToAdd := models.UserRole{
		RoleName: "new_role",
	}
	r := repositories.GetUserRoleRepository()
	if err := r.Create(&roleToAdd); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	userRoleTest = roleToAdd
}

func TestUpdateUserRole(t *testing.T) {
	r := repositories.GetUserRoleRepository()

	roleToUpdate, err := r.GetByID(fmt.Sprint(userRoleTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	roleToUpdate.RoleName = "UpdatedRoleName"
	if err := r.Update(roleToUpdate); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	expect := userRoleTest.RoleName != roleToUpdate.RoleName
	if !expect {
		t.Fatalf("UpdateUser test failed")
	}
}

func TestGetAllUserRoles(t *testing.T) {
	r := repositories.GetUserRoleRepository()
	if _, err := r.GetAll(); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestGetUserRoleById(t *testing.T) {
	r := repositories.GetUserRoleRepository()
	if _, err := r.GetByID(fmt.Sprint(userRoleTest.ID)); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

func TestDeleteRole(t *testing.T) {
	r := repositories.GetUserRoleRepository()

	roleToDelete, err := r.GetByID(fmt.Sprint(userRoleTest.ID))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if err := r.Delete(roleToDelete); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = r.GetByID(fmt.Sprint(userTest.ID))
	if err == nil {
		t.Fatalf("Expected user to be deleted, but it still exists")
	}
}
