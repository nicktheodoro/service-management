package repositories

import (
	"service-management/internal/pkg/database"

	"github.com/jinzhu/gorm"
)

// Find
func Find(where interface{}, out interface{}, associations []string, orders ...string) error {
	db := database.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	db = db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// First
func First(where interface{}, out interface{}, associations []string) (notFound bool, err error) {
	db := database.GetDB()
	for _, a := range associations {
		db = db.Preload(a)
	}
	err = db.Where(where).First(out).Error
	if err != nil {
		notFound = gorm.IsRecordNotFoundError(err)
	}
	return
}

// Create
func Create(value interface{}) error {
	return database.GetDB().Create(value).Error
}

// Updates
func Updates(where interface{}, value interface{}) error {
	return database.GetDB().Model(where).Updates(value).Error
}

// Delete
func Delete(value interface{}) (err error) {
	return database.GetDB().Delete(value).Error
}
