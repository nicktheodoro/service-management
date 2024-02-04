package budgets

import (
	"service-management/internal/pkg/models"
	"service-management/internal/pkg/models/providers"
	"service-management/internal/pkg/models/services"
)

type BudgetItem struct {
	models.Model
	BudgetID   uint64             `gorm:"not null;" json:"-"`
	Provider   providers.Provider `json:"provider"`
	ProviderID uint64             ` gorm:"not null;" json:"provider_id" form:"provider_id"`
	Service    services.Service   `json:"service"`
	ServiceID  uint64             ` gorm:"not null;" json:"-" form:"service_id"`
}
