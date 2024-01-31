package services

import (
	"service-management/internal/pkg/models"
)

type Service struct {
	models.Model
	Description string        `gorm:"not null;" json:"description" form:"description"`
	Value       int64         `gorm:"not null;" json:"value" form:"value"`
	Note        string        `json:"note" form:"note"`
	Status      ServiceStatus `gorm:"not null;" json:"status" form:"status"`
}
