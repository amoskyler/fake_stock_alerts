/*
Package transactioncounter provides a data structure for enumerating and describing a collection of transactions
*/
package transactioncounter

import (
	"errors"
	"fmt"
	"time"

	"github.com/amoskyler/fake_stock_alerts/transaction"
)

type (
	// TickerMap is a hash keyed by the Ticker id with a value being the total count of transactions
	TickerMap map[transaction.Ticker]int
	// Counter is responsible for enumerating and describing properties of a group of Transactions
	Counter struct {
		transactions         []transaction.Transaction
		TickerMap            TickerMap
		filteredTransactions []transaction.Transaction
	}
)

// New constructs a new transactioncounter - it builds the counter with copies of the transactions
func New(transactions []transaction.Transaction) *Counter {
	counter := &Counter{transactions, map[transaction.Ticker]int{}, []transaction.Transaction{}}
	return counter
}

// GetTransactions returns a copy of the internal transaction slice
func (counter *Counter) GetTransactions() []transaction.Transaction {
	return counter.transactions
}

// AddTransactions adds a slice of transactions to the transaction slice - if applyTransaction is true, returns the updated transaction TickerMap
func (counter *Counter) AddTransactions(transactions []transaction.Transaction, applyTransactions bool) (TickerMap, error) {
	for _, t := range transactions {
		counter.AddTransaction(t, false)
	}

	if applyTransactions {
		return counter.ApplyAllTransactions(true)
	}
	return counter.TickerMap, nil
}

// AddTransaction adds a transaction to the transaction slice - if applyTransaction is true, returns the updated transaction TickerMap
func (counter *Counter) AddTransaction(t transaction.Transaction, applyTransaction bool) (TickerMap, error) {
	counter.transactions = append(counter.transactions, t)

	if applyTransaction {
		return counter.ApplyAllTransactions(true)
	}
	return counter.TickerMap, nil
}

// ApplyAllTransactions naively calculates the TickerMap.
// If a true applyFilters parameter is passed in, the TickerMap is calculated with a refreshed filteredTransactions list
func (counter *Counter) ApplyAllTransactions(applyFilters bool) (TickerMap, error) {
	var transactions *[]transaction.Transaction

	if applyFilters {
		transactions = counter.applyFilters()
	} else {
		transactions = &counter.transactions
	}

	for i, transaction := range *transactions {
		if ok := validateTransaction(transaction); !ok {
			return counter.TickerMap, fmt.Errorf(fmt.Sprintf("Invalid transaction %+v", transaction))
		}
		_, err := counter.ApplyTransaction(transaction)
		if err != nil {
			return counter.TickerMap, fmt.Errorf(fmt.Sprintf("Invalid transaction on index %d of %d. PrevErr: %s", i, len(counter.transactions), err))
		}
	}

	return counter.TickerMap, nil
}

func (counter *Counter) applyFilters() *[]transaction.Transaction {
	filtered := &[]transaction.Transaction{}
	for _, t := range counter.transactions {
		if ok := applyFilter(t); !ok {
			continue
		}

		*filtered = append(*filtered, t)
	}

	counter.filteredTransactions = *filtered
	return filtered
}

func applyFilter(t transaction.Transaction) bool {
	// apply filters here...
	elapsed := time.Since(t.Date)

	if elapsed > time.Hour*24*7 {
		return false
	}
	return true
}

func validateTransaction(t transaction.Transaction) bool {
	if time.Now().Before(t.Date) {
		return false
	}

	return true
}

// ApplyTransaction Applies an individual transaction to the TickerMap - Note: Does not pass through filter logic
func (counter *Counter) ApplyTransaction(t transaction.Transaction) (TickerMap, error) {
	tickerMap := counter.TickerMap
	if _, ok := tickerMap[t.Ticker]; !ok {
		tickerMap[t.Ticker] = 0
	}

	switch t.Type {
	case transaction.Buy:
		tickerMap[t.Ticker]++
		break
	case transaction.Sell:
		tickerMap[t.Ticker]--
		break
	default:
		return tickerMap, errors.New("Unsupported transaction type")
	}

	if tickerMap[t.Ticker] == 0 {
		delete(tickerMap, t.Ticker)
	}

	counter.TickerMap = tickerMap

	return tickerMap, nil
}

// ToSprintf verbosly returns the transactions associated to a counter
func (counter *Counter) ToSprintf() string {
	transactions := counter.transactions
	out := ""
	for i, t := range transactions {
		out += fmt.Sprintf("Transaction-%d:\n\tDate: %s\n\tTicker: %s\n\tType: %s\n", i, t.Date, t.Ticker, t.Type)
	}

	return out
}
