package users

import "service-management/internal/pkg/models"

type UserRole struct {
	models.Model
	RoleName string `gorm:"not null;" json:"role_name"`
}
