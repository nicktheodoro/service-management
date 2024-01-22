package users

import (
	"service-management/internal/pkg/models"
	"time"
)

type User struct {
	models.Model
	Email     string   `gorm:"column:email;not null;unique_index:email" json:"email" form:"email"`
	Firstname string   `gorm:"column:firstname;not null;" json:"firstname" form:"firstname"`
	Lastname  string   `gorm:"column:lastname;not null;" json:"lastname" form:"lastname"`
	Hash      string   `gorm:"column:hash;not null;" json:"hash"`
	Role      UserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
