package app

import "sync"

type Stock struct {
	LTP     float32
	Company string
	mu      sync.RWMutex
}

func NewStock(ltp float32, company string) *Stock {
	return &Stock{
		LTP:     ltp,
		Company: company,
	}
}

func (s *Stock) PriceUpdate(price float32) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.LTP = price
}

func (s *Stock) ShowPrice() float32 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.LTP
}
