`service.go` contains a `GenerateReport` function that:

- processes a batch of users from the (mock) DB repository
- calculates their win rates
- generates a report of top players while gracefully accumulating errors for invalid profiles

To run:

```sh
# Set up dependencies
go mod tidy

# Run the main function to generate an example report
go run .

# Run the tests
go test -v
```
