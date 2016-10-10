package transactioncounter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// test the transaction map applies our simple transaction rules correctly
func TestGenerateRandomTransactionCounter(test *testing.T) {
	assert := assert.New(test)

	counter := GenerateRandomTransactionCounter(100)

	// test counter has the correct number of transactions
	assert.Equal(100, len(counter.GetTransactions()))
}

// test
func TestGenerateRandomTickerMap(test *testing.T) {
	assert := assert.New(test)

	tickerMap, err := GenerateRandomTickerMap(100)

	assert.Nil(err)
	assert.NotEqual(len(tickerMap), 0)
}
