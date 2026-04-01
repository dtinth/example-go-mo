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

Expected result from `go run .`:

```
Processed 3 users. Found 1 errors.
Job finished with accumulated errors:
 skipped user 3: player is missing username

Top Players:
- Player alice (ID: 1) - Win Rate: 80%
```
