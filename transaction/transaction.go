/*
Package transaction provides a basic library for interacting with *fake* BUY/SELL transactions
*/
package transaction

import (
	"log"
	"time"
)

const (
	// Buy is the string constant used to denote a transaction's  (+) Type
	Buy Type = "BUY"
	// Sell is the string constant used to denote a transaction's (-) Type
	Sell Type = "SELL"
)

type (
	// Type represents the type assigned to the "type" (ex. BUY, SELL) of transaction
	Type string
	// Ticker represents the type assigned to the Ticker (ex GOOG, TWLO) value of a transaction
	Ticker string

	// Transaction contains the each property used to describe a transaction
	Transaction struct {
		Date   time.Time
		Ticker Ticker
		Type   Type
	}
)

/*
New is the constructor for a transaction. It accepts string parameters for easy of initialization from raw sources.
*/
func New(date string, ticker string, t string) *Transaction {
	transaction := new(Transaction)
	transaction.Ticker = Ticker(ticker)
	transaction.Type = Type(t)

	if date == "" {
		transaction.Date = time.Now()
	} else {
		transaction.Date = parseTime(date)
	}

	return transaction
}

// parse time in 24 hour ust time: 2006-01-02 15:04
func parseTime(date string) time.Time {
	const timeFormat = "2006-01-02 15:04"
	nDate, err := time.Parse(timeFormat, date)

	if err != nil {
		log.Fatal("Invalid transaction date time...\nERR:", err)
	}

	return nDate
}
