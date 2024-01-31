package services

type ServiceStatus int

const (
	Requested ServiceStatus = iota
	InProgress
	Completed
	Failed
	Canceled
)
