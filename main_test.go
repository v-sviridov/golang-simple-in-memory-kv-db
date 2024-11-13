package main

import (
	"testing"
)

func Test1InMemoryDB(t *testing.T) {
	db := NewInMemoryDB()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	db.Commit()
	res, _ := db.Get("key1")
	if res != "value2" {
		t.Errorf("Expected name to be 'value2', got '%s'", res)
	}
}
func Test2InMemoryDB(t *testing.T) {
	db := NewInMemoryDB()
	db.Set("key1", "value1")
	db.StartTransaction()
	res, _ := db.Get("key1")
	if res != "value1" {
		t.Errorf("Expected name to be 'value1', got '%s'", res)
	}
	db.Set("key1", "value2")
	res, _ = db.Get("key1")
	if res != "value2" {
		t.Errorf("Expected name to be 'value2', got '%s'", res)
	}
	db.Rollback()
	res, _ = db.Get("key1")
	if res != "value1" {
		t.Errorf("Expected name to be 'value1', got '%s'", res)
	}
}

func Test3InMemoryDB(t *testing.T) {
	db := NewInMemoryDB()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	res, err := db.Get("key1")
	if res != "value2" {
		t.Errorf("Expected name to be 'value2', got '%s'", res)
	}
	db.StartTransaction()
	res, err = db.Get("key1")
	if res != "value2" {
		t.Errorf("Expected name to be 'value2', got '%s'", res)
	}
	db.Delete("key1")
	db.Commit()
	res, err = db.Get("key1")
	if err == nil {
		t.Errorf("Expected name to be none, got '%s'", res)
	}
	db.Commit()
	res, err = db.Get("key1")
	if err == nil {
		t.Errorf("Expected name to be none, got '%s'", res)
	}
}

func Test4InMemoryDB(t *testing.T) {
	db := NewInMemoryDB()
	db.Set("key1", "value1")
	db.StartTransaction()
	db.Set("key1", "value2")
	res, err := db.Get("key1")
	if res != "value2" {
		t.Errorf("Expected name to be 'value2', got '%s'", res)
	}
	db.StartTransaction()
	res, _ = db.Get("key1")
	if res != "value2" {
		t.Errorf("Expected name to be 'value2', got '%s'", res)
	}
	db.Delete("key1")
	res, err = db.Get("key1")
	if err == nil {
		t.Errorf("Expected name to be none, got '%s'", res)
	}
	db.Rollback()
	res, _ = db.Get("key1")
	if res != "value2" {
		t.Errorf("Expected name to be 'value2', got '%s'", res)
	}
	db.Commit()
	res, _ = db.Get("key1")
	if res != "value2" {
		t.Errorf("Expected name to be 'value2', got '%s'", res)
	}
}
