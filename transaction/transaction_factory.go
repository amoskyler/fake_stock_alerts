package transaction

import (
	"math/rand"

	"github.com/amoskyler/fake_stock_alerts/generators"
)

// GenerateRandomTransactions returns a collection of randomly generated transactions
func GenerateRandomTransactions(num int) []Transaction {
	transactions := []Transaction{}
	for i := 1; i <= num; i++ {
		transactions = append(transactions, GenerateRandomTransaction())
	}

	return transactions
}

// GenerateRandomTransaction initializes and returns a a randomly generated Transaction
func GenerateRandomTransaction() Transaction {
	tTime := generators.GenerateRandomPastDate(1, 10).Format("2006-01-02 15:04")
	ticker := getRandomTicker()
	tType := getRandomType()

	nTransaction := New(tTime, ticker, tType)

	return *nTransaction
}

func getRandomTicker() string {
	ticker := []string{"GOOG", "AAPL", "MSFT", "FB", "TWLO", "TWTR", "TSLA"}

	rIdx := rand.Intn(len(ticker))
	return ticker[rIdx]
}

func getRandomType() string {
	types := []string{"BUY", "SELL"}

	var rIdx int
	// randomly choose to pos 0 more often to generate slightly less equal "NET" results
	if rand.Intn(4)%2 == 0 {
		rIdx = 0
	} else {
		rIdx = rand.Intn(len(types))
	}

	return types[rIdx]
}
