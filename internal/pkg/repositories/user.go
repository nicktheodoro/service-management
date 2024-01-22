package repositories

import (
	"service-management/internal/pkg/database"
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

func (r *UserRepository) Get(id string) (*models.User, error) {
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

func (r *UserRepository) All() (*[]models.User, error) {
	var users []models.User
	err := Find(&models.User{}, &users, []string{"Role"}, "id asc")
	return &users, err
}

func (r *UserRepository) Query(q *models.User) (*[]models.User, error) {
	var users []models.User
	err := Find(&q, &users, []string{"Role"}, "id asc")
	return &users, err
}

func (r *UserRepository) Add(user *models.User) error {
	err := Create(&user)
	err = Save(&user)
	return err
}

func (r *UserRepository) Update(user *models.User) error {
	var userRole models.UserRole
	_, err := First(models.UserRole{UserID: user.ID}, &userRole, []string{})
	userRole.RoleName = user.Role.RoleName
	err = Save(&userRole)
	err = database.GetDB().Omit("Role").Save(&user).Error
	user.Role = userRole
	return err
}

func (r *UserRepository) Delete(user *models.User) error {
	err := database.GetDB().Unscoped().Delete(models.UserRole{UserID: user.ID}).Error
	err = database.GetDB().Unscoped().Delete(&user).Error
	return err
}
