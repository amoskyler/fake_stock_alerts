package alert

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/amoskyler/fake_stock_alerts/user"
	"github.com/stretchr/testify/assert"
)

func init() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	fmt.Println("Application initialized with seed", seed, "BEHOLD MY RANDOMNESS", rand.Float64())
}

// TestGenerateAlerts test the transaction map applies our simple transaction rules correctly
func TestGenerateAlerts(test *testing.T) {
	assert := assert.New(test)

	user := user.GenerateRandomUser(10, 1, 0)
	alerts := GenerateAlerts(user)

	// test the alerts are sorted according to our rules
	if len(alerts) > 1 {
		assert.Equal(isSortedGreatestToLeast(alerts), true, "slice should be sorted greatest to least")
		assert.NotEqual(alerts[len(alerts)-1].NetTransactions, 0, "There should be no alerts for net 0 transactions")
	}

}

func BenchmarkGenerateAlerts(b *testing.B) {
	user := user.GenerateRandomUser(10, 1, 0)
	for i := 0; i < b.N; i++ {
		GenerateAlerts(user)
	}
}

func isSortedGreatestToLeast(alerts Alerts) bool {
	for i, alert := range alerts[1:] {
		if alerts[i].NetTransactions < alert.NetTransactions {
			return false
		}
	}

	return true
}
