package app

type Order interface {
	Execute() error
	GetStatus() OrderStatus
	SetStatus(status OrderStatus)
}

type BaseOrder struct {
	OrderId  string
	Stock    *Stock
	Account  *Account
	Quantity uint32
	Price    float32
	Status   OrderStatus
}

func (o *BaseOrder) GetStatus() OrderStatus {
	return o.Status
}

func (o *BaseOrder) SetStatus(status OrderStatus) {
	o.Status = status
}
