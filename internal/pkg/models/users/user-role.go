package users

import (
	"service-management/internal/pkg/models"
	"time"
)

type UserRole struct {
	models.Model
	RoleName string `gorm:"not null;" json:"role_name"`
}

func (m *UserRole) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *UserRole) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
