# golang-simple-in-memory-kv-db
This is a basic implementation of an in-memory key-value database with transaction support in Golang.

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Usage](#usage)
- [Testing](#testing)

## Introduction

This database is designed to provide a lightweight in-memory storage solution with support for key-value pairs and transactions. It allows users to perform basic operations such as `GET`, `SET`, and `DELETE` on data, as well as start, commit, and rollback transactions.

## Features

- In-memory key-value storage
- Support for transactions
- Nested transactions
- Basic operations: `GET`, `SET`, `DELETE`
- Commit changes in transactions
- Rollback changes in transactions


## Usage

```go
package main

import (
    "fmt"
)

func main() {
	db := NewInMemoryDB()

	// Set key-value pairs
	db.Set("name", "Alice")
	db.Set("age", "30")

	// Start a transaction
	db.StartTransaction()

	// Modify data within the transaction
	db.Set("name", "Bob")
	db.Delete("age")

	// Print the values
	fmt.Println("Values inside transaction:")
	res, _ := db.Get("name")
	fmt.Println("Name:", res) // Should return "Bob"
	res, _ = db.Get("age")
	fmt.Println("Age:", res) // Age should not be found

	// Commit the transaction
	db.Commit()

	// Print the modified values
	fmt.Println("\nValues after commit:")
	res, _ = db.Get("name")
	fmt.Println("Name:", res) // Should return "Bob"
	res, _ = db.Get("age")
	fmt.Println("Age:", res) // Age should be empty as it's deleted within the transaction
}
```

```bash
go run .
```

## Testing

To run tests, execute:

```bash
go test -v
```

---
