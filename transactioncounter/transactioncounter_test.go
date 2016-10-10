package transactioncounter

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/amoskyler/fake_stock_alerts/generators"
	"github.com/amoskyler/fake_stock_alerts/transaction"
	"github.com/stretchr/testify/assert"
)

func init() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	fmt.Println("Application initialized with seed", seed, "BEHOLD MY RANDOMNESS", rand.Float64())
}

// test the transaction map applies our simple transaction rules correctly
func TestApplyAllTransactions(test *testing.T) {
	assert := assert.New(test)

	/*
	   Tests the simple addition property of our transaction map.
	   All transaction dates should be within filter paramters
	*/
	transactions := []transaction.Transaction{}
	transactions = append(transactions, *transaction.New(generators.GenerateRandomPastDate(1, 6).Format("2006-01-02 15:04"), "GOOG", "BUY"))  // +1
	transactions = append(transactions, *transaction.New(generators.GenerateRandomPastDate(1, 6).Format("2006-01-02 15:04"), "GOOG", "SELL")) // -1
	//= 0

	transactions = append(transactions, *transaction.New(generators.GenerateRandomPastDate(1, 6).Format("2006-01-02 15:04"), "TWLO", "BUY"))  // +1
	transactions = append(transactions, *transaction.New(generators.GenerateRandomPastDate(1, 6).Format("2006-01-02 15:04"), "TWLO", "BUY"))  // +1
	transactions = append(transactions, *transaction.New(generators.GenerateRandomPastDate(1, 6).Format("2006-01-02 15:04"), "TWLO", "SELL")) // -1
	//= 1

	transactions = append(transactions, *transaction.New(generators.GenerateRandomPastDate(1, 6).Format("2006-01-02 15:04"), "TSLA", "SELL")) // -1
	transactions = append(transactions, *transaction.New(generators.GenerateRandomPastDate(1, 6).Format("2006-01-02 15:04"), "TSLA", "SELL")) // -1
	//=-2
	nCounter := New(transactions)
	_, err := nCounter.ApplyAllTransactions(true)
	assert.Nil(err)

	assert.Equal(0, nCounter.TickerMap["GOOG"], "Equal number of buys/sells should add up to 0")
	assert.Equal(1, nCounter.TickerMap["TWLO"], "2 buys and 1 sells should be net 1")
	assert.Equal(-2, nCounter.TickerMap["TSLA"], "Two sells should be -2")

	/*
	   Tests filtering of transactions > 1 week out of total transactions
	*/
	// clear transaction array & counter
	transactions = []transaction.Transaction{}

	transactions = append(transactions, *transaction.New(generators.GenerateRandomPastDate(8, 30).Format("2006-01-02 15:04"), "GOOG", "BUY")) // filter
	transactions = append(transactions, *transaction.New(generators.GenerateRandomPastDate(1, 6).Format("2006-01-02 15:04"), "GOOG", "SELL")) // -1
	// = -1

	nCounter = New(transactions)
	_, err = nCounter.ApplyAllTransactions(true)
	assert.Nil(err)

	assert.Equal(-1, nCounter.TickerMap["GOOG"], "Out of two values, only the one which has occured within a week should be applied")
	/*
	   Tests the filtering/erroring of invalid transactions
	*/
	// clear transaction array & counter
	transactions = []transaction.Transaction{}

	transactions = append(transactions, *transaction.New(generators.GenerateRandomFutureDate(8, 30).Format("2006-01-02 15:04"), "GOOG", "BUY")) // filter
	transactions = append(transactions, *transaction.New(generators.GenerateRandomFutureDate(8, 30).Format("2006-01-02 15:04"), "GOOG", "BUY")) // filter

	nCounter = New(transactions)
	_, err = nCounter.ApplyAllTransactions(true)

	assert.NotNil(err)
}

//TODO: Fix test
// func TestAddTransactions(test *testing.T) {
// 	assert := assert.New(test)

// 	transactions := transaction.GenerateRandomTransactions(5)
// 	nCounter := New(transactions)

// 	tickerMap := nCounter.TickerMap
// 	// Test that transactions are not applied when the false flag is given
// 	tm, err := nCounter.AddTransactions(transactions, false)
// 	assert.Equal(true, reflect.DeepEqual(tickerMap, tm), "The ticker maps should be equal in value")
// 	assert.Nil(err)

// 	// Test that transactions are correctly applied when the true flag is given
// 	nCounter = New(transactions)
// 	tickerMap = nCounter.TickerMap
// 	fmt.Println(tickerMap)
// 	tm, err = nCounter.AddTransaction(transactions[0], true)

// 	fmt.Println(tm, "\n", tickerMap, reflect.DeepEqual(tickerMap, tm))
// 	assert.Equal(false, reflect.DeepEqual(tickerMap, tm), "The ticker map values should differ in value")
// 	assert.Nil(err)

// }
