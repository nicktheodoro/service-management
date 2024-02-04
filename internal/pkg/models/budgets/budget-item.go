package budgets

import (
	"service-management/internal/pkg/models"
	"service-management/internal/pkg/models/services"
)

type BudgetItem struct {
	models.Model
	BudgetID   uint64           `gorm:"not null;" json:"-"`
	ProviderID uint64           ` gorm:"not null;" json:"provider_id" form:"provider_id"`
	ServiceID  uint64           ` gorm:"not null;" json:"-" form:"service_id"`
	Service    services.Service `json:"service"`
}
