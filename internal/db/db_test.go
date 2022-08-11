package db_test

import (
	"testing"
	"time"

	"githuh.com/1995parham-teaching/students/internal/db"
)

func TestConnect(t *testing.T) {
	d, err := db.New(db.Config{
		URL:               "mongodb://127.0.0.1:27017",
		ConnectionTimeout: time.Second,
	})
	if err != nil {
		t.Fatalf("database connection error %s", err)
	}

	if d == nil {
		t.Fatalf("database is nil")
	}

	if d.Name() != "students" {
		t.Fatalf("database is not students")
	}
}

func TestNotConnect(t *testing.T) {
	_, err := db.New(db.Config{
		URL:               "mongodb://notfound",
		ConnectionTimeout: time.Second,
	})
	if err == nil {
		t.Fatalf("database connection must have error")
	}
}
