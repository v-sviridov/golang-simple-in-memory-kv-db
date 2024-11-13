package main

const (
	ActionSet    = "set"
	ActionDelete = "delete"
)

// Transaction represents a transaction with a list of operations
type Transaction struct {
	Operations []Operation
}

// Operation represents a database operation
type Operation struct {
	Action string // "set", "delete"
	KeyValue
}

// Set implements method of DbDataActions interface
func (tr *Transaction) Set(kv KeyValue) {
	for i, op := range tr.Operations {
		if op.Key == kv.Key {
			if op.Action == ActionDelete {
				tr.Operations[i].Action = ActionSet
			}
			tr.Operations[i].KeyValue.Value = kv.Value
			return
		}
	}
	tr.Operations = append(tr.Operations, Operation{
		Action:   ActionSet,
		KeyValue: kv,
	})
}

// Get implements method of DbDataActions interface
func (tr *Transaction) Get(key string) (*string, bool) {
	for i := len(tr.Operations) - 1; i >= 0; i-- {
		op := tr.Operations[i]
		if op.Key == key && op.Action == ActionSet {
			return &op.Value, true
		}
		if op.Key == key && op.Action == ActionDelete {
			return nil, true
		}
	}
	return nil, false
}

// Delete implements method of DbDataActions interface
func (tr *Transaction) Delete(key string) {
	for i, op := range tr.Operations {
		if op.Key == key {
			if op.Action == ActionSet {
				tr.Operations[i].Action = ActionDelete
			} // else do nothing because it's already deleted
			return
		}
	}
	tr.Operations = append(tr.Operations, Operation{
		Action:   ActionDelete,
		KeyValue: KeyValue{Key: key},
	})
}
