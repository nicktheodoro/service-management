package customers

import (
	"service-management/internal/pkg/models"
)

type Customer struct {
	models.Model
	Name     string          `gorm:"not null;" json:"name" form:"name"`
	Email    string          `gorm:"not null;unique_index:email" json:"email" form:"email"`
	Phone    string          `gorm:"not null;unique;" json:"phone" form:"phone"`
	Document string          `gorm:"not null;unique_index:document" json:"document" form:"document"`
	Address  CustomerAddress `json:"address" form:"address"`
}
