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

func (r *UserRepository) Get() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{"Role"}, "id asc")
	return &users, err
}

func (r *UserRepository) GetByID(id string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &user, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) GetByUsername(email string) (*models.User, error) {
	var user models.User
	where := models.User{}
	where.Email = email
	_, err := First(&where, &user, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) Create(user *models.User) error {
	err := Create(&user)
	if err != nil {
		return err
	}

	return Save(&user)
}

func (r *UserRepository) Update(user *models.User) error {
	where := models.User{}
	where.ID = user.ID

	err := Updates(where, user)
	if err != nil {
		return err
	}

	return Save(&user)
}

func (r *UserRepository) Delete(user *models.User) error {
	return Delete(user)
}
