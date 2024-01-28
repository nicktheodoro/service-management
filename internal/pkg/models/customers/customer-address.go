package customers

import "service-management/internal/pkg/models"

type CustomerAddress struct {
	models.Model
	Name       string `gorm:"not null;" json:"name" form:"name"`
	ZipCode    string `gorm:"not null;" json:"zip_code" form:"zip_code"`
	Street     string `gorm:"not null;" json:"street" form:"street"`
	Number     string `gorm:"not null;" json:"number" form:"number"`
	Complement string `gorm:"not null;" json:"complement" form:"complement"`
	District   string `gorm:"not null;" json:"district" form:"district"`
	City       string `gorm:"not null;" json:"city" form:"city"`
	State      string `gorm:"not null;" json:"state" form:"state"`
	CustomerID uint64 `json:"customer_id" form:"customer_id"`
}
