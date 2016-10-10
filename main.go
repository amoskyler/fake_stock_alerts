package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/amoskyler/fake_stock_alerts/alert"
	"github.com/amoskyler/fake_stock_alerts/user"
)

func init() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	fmt.Println("Application initialized with seed", seed, "BEHOLD MY RANDOMNESS", rand.Float64(), "\n")
}

func main() {
	users := user.GenerateRandomUsers(5, 1, 20, 20)

	for i, user := range users {
		alerts := alert.GenerateAlerts(user)
		fmt.Println("Alerts for user", i, alerts.ToString(), "\n")
	}

}
