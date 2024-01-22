package models

import "time"

type Model struct {
	ID        uint64    `gorm:"column:id;primary_key;auto_increment;" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;not null;" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;" json:"updated_at"`
}
