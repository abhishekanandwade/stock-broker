package app

type BuyOrder struct {
	BaseOrder
}

func NewBuyOrder(orderId string, stock *Stock, account *Account, qty uint32, price float32) *BuyOrder {
	return &BuyOrder{
		BaseOrder: BaseOrder{
			OrderId:  orderId,
			Stock:    stock,
			Account:  account,
			Quantity: qty,
			Price:    price,
			Status:   OrderPending,
		},
	}
}

func (b *BuyOrder) Execute() error {
	if err := b.Account.Withdraw(b.Price * float32(b.Quantity)); err != nil {
		b.Status = OrderRejected
		return err
	}

	b.Account.Portfolio.AddStockToPortfolio(b.Stock, b.Quantity)
	b.Status = OrderExecuted
	return nil
}
