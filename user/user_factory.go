package user

import (
	"log"
	"math/rand"

	"github.com/amoskyler/fake_stock_alerts/transaction"
	"github.com/amoskyler/fake_stock_alerts/transactioncounter"

	"github.com/nu7hatch/gouuid"
)

// GenerateRandomUser recursively generate a user + friends list. If third numTransactions param is 0, generates between 4 and 15 transactions per user (nested to numFriends..)
// be warned it gets big fast... you get the picture
func GenerateRandomUser(friendsListDepth int, numFriends int, nTransactions int) User {
	uuid := GenerateRandomUUID()
	if friendsListDepth < 1 {
		return *New(uuid, FriendsList{}, transactioncounter.Counter{})
	}
	friendsListDepth--
	friendsList := FriendsList{GenerateRandomUser(friendsListDepth, numFriends, nTransactions)}

	if nTransactions == 0 {
		nTransactions = rand.Intn(15-4) + 4
	}
	user := New(uuid, friendsList, *transactioncounter.New(transaction.GenerateRandomTransactions(nTransactions)))
	return *user
}

// GenerateRandomUsers generates a list of random user using a parametarized call to GenerateRandomUser
// if the third nTransactions paramter is 0, generates between 4 and 15 transactions per user
func GenerateRandomUsers(num int, friendsListDepth int, numFriends int, nTransactions int) []User {
	users := []User{}

	if nTransactions == 0 {
		nTransactions = rand.Intn(15-4) + 4
	}

	for i := 1; i <= num; i++ {
		users = append(users, GenerateRandomUser(friendsListDepth, numFriends, nTransactions))
	}
	return users
}

// GenerateRandomUUID generates a single random uuid
func GenerateRandomUUID() uuid.UUID {
	u4, err := uuid.NewV4()
	if err != nil {
		log.Fatal("Failed to generate a random UUID")
	}
	return *u4
}
