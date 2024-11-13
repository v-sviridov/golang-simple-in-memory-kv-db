package main

import "fmt"

func main() {
	infoMessage := `-----
-- This is a simple implementation of an in-memory key-value database with transaction support.
-----
This implementation supports the following operations:
- Set(key, value) - sets the value for the given key
- Get(key) - retrieves the value associated with the given key
- Delete(key) - deletes the value associated with the given key
- StartTransaction() - starts a new transaction
- Commit() - commits the current transaction
- Rollback() - rolls back the current transaction
-----
-- Run tests with 'go test -v'
-----`
	fmt.Println(infoMessage)
}
