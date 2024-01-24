package repositories

import (
	models "service-management/internal/pkg/models/users"
	"strconv"
)

type UserRoleRepository struct{}

var repository *UserRoleRepository

func GetUserRoleRepository() *UserRoleRepository {
	if repository == nil {
		repository = &UserRoleRepository{}
	}
	return repository
}

func (r *UserRoleRepository) GetAll() (*[]models.UserRole, error) {
	var m []models.UserRole
	err := Find(&models.UserRole{}, &m, []string{}, "id asc")
	return &m, err
}

func (r *UserRoleRepository) GetByID(id string) (*models.UserRole, error) {
	var m models.UserRole
	where := models.UserRole{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &m, []string{})
	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *UserRoleRepository) Create(model *models.UserRole) error {
	return Create(&model)
}

func (r *UserRoleRepository) Update(model *models.UserRole) error {
	where := models.UserRole{}
	where.ID = model.ID

	return Updates(where, model)
}

func (r *UserRoleRepository) Delete(model *models.UserRole) error {
	return Delete(model)
}
