# fake_stock_alerts
Have you ever wanted fake stock alerts? Well now you can have some. Follow randomly generated Buy/Sell stock alerts for a
randomly generated user who has imaginary friends one randomly generated execution at a time! 

Build the binary
--

``` bash
docker run --rm -v $PWD:/usr/local/go/src/github.com/amoskyler/fake_stock_alerts -w /usr/local/go/src/github.com/amoskyler/fake_stock_alerts golang:1.7.1 go get; go build -v
```

Run tests (with coverage)
--

``` bash
docker run --rm -v $PWD:/usr/local/go/src/github.com/amoskyler/fake_stock_alerts -w /usr/local/go/src/github.com/amoskyler/fake_stock_alerts golang:1.7.1 go get; go test ./... -race -v
```

## Intro & Design

I implemented this little application in go as an intro into the language.
Prior to this I've written a bit of go packet code but in this project I explored
and got introduced to and familiarized with a lot of basic go principles.


General
--

1. The method `alert.GenerateAlerts(*User) Alerts` returns a custom list of alerts for the user provided user. It returns a slice of Alert in the type Alerts
    `Alerts.ToString() string` outputs a list in the format `"'<date>,<BUY|SELL>,<ticker>','<date>,<BUY|SELL>,<ticker>'"`
2. Unit tests have been implemented for the transaction counter logic [(here)](./transactioncounter/transactioncounter_test.go)
   as well as the Alert collection sorting logic [(here)](./alert/alert_test.go).
   Additional Tests: [here](./transactioncounter/transactioncounter_factory_test.go)
3. Good future testing
    -  Test struct/type constructor functions
    -  Test TickerMap state updates in the transactioncounter when new transactions are added
    -  Test alert property validity on generation
4. Alerts implements the sortable interface (Len, Less, Swap) which returns the alerts by highest volume moved with **O(n\*log(n))** complexity
    - This means alerts with net high SELLs will appear above a ticker with low BUYs
    - This can almost certainly be optimized as we can most likely generate the user's Alerts as we consume their friend's list of transactions (which is **O(n)**, in our case)
    - See [alert.go](./alert/alert.go)


There are *four* primary components which make the fake stock alerter a reality.
-

transaction.Transaction
--
A simple data structure which defines a single transaction (ticker, type, date).
  - generate with `transaction.GenerateRandomTransaction`


transactioncounter.Counter
--
The transaction counter is essentially a transaction collection.
Counter provides helpers for adding/calculating the TickerMap
(which stores all tickers associated to their net transaction flow)
When transactions are added to the Counter the updated TickerMap
is calculated via a simple foreach loop, incr/decrementing a hash, keyed by the ticker string
as it consumes a list of transactions.
The hash is calcualted in **O(n)** time as each transaction is consumed 
  - see [transactioncounter.go](./transactioncounter/transactioncounter.go)

The Counter also filters/errors out invalid transaction structures and applies date range validation
in order to filter out of interest transactions. Counter internally tracks both filtered and un-filtered transaction lists 

alert.Alert
--
Alert implements the sortable interface so it can be used with the native go package, "sort".
It implements it's sort a little differently than most high/low algorithms - it calculates position using the highest *absolute* values

user.User
--
A user has a friends list, as well as a series of transactions

## Basic Usage

``` golang
user := user.GenerateRandomUser(1, 20, 20)

alerts := alert.GenerateAlerts(user)
fmt.Println(alerts.ToString())

/* Outputs ~:
[
Alerts for user 0
5 Alerts

	"2,BUY,TWTR",
	"1,SELL,TWLO",
	"1,BUY,AAPL",
	"1,SELL,FB",
	"1,SELL,TSLA"
]
*/
```