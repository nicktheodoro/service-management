package repositories

import (
	models "service-management/internal/pkg/models/services"
	"strconv"
)

type ServiceRepository struct{}

var serviceRepository *ServiceRepository

func GetServiceRepository() *ServiceRepository {
	if serviceRepository == nil {
		serviceRepository = &ServiceRepository{}
	}
	return serviceRepository
}

func (r *ServiceRepository) GetAll() (*[]models.Service, error) {
	var m []models.Service
	err := Find(&models.Service{}, &m, []string{}, "id asc")
	return &m, err
}

func (r *ServiceRepository) GetByID(id string) (*models.Service, error) {
	var m models.Service
	where := models.Service{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &m, []string{})

	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *ServiceRepository) Create(model *models.Service) error {
	return Create(&model)
}

func (r *ServiceRepository) Update(model *models.Service) error {
	where := models.Service{}
	where.ID = model.ID
	return Updates(where, model)
}

func (r *ServiceRepository) Delete(model *models.Service) error {
	return Delete(model)
}
