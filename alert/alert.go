package alert

import (
	"errors"
	"fmt"
	"math"
	"sort"

	"github.com/amoskyler/fake_stock_alerts/transaction"
	"github.com/amoskyler/fake_stock_alerts/user"
)

// Alert is that the best you can do
type Alert struct {
	NetTransactions int
	TransactionType transaction.Type
	Ticker          transaction.Ticker
}

// Alerts is a slice of type Alert which implements the sortable interface
type Alerts []Alert

/*
  Implement sortable interface Len, Less, and Swap
*/
func (slice Alerts) Len() int {
	return len(slice)
}

// returns the lesser of two absolute values: this means higher absolute values will appear higher up in the list
func (slice Alerts) Less(i int, j int) bool {
	return math.Abs(float64(slice[i].NetTransactions)) < math.Abs(float64(slice[j].NetTransactions))
}

func (slice Alerts) Swap(i int, j int) {
	slice[i], slice[j] = slice[j], slice[i]

}

// GenerateAlerts creates a slice of Alerts out of a user's friends list
func GenerateAlerts(u user.User) Alerts {
	transactions := []transaction.Transaction{}
	for _, user := range u.GetFriendsListForUser() {
		transactions = append(transactions, user.GetTradeTransactionsForUser()...)
	}

	tickerMap, err := u.Transactions.AddTransactions(transactions, true)
	if err != nil {
		panic("Generated un-recoverable error while processing transactions for alerts")
	}

	alerts := Alerts{}
	for ticker, net := range tickerMap {
		alert, err := GenerateAlert(ticker, net)
		if err != nil {
			fmt.Println(err)
			continue
		}
		alerts = append(alerts, alert)
	}

	sort.Sort(sort.Reverse(alerts))
	return alerts
}

// GenerateAlert initializes a single alert
func GenerateAlert(ticker transaction.Ticker, netTransactions int) (Alert, error) {
	var (
		transactionType transaction.Type
		alert           Alert
	)

	if netTransactions == 0 {
		return alert, errors.New("No alert type exists for a net transaction value of 0")
	}

	if netTransactions > 0 {
		transactionType = transaction.Buy
	} else if netTransactions < 1 {
		transactionType = transaction.Sell
	}

	alert.NetTransactions = netTransactions
	alert.TransactionType = transactionType
	alert.Ticker = ticker

	return Alert{netTransactions, transactionType, ticker}, nil
}
