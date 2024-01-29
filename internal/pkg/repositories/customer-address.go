package repositories

import (
	models "service-management/internal/pkg/models/customers"
	"strconv"
)

type CustomerAddressRepository struct{}

var customerAddressRep *CustomerAddressRepository

func GetCustomerAddressRepository() *CustomerAddressRepository {
	if customerAddressRep == nil {
		customerAddressRep = &CustomerAddressRepository{}
	}
	return customerAddressRep
}

func (r *CustomerAddressRepository) GetAll() (*[]models.CustomerAddress, error) {
	var m []models.CustomerAddress
	err := Find(&models.CustomerAddress{}, &m, []string{}, "id asc")
	return &m, err
}

func (r *CustomerAddressRepository) GetByID(id string) (*models.CustomerAddress, error) {
	var m models.CustomerAddress
	where := models.CustomerAddress{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &m, []string{})

	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *CustomerAddressRepository) Create(model *models.CustomerAddress) error {
	return Create(&model)
}

func (r *CustomerAddressRepository) Update(model *models.CustomerAddress) error {
	where := models.CustomerAddress{}
	where.ID = model.ID
	return Updates(where, model)
}

func (r *CustomerAddressRepository) Delete(model *models.CustomerAddress) error {
	return Delete(model)
}
