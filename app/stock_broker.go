package app

import (
	"fmt"
	"sync"
)

type StockBroker struct {
	Accounts       map[int]*Account
	Stocks         map[string]*Stock
	AccountIdCount int
	OrderQueue     chan Order
	mu             sync.RWMutex
}

var (
	stockBrokerInstance *StockBroker
	once                sync.Once
)

func NewStockBroker() *StockBroker {
	stockBrokerInstance = &StockBroker{
		Accounts:   make(map[int]*Account),
		Stocks:     make(map[string]*Stock),
		OrderQueue: make(chan Order, 100),
	}

	return stockBrokerInstance
}

func (s *StockBroker) CreateAccount(name string, email string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	fmt.Println("––––––––––––––––––––––––––––––––––———————–")
	uniqNo := s.AccountIdCount + 1
	user := NewUser(uniqNo, name, email)
	portfolio := NewPortfolio()
	account := NewAccount(user, portfolio)

	s.Accounts[uniqNo] = account
	fmt.Println("Account Created Successfully, Your Account Number is ", uniqNo)
}

func (s *StockBroker) AddFundsToAccount(accountId int, amount float32) {
	account := s.Accounts[accountId]
	account.Deposit(amount)

	fmt.Println("––––––––––––––––––––––––––––––––––———————–")
	fmt.Println("Funds Added Successfully")

}

func (s *StockBroker) GetAllAccounts() {
	s.mu.RLock()
	defer s.mu.RUnlock()

	fmt.Println("––––––––––––––––––––––––––––––––––———————–")
	for accountId, account := range s.Accounts {
		fmt.Printf("AccountId: %d   Name: %s\n", accountId, account.User.Name)
	}
}

func (s *StockBroker) GetAccountDetails(accountId int) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	account := s.Accounts[accountId]
	fmt.Println("––––––––––––––––––––––––––––––––––———————–")
	fmt.Println("Name: ", account.User.Name)
	fmt.Println("Email: ", account.User.Email)
	fmt.Println("Balance: ", account.Balance)
	fmt.Println("Protfolio Value: ", account.Portfolio.TotalVlaue)
}

func (s *StockBroker) GetAccount(accountId int) *Account {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.Accounts[accountId]
}

func (s *StockBroker) AddStocks(symbol string, company string, value float32) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	stock := NewStock(value, company)
	s.Stocks[symbol] = stock

	fmt.Println("––––––––––––––––––––––––––––––––––———————–")
	fmt.Println("Stock added")
}

func (s *StockBroker) GetStock(symbol string) *Stock {
	stock, exists := s.Stocks[symbol]
	if !exists {
		fmt.Println("Invalid stock symbol")
	}
	return stock
}

func (s *StockBroker) GetAllStocks() {
	fmt.Println("––––––––––––––––––––––––––––––––––———————–")
	for symbol, stock := range s.Stocks {
		fmt.Println("Symbol: ", symbol)
		fmt.Println("Company: ", stock.Company)
		fmt.Println("Price: ", stock.LTP)

		fmt.Println("********")
	}
}

func (s *StockBroker) ProcessOrder(order Order) {
	fmt.Println("––––––––––––––––––––––––––––––––––———————–")
	if err := order.Execute(); err != nil {
		fmt.Println("Order Processing Failed: ", err.Error())
		return
	}

	fmt.Println("Order processed Successfully")
}
