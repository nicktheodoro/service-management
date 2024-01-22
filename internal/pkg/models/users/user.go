package users

import (
	"service-management/internal/pkg/models"
	"time"
)

type User struct {
	models.Model
	Email     string   `gorm:"not null;unique_index:email" json:"email" form:"email"`
	Firstname string   `gorm:"not null;" json:"firstname" form:"firstname"`
	Lastname  string   `gorm:"not null;" json:"lastname" form:"lastname"`
	Role      UserRole `gorm:"foreignKey:RoleID"`
	Hash      string   `gorm:"not null;" json:"-"`
	RoleID    uint64   `gorm:"not null;" json:"-"`
}

func (m *User) BeforeCreate() error {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return nil
}

func (m *User) BeforeUpdate() error {
	m.UpdatedAt = time.Now()
	return nil
}
