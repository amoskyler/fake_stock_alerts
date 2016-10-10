/*
 	Package user describes an individual user
 * Given Functions and their assumptions
 *
 * getFriendsListForUser(): []string {}
 * -? assume getFriendsListForUser is called from user context, so no params required
 *
 * getTradeTransactionsForUser(): []string {}
 * -? Same assumption as above - in this case the array returned by this method will be all transactions made by any given user
*/

package user

import (
	"fmt"

	"github.com/amoskyler/fake_stock_alerts/transaction"
	"github.com/amoskyler/fake_stock_alerts/transactioncounter"
	"github.com/nu7hatch/gouuid"
)

type (
	// FriendsList is a collection of other users
	FriendsList []User

	// User defines a basic user
	User struct {
		UUID         uuid.UUID
		FriendsList  FriendsList
		Transactions transactioncounter.Counter
	}
)

// ToSprintf returns a print friendly output of a friendsList
func (list FriendsList) ToSprintf() string {
	out := ""
	for i, user := range list {
		out += fmt.Sprintf("\tFriendsList: %d: %+v\n", i, user)
	}

	return out
}

// New returns a newly constructed User
func New(uuid uuid.UUID, friendsList FriendsList, transactions transactioncounter.Counter) *User {
	u := &User{}
	u.UUID = uuid
	u.FriendsList = friendsList
	u.Transactions = transactions

	return u
}

func (user *User) GetTradeTransactionsForUser() []transaction.Transaction {
	return user.Transactions.GetTransactions()
}

func (user *User) GetFriendsListForUser() FriendsList {
	return user.FriendsList
}
