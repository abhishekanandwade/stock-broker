package main

import "stock-market/app"

func main() {
	sb := app.NewStockBroker()

	sb.CreateAccount("Test1", "test@mail.com")

	sb.GetAllAccounts()

	sb.AddFundsToAccount(1, 300.00)

	sb.GetAccountDetails(1)

	sb.AddStocks("RIL", "Reliance", 124.93)

	sb.GetAllStocks()

	account := sb.GetAccount(1)

	order1 := app.NewBuyOrder("1", sb.GetStock("RIL"), account, 2, 124.93)

	sb.ProcessOrder(order1)

}
