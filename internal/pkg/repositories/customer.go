package repositories

import (
	models "service-management/internal/pkg/models/customers"
	"strconv"
)

type CustomerRepository struct{}

var customerRep *CustomerRepository

func GetCustomerRepository() *CustomerRepository {
	if customerRep == nil {
		customerRep = &CustomerRepository{}
	}
	return customerRep
}

func (r *CustomerRepository) GetAll() (*[]models.Customer, error) {
	var m []models.Customer
	err := Find(&models.Customer{}, &m, []string{"Address"}, "id asc")
	return &m, err
}

func (r *CustomerRepository) GetByID(id string) (*models.Customer, error) {
	var m models.Customer
	where := models.Customer{}
	where.ID, _ = strconv.ParseUint(id, 10, 64)
	_, err := First(&where, &m, []string{"Address"})
	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *CustomerRepository) GetByEmail(email string) (*models.Customer, error) {
	var m models.Customer
	where := models.Customer{}
	where.Email = email
	_, err := First(&where, &m, []string{"Role"})
	if err != nil {
		return nil, err
	}
	return &m, err
}

func (r *CustomerRepository) Create(model *models.Customer) error {
	return Create(&model)
}

func (r *CustomerRepository) Update(model *models.Customer) error {
	where := models.Customer{}
	where.ID = model.ID
	return Updates(where, model)
}

func (r *CustomerRepository) Delete(model *models.Customer) error {
	return Delete(model)
}
