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

func (r *CustomerAddressRepository) GetAll(customerID string) (*[]models.CustomerAddress, error) {
	var m []models.CustomerAddress
	w := models.CustomerAddress{}
	w.CustomerID, _ = strconv.ParseUint(customerID, 10, 64)

	err := Find(&w, &m, []string{}, "id asc")
	return &m, err
}

func (r *CustomerAddressRepository) GetByID(customerID, id string) (*models.CustomerAddress, error) {
	var m models.CustomerAddress
	w := models.CustomerAddress{}
	w.ID, _ = strconv.ParseUint(id, 10, 64)
	w.CustomerID, _ = strconv.ParseUint(customerID, 10, 64)
	_, err := First(&w, &m, []string{})

	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *CustomerAddressRepository) Create(model *models.CustomerAddress) error {
	return Create(&model)
}

func (r *CustomerAddressRepository) Update(model *models.CustomerAddress) error {
	w := models.CustomerAddress{}
	w.ID = model.ID
	return Updates(&w, &model)
}

func (r *CustomerAddressRepository) Delete(model *models.CustomerAddress) error {
	return Delete(model)
}
