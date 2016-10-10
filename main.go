package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/amoskyler/fake_stock_alerts/alert"
	"github.com/amoskyler/fake_stock_alerts/user"
)

func init() {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)
	fmt.Println("Application initialized with seed", seed, "BEHOLD MY RANDOMNESS", rand.Float64())
}

func main() {
	users := user.GenerateRandomUsers(1, 1, 20, 20)

	for _, user := range users {
		alerts := alert.GenerateAlerts(user)
		b, err := json.MarshalIndent(alerts, "", "	")
		if err != nil {
			fmt.Println("ERROR!", err)
		}

		fmt.Println(string(b))
	}

}
