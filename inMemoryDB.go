package main

import (
	"errors"
	"fmt"
)

// DataActions is interface for DbData and Transaction
type DataActions interface {
	Set(kv KeyValue)
	Get(key string) (*string, bool)
	Delete(key string)
}

// KeyValue represents a key-value pair
type KeyValue struct {
	Key   string
	Value string
}

// InMemoryDB represents an in-memory key-value database
type InMemoryDB struct {
	data         *DbData
	transactions []*Transaction
}

// NewInMemoryDB initializes a new in-memory database
func NewInMemoryDB() *InMemoryDB {
	return &InMemoryDB{
		data: &DbData{
			data: make(map[string]string),
		},
	}
}

// Get retrieves the value associated with the given key
func (db *InMemoryDB) Get(key string) (string, error) {
	err := errors.New(fmt.Sprintf("key %s not found", key))
	if len(db.transactions) > 0 {
		for i := len(db.transactions) - 1; i >= 0; i-- {
			transaction := db.transactions[i]
			if value, ok := transaction.Get(key); ok {
				if value == nil {
					return "", err
				}
				return *value, nil
			}
		}
	}
	if value, ok := db.data.Get(key); ok {
		return *value, nil
	}
	return "", err
}

// Set sets the value associated with the given key
func (db *InMemoryDB) Set(key, value string) {
	db.GetDataActions().Set(KeyValue{Key: key, Value: value})
}

// Delete deletes the value associated with the given key
func (db *InMemoryDB) Delete(key string) {
	db.GetDataActions().Delete(key)
}

// GetDataActions returns the object of the current transaction or the database.data
// they both implement the DataActions interface
func (db *InMemoryDB) GetDataActions() DataActions {
	if len(db.transactions) == 0 {
		return db.data
	}
	return db.transactions[len(db.transactions)-1]
}

// StartTransaction starts a new transaction
func (db *InMemoryDB) StartTransaction() {
	db.transactions = append(db.transactions, &Transaction{})
}

// Commit commits all changes in the current transaction
func (db *InMemoryDB) Commit() {
	lastTransaction, ok := db.popTransaction()
	if ok {
		data := db.GetDataActions()
		for _, op := range lastTransaction.Operations {
			switch op.Action {
			case ActionSet:
				data.Set(op.KeyValue)
			case ActionDelete:
				data.Delete(op.Key)
			}
		}
	}
}

// Rollback discards all changes in the current transaction
func (db *InMemoryDB) Rollback() {
	db.popTransaction()
}

// popTransaction pops the last transaction from the stack
func (db *InMemoryDB) popTransaction() (*Transaction, bool) {
	if len(db.transactions) == 0 {
		return nil, false
	}
	lastTransaction := db.transactions[len(db.transactions)-1]
	db.transactions = db.transactions[:len(db.transactions)-1]
	return lastTransaction, true
}
