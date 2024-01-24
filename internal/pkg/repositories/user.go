package repositories

import (
	models "service-management/internal/pkg/models/users"
	"strconv"
)

type UserRepository struct{}

var userRepository *UserRepository

func GetUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}

func (r *UserRepository) GetAll() (*[]models.User, error) {
	var m []models.User
	err := Find(&models.User{}, &m, []string{"Role"}, "id asc")
	return &m, err
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	var m models.User
	where := models.User{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &m, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var m models.User
	where := models.User{}
	where.Email = email
	_, err := First(&where, &m, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *UserRepository) Create(model *models.User) error {
	return Create(&model)
}

func (r *UserRepository) Update(model *models.User) error {
	where := models.User{}
	where.ID = model.ID

	return Updates(where, model)
}

func (r *UserRepository) Delete(model *models.User) error {
	return Delete(model)
}
