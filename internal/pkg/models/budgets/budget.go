package budgets

import "service-management/internal/pkg/models"

type Budget struct {
	models.Model
	Description string       `gorm:"not null;" json:"description" form:"description"`
	Discount    int64        `json:"discount" form:"discount"`
	Surcharge   int64        `json:"surcharge" form:"surcharge"`
	Note        string       `json:"note" form:"note"`
	Status      BudgetStatus `gorm:"not null;" json:"status" form:"status"`
	CustomerID  uint64       `json:"customer_id" form:"customer_id"`
	Items       []BudgetItem `json:"-"`
}
