package transactioncounter

import "github.com/amoskyler/fake_stock_alerts/transaction"

// GenerateRandomTransactionCounter returns a brand new Counter with nTransactions added
func GenerateRandomTransactionCounter(nTransactions int) Counter {
	transactions := transaction.GenerateRandomTransactions(nTransactions)
	counter := New(transactions)

	return *counter
}

// GenerateRandomTickerMap returns a random TickerMap calculated over nTransactions where no filtering occurs
func GenerateRandomTickerMap(nTransactions int) (TickerMap, error) {
	counter := GenerateRandomTransactionCounter(nTransactions)

	return counter.ApplyAllTransactions(false)
}
