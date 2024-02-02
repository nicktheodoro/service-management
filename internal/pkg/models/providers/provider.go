package providers

import "service-management/internal/pkg/models"

type Provider struct {
	models.Model
	Name     string `gorm:"not null;" json:"name" form:"name"`
	Email    string `gorm:"not null;unique;" json:"email" form:"email"`
	Phone    string `gorm:"not null;unique;" json:"phone" form:"phone"`
	Document string `gorm:"not null;unique;" json:"document" form:"document"`
	Note     string `json:"note" form:"note"`
}
