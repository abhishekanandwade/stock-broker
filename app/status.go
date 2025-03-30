package app

type OrderStatus int

const (
	OrderPending OrderStatus = iota
	OrderExecuted
	OrderRejected
)
