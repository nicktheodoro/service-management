package budgets

type BudgetStatus string

const (
	Requested  BudgetStatus = "REQUESTED"
	InProgress BudgetStatus = "PROGRESS"
	Completed  BudgetStatus = "COMPLETED"
	Failed     BudgetStatus = "FAILED"
	Canceled   BudgetStatus = "CANCELED"
)
