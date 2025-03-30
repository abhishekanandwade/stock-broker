package app

import (
	"errors"
	"sync"
)

type Portfolio struct {
	Stocks     map[*Stock]uint32
	TotalVlaue float32
	mu         sync.RWMutex
}

func NewPortfolio() *Portfolio {
	return &Portfolio{
		Stocks:     make(map[*Stock]uint32),
		TotalVlaue: 0.0,
	}
}

func (p *Portfolio) AddStockToPortfolio(stock *Stock, qty uint32) {
	p.mu.Lock()
	defer p.mu.Unlock()

	p.Stocks[stock] += qty
	p.TotalVlaue += stock.LTP * float32(qty)
}

func (p *Portfolio) RemoveStockFromPortfolio(stock *Stock, qty uint32) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	avlQty, exists := p.Stocks[stock]
	if !exists {
		return errors.New("stock not found in portfolio")
	}

	if qty > avlQty {
		return errors.New("insufficient stock quantity in portfolio")
	}

	p.Stocks[stock] -= qty
	if p.Stocks[stock] == 0 {
		delete(p.Stocks, stock)
	}

	return nil

}

func (p *Portfolio) GetHoldings() map[string]uint32 {
	p.mu.RLock()
	defer p.mu.RUnlock()
	holdings := make(map[string]uint32)
	for stock, qty := range p.Stocks {
		holdings[stock.Company] = qty
	}

	return holdings
}
