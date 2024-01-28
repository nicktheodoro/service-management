package users

import (
	"service-management/internal/pkg/models"
)

type User struct {
	models.Model
	Email     string   `gorm:"not null;unique_index:email" json:"email" form:"email"`
	Firstname string   `gorm:"not null;" json:"firstname" form:"firstname"`
	Lastname  string   `gorm:"not null;" json:"lastname" form:"lastname"`
	Hash      string   `gorm:"not null;" json:"-"`
	Role      UserRole `json:"role" form:"role"`
	RoleID    uint64   `gorm:"not null;" json:"-"`
}
