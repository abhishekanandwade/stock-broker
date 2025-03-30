package app

type SellOrder struct {
	BaseOrder
}

func NewSellOrder(orderId string, stock *Stock, account *Account, qty uint32, price float32) *SellOrder {
	return &SellOrder{
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

func (s *SellOrder) Execute() error {
	if err := s.Account.Portfolio.RemoveStockFromPortfolio(s.Stock, s.Quantity); err != nil {
		s.Status = OrderRejected
		return err
	}

	s.Account.Deposit(s.Price * float32(s.Quantity))
	s.Status = OrderExecuted
	return nil
}
