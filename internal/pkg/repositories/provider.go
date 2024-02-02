package repositories

import (
	models "service-management/internal/pkg/models/providers"
	"strconv"
)

type ProviderRepository struct{}

var providerRep *ProviderRepository

func GetProviderRepository() *ProviderRepository {
	if providerRep == nil {
		providerRep = &ProviderRepository{}
	}
	return providerRep
}

func (r *ProviderRepository) GetAll() (*[]models.Provider, error) {
	var m []models.Provider
	err := Find(&models.Provider{}, &m, []string{}, "id asc")
	return &m, err
}

func (r *ProviderRepository) GetByID(id string) (*models.Provider, error) {
	var m models.Provider
	where := models.Provider{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &m, []string{})
	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *ProviderRepository) Create(model *models.Provider) error {
	return Create(model)
}

func (r *ProviderRepository) Update(model *models.Provider) error {
	where := models.Provider{}
	where.ID = model.ID
	return Updates(where, model)
}

func (r *ProviderRepository) Delete(model *models.Provider) error {
	return Delete(model)
}
