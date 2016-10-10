## Build the binary
`docker run --rm -v $PWD:/usr/local/go/src/github.com/amoskyler/fake_stock_alerts -w /usr/local/go/src/github.com/amoskyler/fake_stock_alerts golang:1.7.1 go get; go build -v`

## Run tests
`docker run --rm -v $PWD:/usr/local/go/src/github.com/amoskyler/fake_stock_alerts -w /usr/local/go/src/github.com/amoskyler/fake_stock_alerts golang:1.7.1 go get; go test ./... -race -v`