package main

// DbData represents the data of the database
type DbData struct {
	data map[string]string
}

// Set implements method of DbDataActions interface
func (dd *DbData) Set(kv KeyValue) {
	dd.data[kv.Key] = kv.Value
}

// Get implements method of DbDataActions interface
func (dd *DbData) Get(key string) (*string, bool) {
	value, ok := dd.data[key]
	return &value, ok
}

// Delete implements method of DbDataActions interface
func (dd *DbData) Delete(key string) {
	delete(dd.data, key)
}
